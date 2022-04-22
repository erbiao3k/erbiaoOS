package utils

import (
	"strings"
)

/*
检查当前主机名中是否包含localhost，有则修改主机名，且添加集群主机与IP对应关系到/etc/hosts
*/

// addZero IP补0，如：192.168.1.1变为192.168.001.001
func addZero(ip string) string {
	n1 := strings.Split(ip, ".")[0]
	n2 := strings.Split(ip, ".")[1]
	n3 := strings.Split(ip, ".")[2]
	n4 := strings.Split(ip, ".")[3]

	for i := 1; i < 3; i++ {
		if len(n1) < 3 {
			n1 = "0" + n1
		}
		if len(n2) < 3 {
			n2 = "0" + n2
		}
		if len(n3) < 3 {
			n3 = "0" + n3
		}
		if len(n4) < 3 {
			n4 = "0" + n4
		}
	}

	return n1 + "." + n2 + "." + n3 + "." + n4

}

// GenerateHostname 依据IP和节点角色生成主机名
func GenerateHostname(role, ip string) string {
	ipstr := addZero(ip)
	n1 := strings.Split(ipstr, ".")[2]
	n2 := strings.Split(ipstr, ".")[3]
	return role + n1 + n2
}
