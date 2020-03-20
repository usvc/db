package check

import (
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/usvc/db/cmd/db/configuration"
	"github.com/usvc/db/cmd/db/utils"
)

func check(_ *cobra.Command, _ []string) {
	utils.Connect(utils.ConnectOptions{
		Hostname:        configuration.Global.GetString("host"),
		Port:            uint16(configuration.Global.GetUint("port")),
		Username:        configuration.Global.GetString("username"),
		Password:        configuration.Global.GetString("password"),
		RetryCount:      configuration.Global.GetUint("retry-count"),
		RetryIntervalMs: time.Duration(configuration.Global.GetUint("retry-interval-ms")) * time.Millisecond,
		Log:             log,
	})

	os.Exit(0)
}
