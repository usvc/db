package check

import (
	"github.com/spf13/cobra"
	"github.com/usvc/db/cmd/db/configuration"
	"github.com/usvc/go-logger"
)

var (
	cmd *cobra.Command
	log logger.Logger
)

func GetCommand() *cobra.Command {
	if cmd == nil {
		cmd = &cobra.Command{
			Use:   "check",
			Short: "Verify a connection can be made with the provided credentials",
			PreRun: func(_ *cobra.Command, _ []string) {
				log = logger.New(logger.Options{
					Fields: map[string]interface{}{
						"command": "check",
					},
					Format: logger.Format(configuration.Global.GetString(configuration.FlagLogFormat)),
					Type:   logger.Type(configuration.Global.GetString(configuration.FlagLogType)),
				})
				log.Trace("'check' command triggered")
			},
			Run: check,
		}
	}
	return cmd
}
