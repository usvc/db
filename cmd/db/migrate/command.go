package migrate

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
		log = logger.New(logger.Options{
			Format: logger.Format(configuration.Global.GetString("log-format")),
			Type:   logger.Type(configuration.Global.GetString("log-type")),
		})
		cmd = &cobra.Command{
			Use:   "migrate [flags] ./path/to/migrations",
			Short: "Perform database migrations",
			Run:   migrate,
		}
		inited = true
		log.Trace("initialised migrate command")
	}
	return cmd
}
