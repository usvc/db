package rollback

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/usvc/db/cmd/db/configuration"
	"github.com/usvc/db/cmd/db/utils"
	"github.com/usvc/go-db"
	"github.com/usvc/go-db/migration/mysql"
)

func rollback(command *cobra.Command, args []string) {
	// initialise
	utils.Connect(utils.ConnectOptions{
		Database:        configuration.Global.GetString(configuration.FlagDatabase),
		Hostname:        configuration.Global.GetString(configuration.FlagHost),
		Password:        configuration.Global.GetString(configuration.FlagPassword),
		Port:            uint16(configuration.Global.GetUint(configuration.FlagPort)),
		RetryCount:      configuration.Global.GetUint(configuration.FlagRetryCount),
		RetryIntervalMs: time.Duration(configuration.Global.GetUint(configuration.FlagRetryIntervalMs)) * time.Millisecond,
		Username:        configuration.Global.GetString(configuration.FlagUsername),
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

	// rollback!
	var rollbackError error
	migrationTableName := conf.GetString(FlagMigrationsTableName)
	connection := db.Get()
	processedMigrationsCount := uint(0)
	rolledBackMigrationNames := []string{}
	for i := len(migrations) - 1; i >= 0; i-- {
		if uint(len(rolledBackMigrationNames)) >= conf.GetUint(FlagSteps) {
			break
		}
		processedMigrationsCount++
		migration := migrations[i]
		log.Debugf("[%s] rolling back...", migration.Name)
		if err := migration.Validate(migrationTableName, connection); err != nil {
			if err != mysql.NoErrDoesNotExist {
				log.Debugf("[%s] resolving past error: '%s'", migration.Name, err)
				err = migration.Resolve(migrationTableName, connection)
				if err != nil {
					rollbackError = err
					break
				}
			} else {
				log.Infof("[%s] did not exist, continuing to previous migration", migration.Name)
				continue
			}
		}
		if err := migration.Rollback(migrationTableName, connection); err != nil {
			rollbackError = err
			break
		}
		log.Infof("[%s] rolled back successfully", migration.Name)
		rolledBackMigrationNames = append(rolledBackMigrationNames, migration.Name)
	}

	if rollbackError != nil {
		log.Error("rollback did not complete successfully: '%s'", rollbackError)
		utils.ExitErrorApplyingMigrations()
		return
	} else if processedMigrationsCount == uint(len(migrations)) {
		log.Warn("all possible migrations seem to have already been rolled back")
	} else {
		log.Infof("rolled back %v migration(s): [%s] successfully, cheers!", conf.GetUint(FlagSteps), strings.Join(rolledBackMigrationNames, ", "))
	}
	fmt.Println("exit status 0")
	utils.ExitSuccessfully()
}
