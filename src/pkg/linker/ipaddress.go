package linker

import (
	"fmt"
	"net"
)

type IfaceData struct {
	IfaceName  string `json:"iface_name"`
	IpAddr     string `json:"ip_addr"`
}

func getIps() (ips []IfaceData) {

	ifaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			fmt.Println(iface.Name, "is down")
			continue
		}
		if iface.Flags&net.FlagLoopback != 0 {
			fmt.Println(iface.Name, "is loopback")
			continue
		}
		addrs, err := iface.Addrs()
		if err != nil {
			panic(err)
		}
		for _, addr := range addrs {
			ip := getIpFromAddr(addr)
			if ip == "" {
				continue
			}
			fmt.Println(iface.Name, ip)
			ips = append(ips, IfaceData{IfaceName: iface.Name, IpAddr: ip})
		}
	}
	return ips
}
func getIpFromAddr(addr net.Addr) string {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	if ip == nil || ip.IsLoopback() {
		return ""
	}
	ip = ip.To4()
	if ip == nil {
		return "" // not an ipv4 address
	}

	return ip.String()
}
