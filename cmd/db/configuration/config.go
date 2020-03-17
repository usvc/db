package configuration

import (
	"github.com/usvc/go-config"
)

var (
	Map = config.Map{
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
		"user": &config.String{
			Default:   "username",
			Shorthand: "u",
		},
		"password": &config.String{
			Default:   "password",
			Shorthand: "p",
		},
	}
)
