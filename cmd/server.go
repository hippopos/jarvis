package cmd

import (
	"github.com/spf13/cobra"

	"github.com/hippopos/jarvis/src/server"
)

func newServer() *cobra.Command {
	command := &cobra.Command{
		Use:   "server",
		Short: "server",
		Run: func(cmd *cobra.Command, args []string) {
			server.Server()
		},
	}

	return command
}
