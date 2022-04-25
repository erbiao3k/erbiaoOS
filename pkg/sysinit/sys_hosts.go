package sysinit

import (
	"erbiaoOS/setting"
	"erbiaoOS/utils"
	"erbiaoOS/utils/file"
)

// initHostfile 初始化/etc/hosts
func initHostfile() {
	initHost := SysHost + "# k8s section\n"
	loopExec := func(hostInfo []setting.HostInfo) {
		for _, host := range hostInfo {
			hostname := utils.GenerateHostname(host.Role, host.LanIp)
			initHost = initHost + host.LanIp + " " + hostname + "\n"
		}
	}

	loopExec(setting.K8sMasterHost)

	loopExec(setting.K8sNodeHost)

	initHost = initHost + "# end section \n"
	file.Create(hostsFile, initHost)
}
