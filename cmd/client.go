package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/hippopos/jarvis/src/client"
)

var (
	name   string
	region string
)

func newClient() *cobra.Command {
	command := &cobra.Command{
		Use:   "client",
		Short: "client",
		Run: func(cmd *cobra.Command, args []string) {
			client.Client()
		},
	}
	command.PersistentFlags().StringVarP(&name, "name", "n", "client", "client name")
	command.PersistentFlags().StringVarP(&region, "region", "r", "region", "client region")

	viper.BindPFlags(command.PersistentFlags())
	return command
}
