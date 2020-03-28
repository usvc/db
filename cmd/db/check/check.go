package check

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/usvc/db/cmd/db/configuration"
	"github.com/usvc/db/cmd/db/utils"
)

func check(_ *cobra.Command, _ []string) {
	utils.Connect(utils.ConnectOptions{
		Database:        configuration.Global.GetString(configuration.FlagDatabase),
		Hostname:        configuration.Global.GetString(configuration.FlagHost),
		Password:        configuration.Global.GetString(configuration.FlagPassword),
		Port:            uint16(configuration.Global.GetUint(configuration.FlagPort)),
		RetryCount:      configuration.Global.GetUint(configuration.FlagRetryCount),
		RetryIntervalMs: time.Duration(configuration.Global.GetUint(configuration.FlagRetryIntervalMs)) * time.Millisecond,
		Username:        configuration.Global.GetString(configuration.FlagUsername),
		Log:             log,
	})

	utils.ExitSuccessfully()
}
