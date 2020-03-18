package check

import (
	"os"
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
		conf = &configuration.Global
		cmd = &cobra.Command{
			Use:   "check",
			Short: "verify a connection can be made with the provided credentials",
			Run:   run,
		}
		inited = true
		log.Trace("initialised check command")
	}
	return cmd
}

func run(_ *cobra.Command, _ []string) {
	dbOptions := db.Options{
		Hostname: conf.GetString("host"),
		Port:     uint16(conf.GetUint("port")),
		Username: conf.GetString("user"),
		Password: conf.GetString("password"),
	}
	log.Debugf("hostname: %s", dbOptions.Hostname)
	log.Debugf("port: %v", dbOptions.Port)
	log.Debugf("username: %s", dbOptions.Username)
	log.Debugf("password: %v", len(dbOptions.Password) > 0)
	db.Init(dbOptions)
	retriesLeft := conf.GetUint("retry-count")
	retryInterval := conf.GetUint("retry-interval-ms")
	for !check() {
		if retriesLeft == 0 {
			log.Error("no more retries left, giving up")
			os.Exit(1)
			break
		}
		log.Debugf("retrying in %v (%v tries left)...", time.Duration(retryInterval)*time.Millisecond, retriesLeft)
		<-time.After(time.Duration(retryInterval) * time.Millisecond)
		retriesLeft--
	}
	os.Exit(0)
}

func check() bool {
	if dbCheckErr := db.Check(); dbCheckErr != nil {
		log.Warnf("failed to access database (%s)", dbCheckErr)
		return false
	}
	return true
}
