package check

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/usvc/db/cmd/db/configuration"
	"github.com/usvc/db/cmd/db/utils"
)

func check(_ *cobra.Command, _ []string) {
	utils.Connect(utils.ConnectOptions{
		Hostname:        configuration.Global.GetString(configuration.FlagHost),
		Port:            uint16(configuration.Global.GetUint(configuration.FlagPort)),
		Username:        configuration.Global.GetString(configuration.FlagUsername),
		Password:        configuration.Global.GetString(configuration.FlagPassword),
		RetryCount:      configuration.Global.GetUint(configuration.FlagRetryCount),
		RetryIntervalMs: time.Duration(configuration.Global.GetUint(configuration.FlagRetryIntervalMs)) * time.Millisecond,
		Log:             log,
	})

	utils.ExitSuccessfully()
}
