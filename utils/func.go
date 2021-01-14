package utils

import (
	"github.com/mitchellh/mapstructure"
	"net"
)

func DecodeInterface(in interface{}, out interface{}) {
	err := mapstructure.Decode(in, out)
	if err != nil {
		panic("decode interface error " + err.Error())
	}
}

func LocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
