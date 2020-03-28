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

func migrate(_ *cobra.Command, args []string) {
	// initialise
	utils.Connect(utils.ConnectOptions{
		Hostname:        configuration.Global.GetString("host"),
		Port:            uint16(configuration.Global.GetUint("port")),
		Username:        configuration.Global.GetString("username"),
		Password:        configuration.Global.GetString("password"),
		RetryCount:      configuration.Global.GetUint("retry-count"),
		RetryIntervalMs: time.Duration(configuration.Global.GetUint("retry-interval-ms")) * time.Millisecond,
		Log:             log,
	})
	pathToMigrations := path.Join(args...)
	if !filepath.IsAbs(pathToMigrations) {
		cwd, err := os.Getwd()
		if err != nil {
			log.Errorf("unable to access the current working directory")
			os.Exit(1)
		}
		pathToMigrations = path.Join(cwd, pathToMigrations)
	}

	// debug
	log.Debugf("migrations    : %s", pathToMigrations)

	// get file system list
	migrations, err := mysql.NewFromDirectory(pathToMigrations)
	if err != nil {
		log.Errorf("an error happened while retrieving migrations: '%s'", err)
		os.Exit(1)
	}
	sort.Sort(migrations)

	// migrate!
	var migrationError error
	migrationTableName := "migrations"
	connection := db.Get()
	appliedMigrationsCount := uint(0)
	lastAppliedMigrationIndex := 0
	for i := 0; i < len(migrations); i++ {
		if !conf.GetBool("all the way") && appliedMigrationsCount >= conf.GetUint("steps") {
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
		os.Exit(1)
	} else if lastAppliedMigrationIndex >= len(migrations)-1 {
		log.Info("migrations are up-to-date, cheers!")
	} else {
		unappliedMigrations := []string{}
		for j := lastAppliedMigrationIndex + 1; j < len(migrations); j++ {
			unappliedMigrations = append(unappliedMigrations, migrations[j].Name)
		}
		log.Infof("%v migrations left to go: %v", len(unappliedMigrations), unappliedMigrations)
	}
	os.Exit(0)
}
