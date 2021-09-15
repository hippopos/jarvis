package server

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/Tsui89/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/hippopos/jarvis/src/pkg/linker"
)

var clients map[string]linker.SendData

func Server() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		port := viper.GetString("port")
		conn, err := net.Listen("tcp", ":"+port)
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		clients = make(map[string]linker.SendData)
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
		wg.Done()
	}()

	wg.Add(1)
	go restServer()
	wg.Wait()
}
func process(c net.Conn) {
	data, err := ioutil.ReadAll(c)
	if err != nil && err != io.EOF {
		fmt.Println(err.Error())
	}
	var recvData linker.SendData
	err = json.Unmarshal(data, &recvData)
	if err != nil {
		panic(err)
	}
	clientKey := strings.Join([]string{recvData.Name, recvData.Region}, "-")
	clients[clientKey] = recvData
}

func restServer() {
	r := gin.Default()

	r.Use(corsMiddleware)
	r.GET("/clients", clientsHandler)
	r.Run(":10000")
}

func clientsHandler(c *gin.Context) {
	//data, _ := json.Marshal(clients)
	br := response.NewBaseResponse()
	response.ResponseData(c, clients, *br)
}

func corsMiddleware(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Origin, X-Auth-Token, Accept, X-Custom-Header")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	c.Writer.Header().Add("Access-Control-Max-Age", "3600")
	if c.Request.Method == http.MethodOptions {
		c.Writer.WriteHeader(204)
		return
	}
	c.Next()
}