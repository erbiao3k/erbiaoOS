package utils

import "net"

var CurrentIP, _ = GetInnerIP()

func GetInnerIP() (ip string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err == nil {
		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					return ipnet.IP.String(), nil
				}
			}
		}
	}
	panic(err)
	return "", err
}
