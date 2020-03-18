package main

import (
	"github.com/spf13/cobra"
	"github.com/usvc/db/cmd/db/check"
	"github.com/usvc/db/cmd/db/configuration"
	"github.com/usvc/db/cmd/db/migrate"
)

func GetCommand() *cobra.Command {
	command := &cobra.Command{
		Use: "db",
		Run: run,
	}
	configuration.Global.ApplyToCobraPersistent(command)
	command.AddCommand(check.GetCommand())
	command.AddCommand(migrate.GetCommand())
	return command
}

func run(command *cobra.Command, args []string) {
	command.Help()
}

func main() {
	GetCommand().Execute()
}
