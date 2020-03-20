package check

import (
	"github.com/spf13/cobra"
	"github.com/usvc/db/cmd/db/configuration"
	"github.com/usvc/go-logger"
)

var (
	cmd    *cobra.Command
	log    logger.Logger
	inited bool
)

func GetCommand() *cobra.Command {
	if !inited {
		cmd = &cobra.Command{
			Use:   "check",
			Short: "Verify a connection can be made with the provided credentials",
			PreRun: func(_ *cobra.Command, _ []string) {
				log = logger.New(logger.Options{
					Format: logger.Format(configuration.Global.GetString("log-format")),
					Type:   logger.Type(configuration.Global.GetString("log-type")),
				})
			},
			Run: check,
		}
		inited = true
	}
	return cmd
}
