package migrate

import "github.com/usvc/go-config"

const (
	FlagLatest              = "latest"
	FlagMigrationsTableName = "table-name"
	FlagSteps               = "steps"
)

var (
	conf = &config.Map{
		FlagLatest: &config.Bool{
			Shorthand: "A",
			Usage:     "defines whether migrations should run to completion",
		},
		FlagMigrationsTableName: &config.String{
			Default:   "migrations",
			Shorthand: "n",
			Usage:     "defines the name of the table used to store migration steps",
		},
		FlagSteps: &config.Uint{
			Default:   1,
			Shorthand: "s",
			Usage:     "defines the number of steps to migrate",
		},
	}
)
