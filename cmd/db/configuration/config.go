package configuration

import (
	"github.com/usvc/go-config"
	"github.com/usvc/go-logger"
)

var (
	Global = config.Map{
		"driver": &config.String{
			Default:   "mysql",
			Shorthand: "d",
			Usage:     "defines the database driver to use (only MySQL is supported for now)",
		},
		"host": &config.String{
			Default:   "localhost",
			Shorthand: "H",
			Usage:     "defines the hostname at which the database server can be reached at",
		},
		"log-format": &config.String{
			Default:   string(logger.FormatText),
			Shorthand: "f",
			Usage:     "defines the format of the logger according to `github.com/usvc/go-logger`",
		},
		"log-type": &config.String{
			Default:   string(logger.TypeStdout),
			Shorthand: "t",
			Usage:     "defines the type of the logger according to `github.com/usvc/go-logger`",
		},
		"port": &config.Uint{
			Default:   3306,
			Shorthand: "P",
			Usage:     "defines the port at which the database server can be reached",
		},
		"username": &config.String{
			Default:   "user",
			Shorthand: "u",
			Usage:     "defines the username of the user used to login to the database server",
		},
		"password": &config.String{
			Default:   "password",
			Shorthand: "p",
			Usage:     "defines the password of the user used to login to the database server",
		},
		"retry-interval-ms": &config.Uint{
			Default:   3000,
			Shorthand: "R",
			Usage:     "defines the number of milliseconds in between database connection retry attempts",
		},
		"retry-count": &config.Uint{
			Default:   5,
			Usage:     "defines the number of times the application will re-attempt a failed connection to the database",
			Shorthand: "r",
		},
	}
)
