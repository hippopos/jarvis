package linker

import (
	"encoding/json"
	"net"
)

type Client interface {
	Connect() error
	Send(data ...interface{}) error
	Close() error
}

type LinkerClient struct {
	ServerAddress string
	Conn          net.Conn
}

//server 单实例
// client
func NewLinkerClient(s string) (*LinkerClient, error) {
	var lc LinkerClient
	var err error
	lc.ServerAddress = s
	lc.Conn, err = net.Dial("tcp", lc.ServerAddress)
	if err != nil {
		return nil, err
	}
	return &lc, nil
}

func (lc *LinkerClient) Connect() error {
	return nil
}

func (lc *LinkerClient) Send(data ...interface{}) error {
	databytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	length := len(databytes)
	wn := 0
	for ; ; {
		if wn == length {
			break
		}
		n, err := lc.Conn.Write(databytes[wn:length])
		if err != nil {
			break
		}
		wn += n
	}
	return nil
}
