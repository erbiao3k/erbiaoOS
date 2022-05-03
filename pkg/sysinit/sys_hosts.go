package sysinit

import (
	"erbiaoOS/utils"
	"erbiaoOS/utils/file"
	"erbiaoOS/vars"
)

// initHostfile 初始化/etc/hosts
func initHostfile() {
	_, _, K8sClusterHost := vars.ClusterHostInfo()

	var SysHost string
	if !file.Exist(HostsFileBak) {
		file.Copy(HostsFileBak, hostsFile)
		SysHost = file.Read(hostsFile)
	} else {
		SysHost = file.Read(HostsFileBak)
	}

	initHost := SysHost + "# k8s section\n"

	for _, host := range K8sClusterHost {
		hostname := utils.GenerateHostname(host.Role, host.LanIp)
		initHost = initHost + host.LanIp + " " + hostname + "\n"
	}

	initHost = initHost + "# end section \n"
	file.Create(vars.TempDir+"hosts", initHost)
}
