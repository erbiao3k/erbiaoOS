package net

import "net"

var CurrentIP, _ = InnerIP()

// InnerIP 获取内网IP地址
func InnerIP() (ip string, err error) {
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
