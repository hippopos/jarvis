package client

import (
	"time"

	"github.com/spf13/viper"

	"github.com/hippopos/jarvis/src/pkg/linker"
)

func Client() {
	//t := time.NewTicker(2 * time.Second)
	cn := viper.GetString("name")
	rg := viper.GetString("region")
	sd := viper.GetString("server-addr")
	for {
		client, err := linker.NewLinkerClient(cn, rg, sd)
		if err != nil {
			panic(err)
		}
		client.Heartbeat()
		client.Conn.Close()
		time.Sleep(30 * time.Minute)
	}
}
