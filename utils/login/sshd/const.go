package sshd

import (
	"erbiaoOS/pkg/config"
)

// LoopRemoteExec 单个Linux指令在清单上的机器执行
func LoopRemoteExec(hosts []config.HostInfo, cmd string) {
	for _, host := range hosts {

		RemoteExec(&host, cmd)
	}
}

// LoopRemoteMultiExec 多个Linux指令在清单上的机器执行
func LoopRemoteMultiExec(hosts []config.HostInfo, cmds []string) {
	for _, host := range hosts {
		for _, cmd := range cmds {
			RemoteExec(&host, cmd)
		}
	}
}
