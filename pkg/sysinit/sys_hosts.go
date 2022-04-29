package sysinit

import (
	myConst "erbiaoOS/const"
	"erbiaoOS/pkg/config"
	"erbiaoOS/utils"
	"erbiaoOS/utils/file"
)

// initHostfile 初始化/etc/hosts
func initHostfile() {
	initHost := SysHost + "# k8s section\n"
	loopExec := func(hostInfo []config.HostInfo) {
		for _, host := range hostInfo {
			hostname := utils.GenerateHostname(host.Role, host.LanIp)
			initHost = initHost + host.LanIp + " " + hostname + "\n"
		}
	}

	loopExec(config.K8sMasterHost)

	loopExec(config.K8sNodeHost)

	initHost = initHost + "# end section \n"
	file.Create(myConst.TempDir+"hosts", initHost)
}
