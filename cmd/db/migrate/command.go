package migrate

import (
	"os"
	"path"
	"path/filepath"
	"sort"
	"time"

	"github.com/spf13/cobra"
	"github.com/usvc/db/cmd/db/configuration"
	"github.com/usvc/go-db"
	"github.com/usvc/go-db/migration/mysql"
	"github.com/usvc/go-logger"
)

var (
	cmd    *cobra.Command
	log    logger.Logger
	inited bool
)

func GetCommand() *cobra.Command {
	if !inited {
		log = logger.New()
		cmd = &cobra.Command{
			Use:   "migrate [flags] ./path/to/migrations",
			Short: "perform database migrations",
			Run:   run,
		}
		conf.ApplyToCobra(cmd)
		inited = true
		log.Trace("initialised migrate command")
	}
	return cmd
}

func run(_ *cobra.Command, args []string) {
	// initialise
	dbOptions := db.Options{
		Hostname: configuration.Global.GetString("host"),
		Port:     uint16(configuration.Global.GetUint("port")),
		Username: configuration.Global.GetString("username"),
		Password: configuration.Global.GetString("password"),
	}
	retriesLeft := configuration.Global.GetUint("retry-count")
	retryInterval := time.Duration(configuration.Global.GetUint("retry-interval-ms")) * time.Millisecond
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
	log.Debugf("hostname      : %s", dbOptions.Hostname)
	log.Debugf("port          : %v", dbOptions.Port)
	log.Debugf("username      : %s", dbOptions.Username)
	log.Debugf("password      : %v", len(dbOptions.Password) > 0)
	log.Debugf("migrations    : %s", pathToMigrations)
	log.Debugf("retry interval: %v", retryInterval)
	log.Debugf("retry limit   : %v", retriesLeft)

	// make contact with database
	if err := db.Init(dbOptions); err != nil {
		log.Warn("an error happened while initialising the database connection: '%s'", err)
	}
	for !check() {
		if retriesLeft == 0 {
			log.Error("no more retries left, giving up and exitting...")
			os.Exit(1)
			break
		}
		log.Debugf("retrying in %v (%v tries left)...", retryInterval, retriesLeft)
		<-time.After(retryInterval)
		retriesLeft--
	}
	log.Infof(
		"successfully connected to '%s@%s:%v' (using password: %v)",
		dbOptions.Username,
		dbOptions.Hostname,
		dbOptions.Port,
		len(dbOptions.Password) > 0,
	)

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
	for i := 0; i < len(migrations); i++ {
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
	}
	if migrationError != nil {
		log.Error("migrations did not complete successfully: '%s'", migrationError)
		os.Exit(1)
	}
	log.Info("migrations are up-to-date, cheers!")
	os.Exit(0)
}

func check() bool {
	if dbCheckErr := db.Check(); dbCheckErr != nil {
		log.Warnf("failed to access database (%s)", dbCheckErr)
		return false
	}
	return true
}
