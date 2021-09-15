package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/hippopos/jarvis/src/server"
)

var port string

func newServer() *cobra.Command {
	command := &cobra.Command{
		Use:   "server",
		Short: "server",
		Run: func(cmd *cobra.Command, args []string) {
			server.Server()
		},
	}
	command.PersistentFlags().StringVarP(&port, "port", "p", "9999", "server port")

	viper.BindPFlags(command.PersistentFlags())
	return command
}
