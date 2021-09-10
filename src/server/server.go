package server

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"time"
)

func Server() {
	conn, err := net.Listen("tcp", ":9999")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	for {
		select {
		case <-time.After(1 * time.Second):
		}
		c, err := conn.Accept()
		if err != nil {
			fmt.Println(err.Error())
		}
		go process(c)

	}
}
func process(c net.Conn){
	data, err := ioutil.ReadAll(c)
	if err != nil && err != io.EOF {
		fmt.Println(err.Error())
	}
	fmt.Println(string(data))
}