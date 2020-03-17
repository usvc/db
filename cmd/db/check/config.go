package check

import "github.com/usvc/go-config"

var localConf = config.Map{
	"retry-interval-ms": &config.Uint{
		Default:   3000,
		Shorthand: "R",
	},
	"retry-count": &config.Uint{
		Default:   0,
		Shorthand: "r",
	},
}
