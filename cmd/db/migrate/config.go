package migrate

import "github.com/usvc/go-config"

var (
	conf = &config.Map{
		"all the way": &config.Bool{
			Shorthand: "A",
		},
		"steps": &config.Uint{
			Default:   1,
			Shorthand: "s",
		},
	}
)
