package rollback

import "github.com/usvc/go-config"

const (
	FlagMigrationsTableName = "table-name"
	FlagSteps               = "steps"
)

var (
	conf = &config.Map{
		FlagMigrationsTableName: &config.String{
			Default:   "migrations",
			Shorthand: "n",
			Usage:     "defines the name of the table used to store migration steps",
		},
		FlagSteps: &config.Uint{
			Default:   1,
			Shorthand: "s",
			Usage:     "defines the number of steps to roll back",
		},
	}
)
