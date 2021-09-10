package client

import (
	"fmt"

	"github.com/hippopos/jarvis/src/pkg/linker"
)

func Client() {
	//t := time.NewTicker(2 * time.Second)
	client, err := linker.NewLinkerClient("localhost:9999")
	if err != nil {
		panic(err)
	}
	for _, v := range getIps() {
		fmt.Println(v)
		client.Send(v)
	}
	client.Conn.Close()
}
