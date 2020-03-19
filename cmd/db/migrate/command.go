package migrate

import (
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
	"github.com/usvc/db/cmd/db/configuration"
	"github.com/usvc/go-config"
	"github.com/usvc/go-db"
	"github.com/usvc/go-logger"
)

var (
	cmd    *cobra.Command
	conf   *config.Map
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
		conf = &configuration.Global
		inited = true
		log.Trace("initialised migrate command")
	}
	return cmd
}

func run(_ *cobra.Command, args []string) {
	// initialise
	dbOptions := db.Options{
		Hostname: conf.GetString("host"),
		Port:     uint16(conf.GetUint("port")),
		Username: conf.GetString("username"),
		Password: conf.GetString("password"),
	}
	retriesLeft := conf.GetUint("retry-count")
	retryInterval := time.Duration(conf.GetUint("retry-interval-ms")) * time.Millisecond
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

	// start
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

}

func check() bool {
	if dbCheckErr := db.Check(); dbCheckErr != nil {
		log.Warnf("failed to access database (%s)", dbCheckErr)
		return false
	}
	return true
}
