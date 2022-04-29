package net

import (
	"log"
	"net"
)

// ParseIPv4 检测是否是合法的IPv4地址
func ParseIPv4(ip string) bool {
	address := net.ParseIP(ip)
	if address == nil {
		return false
	}
	return true
}

// ParseMultiIPv4 多IPv4检测
func ParseMultiIPv4(IPs []string) {
	for _, ip := range IPs {
		if !ParseIPv4(ip) {
			log.Fatalf("IP地址不合法：%s", ip)
		}
	}
}
