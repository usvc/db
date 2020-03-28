package migrate

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
			Use:   "migrate [flags] ./path/to/migrations",
			Short: "Perform database migrations",
			PreRun: func(_ *cobra.Command, _ []string) {
				log = logger.New(logger.Options{
					Fields: map[string]interface{}{
						"command": "migrate",
					},
					Format: logger.Format(configuration.Global.GetString(configuration.FlagLogFormat)),
					Type:   logger.Type(configuration.Global.GetString(configuration.FlagLogType)),
				})
				log.Trace("'migrate' command triggered")
			},
			Run: migrate,
		}
		conf.ApplyToCobra(cmd)
	}
	return cmd
}
