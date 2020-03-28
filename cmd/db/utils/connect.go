package utils

import (
	"database/sql"
	"time"

	"github.com/usvc/go-db"
	"github.com/usvc/go-logger"
)

type ConnectOptions struct {
	Hostname        string
	Port            uint16
	Username        string
	Password        string
	Database        string
	RetryCount      uint
	RetryIntervalMs time.Duration
	Log             logger.Logger
}

func Connect(opts ConnectOptions) *sql.DB {
	dbOptions := db.Options{
		Hostname: opts.Hostname,
		Port:     opts.Port,
		Username: opts.Username,
		Password: opts.Password,
		Database: opts.Database,
	}
	retriesLeft := opts.RetryCount
	retryInterval := opts.RetryIntervalMs
	if opts.Log == nil {
		opts.Log = logger.New(logger.Options{Type: logger.TypeNoOp})
	}

	// debug
	opts.Log.Debugf("hostname      : %s", dbOptions.Hostname)
	opts.Log.Debugf("port          : %v", dbOptions.Port)
	opts.Log.Debugf("username      : %s", dbOptions.Username)
	opts.Log.Debugf("password      : %v", len(dbOptions.Password) > 0)
	opts.Log.Debugf("database      : %v", dbOptions.Database)
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
			ExitErrorConnectingToDatabase()
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
