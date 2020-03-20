package utils

import (
	"database/sql"
	"os"
	"time"

	"github.com/usvc/db/cmd/db/configuration"
	"github.com/usvc/go-db"
	"github.com/usvc/go-logger"
)

type ConnectOptions struct {
	Hostname        string
	Port            uint16
	Username        string
	Password        string
	RetryCount      uint
	RetryIntervalMs time.Duration
	Log             logger.Logger
}

func Connect(opts ConnectOptions) *sql.DB {
	dbOptions := db.Options{
		Hostname: configuration.Global.GetString("host"),
		Port:     uint16(configuration.Global.GetUint("port")),
		Username: configuration.Global.GetString("username"),
		Password: configuration.Global.GetString("password"),
	}
	retriesLeft := configuration.Global.GetUint("retry-count")
	retryInterval := time.Duration(configuration.Global.GetUint("retry-interval-ms")) * time.Millisecond

	// debug
	opts.Log.Debugf("hostname      : %s", dbOptions.Hostname)
	opts.Log.Debugf("port          : %v", dbOptions.Port)
	opts.Log.Debugf("username      : %s", dbOptions.Username)
	opts.Log.Debugf("password      : %v", len(dbOptions.Password) > 0)
	opts.Log.Debugf("retry interval: %v", retryInterval)
	opts.Log.Debugf("retry limit   : %v", retriesLeft)

	// make contact with database
	if err := db.Init(dbOptions); err != nil {
		opts.Log.Warn("an error happened while initialising the database connection: '%s'", err)
	}
	for {
		if err := db.Check(); err == nil {
			opts.Log.Infof(
				"successfully connected to '%s@%s:%v' (using password: %v)",
				dbOptions.Username,
				dbOptions.Hostname,
				dbOptions.Port,
				len(dbOptions.Password) > 0,
			)
			break
		} else if retriesLeft == 0 {
			opts.Log.Error("no more retries left, giving up and exitting...")
			os.Exit(1)
			break
		} else {
			opts.Log.Warnf("database connection failed: '%s'", err)
		}
		opts.Log.Debugf("retrying in %v (%v tries left)...", retryInterval, retriesLeft)
		<-time.After(retryInterval)
		retriesLeft--
	}
	return db.Get()
}
