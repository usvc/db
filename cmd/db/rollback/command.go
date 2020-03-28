package rollback

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
			Aliases: []string{"rb", "down"},
			Use:     "rollback [flags] ./path/to/migrations",
			Short:   "Rolls back database migrations",
			PreRun: func(_ *cobra.Command, _ []string) {
				log = logger.New(logger.Options{
					Fields: map[string]interface{}{
						"command": "rollback",
					},
					Format: logger.Format(configuration.Global.GetString(configuration.FlagLogFormat)),
					Type:   logger.Type(configuration.Global.GetString(configuration.FlagLogType)),
				})
				log.Trace("'rollback' command triggered")
			},
			Run: rollback,
		}
		conf.ApplyToCobra(cmd)
	}
	return cmd
}
