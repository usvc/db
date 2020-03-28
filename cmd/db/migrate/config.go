package migrate

import "github.com/usvc/go-config"

const (
	FlagLatest = "latest"
	FlagSteps  = "steps"
)

var (
	conf = &config.Map{
		FlagLatest: &config.Bool{
			Shorthand: "A",
			Usage:     "defines whether migrations should run to completion (only applies to upwards migrations)",
		},
		FlagSteps: &config.Uint{
			Default:   1,
			Shorthand: "s",
			Usage:     "defines the number of steps to migrate (both upwards/downwards)",
		},
	}
)
