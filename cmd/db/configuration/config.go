package configuration

import (
	"github.com/usvc/go-config"
)

var (
	Global = config.Map{
		"driver": &config.String{
			Default:   "mysql",
			Shorthand: "d",
		},
		"host": &config.String{
			Default:   "localhost",
			Shorthand: "H",
		},
		"port": &config.Uint{
			Default:   3306,
			Shorthand: "P",
		},
		"username": &config.String{
			Default:   "user",
			Shorthand: "u",
		},
		"password": &config.String{
			Default:   "password",
			Shorthand: "p",
		},
		"retry-interval-ms": &config.Uint{
			Default:   3000,
			Shorthand: "R",
		},
		"retry-count": &config.Uint{
			Default:   0,
			Shorthand: "r",
		},
	}
)
