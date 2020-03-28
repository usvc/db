package migrate

import (
	"os"
	"path"
	"path/filepath"
	"sort"
	"time"

	"github.com/spf13/cobra"
	"github.com/usvc/db/cmd/db/configuration"
	"github.com/usvc/db/cmd/db/utils"
	"github.com/usvc/go-db"
	"github.com/usvc/go-db/migration/mysql"
)

func migrate(command *cobra.Command, args []string) {
	// initialise
	utils.Connect(utils.ConnectOptions{
		Hostname:        configuration.Global.GetString(configuration.FlagHost),
		Port:            uint16(configuration.Global.GetUint(configuration.FlagPort)),
		Username:        configuration.Global.GetString(configuration.FlagUsername),
		Password:        configuration.Global.GetString(configuration.FlagPassword),
		RetryCount:      configuration.Global.GetUint(configuration.FlagRetryCount),
		RetryIntervalMs: time.Duration(configuration.Global.GetUint(configuration.FlagRetryIntervalMs)) * time.Millisecond,
		Log:             log,
	})
	pathToMigrations := path.Join(args...)
	if len(pathToMigrations) == 0 {
		log.Errorf("failed to read in a valid path to use as the migrations source directory, see help as follows:")
		command.Help()
		utils.ExitErrorInvalidArguments()
	}
	if !filepath.IsAbs(pathToMigrations) {
		cwd, err := os.Getwd()
		if err != nil {
			log.Errorf("failed to access the current working directory")
			utils.ExitErrorAccessingFileSystem()
		}
		pathToMigrations = path.Join(cwd, pathToMigrations)
	}

	// debug
	log.Debugf("migrations    : %s", pathToMigrations)

	// get file system list
	migrations, err := mysql.NewFromDirectory(pathToMigrations)
	if err != nil {
		log.Errorf("failed to retrieve migrations: '%s'", err)
		utils.ExitErrorReadingMigrations()
	}
	sort.Sort(migrations)

	// migrate!
	var migrationError error
	migrationTableName := "migrations"
	connection := db.Get()
	appliedMigrationsCount := uint(0)
	lastAppliedMigrationIndex := 0
	for i := 0; i < len(migrations); i++ {
		if !conf.GetBool(FlagLatest) && appliedMigrationsCount >= conf.GetUint(FlagSteps) {
			break
		}
		lastAppliedMigrationIndex = i
		migration := migrations[i]
		log.Debugf("[%s] processing...", migration.Name)
		if err := migration.Validate(migrationTableName, connection); err != nil && err != mysql.NoErrDoesNotExist {
			log.Debugf("[%s] resolving past error: '%s'", migration.Name, err)
			err = migration.Resolve(migrationTableName, connection)
			if err != nil {
				migrationError = err
				break
			}
		}
		if err := migration.Apply(migrationTableName, connection); err != nil {
			if err == mysql.NoErrAlreadyApplied {
				log.Debugf("[%s] already been applied", migration.Name)
				continue
			}
			migrationError = err
			break
		}
		log.Infof("[%s] applied successfully", migration.Name)
		appliedMigrationsCount++
	}

	if migrationError != nil {
		log.Error("migrations did not complete successfully: '%s'", migrationError)
		utils.ExitErrorApplyingMigrations()
	} else if lastAppliedMigrationIndex < len(migrations)-1 {
		unappliedMigrations := []string{}
		for j := lastAppliedMigrationIndex + 1; j < len(migrations); j++ {
			unappliedMigrations = append(unappliedMigrations, migrations[j].Name)
		}
		log.Infof("%v migrations left to go: %v", len(unappliedMigrations), unappliedMigrations)
		utils.ExitSuccessfullyWithIncompleteMigrations()
	}
	log.Info("migrations are up-to-date, cheers!")
	utils.ExitSuccessfully()
}
