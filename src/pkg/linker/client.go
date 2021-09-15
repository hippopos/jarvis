package linker

import (
	"encoding/json"
	"net"
	"time"
)

type Client interface {
	Connect() error
	Send(data ...interface{}) error
	Close() error
}

type LinkerClient struct {
	Name          string
	Region        string
	ServerAddress string
	Conn          net.Conn
}

type SendData struct {
	Name      string      `json:"name"`
	Region    string      `json:"region"`
	Data      []IfaceData `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

//server 单实例
// client
func NewLinkerClient(name, region, s string) (*LinkerClient, error) {
	var lc LinkerClient
	var err error
	lc.Name = name
	lc.Region = region
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

func (lc *LinkerClient) Heartbeat() error {

	var sd SendData
	sd.Name = lc.Name
	sd.Region = lc.Region
	sd.Data = getIps()
	sd.Timestamp = time.Now().Unix()
	databytes, err := json.Marshal(sd)
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
