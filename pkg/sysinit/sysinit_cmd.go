package sysinit

import (
	"erbiaoOS/utils"
	"erbiaoOS/utils/file"
	"erbiaoOS/utils/login/sshd"
	"erbiaoOS/utils/net"
	"erbiaoOS/vars"
	"fmt"
)

// SysInit 系统初始化，
//		1、设置主机名；
//		2、关闭SELinux；
//		3、关闭firewalld服务；
//		4、卸载swap；
//		5、配置chrony时间同步；
//		6、优化Linux内核；
//		7、启用iptables；
//		8、开启ipvs；
//		9、安装docker;
func SysInit() {
	_, _, K8sClusterHost := vars.ClusterHostInfo()

	file.Chdir(vars.InitScriptDir)

	for f, cmd := range script {
		file.Create(f, cmd)
	}
	fmt.Println("正在为所有linux服务器上传系统初始化脚本")
	for _, host := range K8sClusterHost {
		if host.LanIp == net.CurrentIP {
			continue
		}

		sshd.Upload(&host, vars.InitScriptDir, vars.InitScriptDir)
	}

	fmt.Println("正在为所有linux服务器设置主机名")
	for _, host := range K8sClusterHost {
		hName := utils.GenerateHostname(host.Role, host.LanIp)
		sshd.RemoteExec(&host, setHostname+hName)
	}

	fmt.Println("初始化集群/etc/hosts文件")
	initHostfile()
	for _, host := range K8sClusterHost {
		sshd.Upload(&host, vars.TempDir+"hosts", SysConfigDir)
	}

	fmt.Println("正在为所有linux服务器关闭SELinux")
	sshd.LoopRemoteExec(K8sClusterHost, disableSELinux)

	fmt.Println("正在为所有linux服务器关闭firewalld服务")
	sshd.LoopRemoteExec(K8sClusterHost, disableFirewalld)

	fmt.Println("正在为所有linux服务器卸载swap")
	sshd.LoopRemoteExec(K8sClusterHost, disableSwap)

	fmt.Println("正在为所有linux服务器配置chrony服务")
	sshd.LoopRemoteExec(K8sClusterHost, fmt.Sprintf("sh -x %sEnableChrony.sh", vars.InitScriptDir))

	fmt.Println("正在为k8s集群节点linux服务器优化内核")
	sshd.LoopRemoteExec(K8sClusterHost, fmt.Sprintf("sh -x %sKernelOptimize.sh", vars.InitScriptDir))

	fmt.Println("正在为k8s集群节点安装基础软件")

	sshd.LoopRemoteExec(K8sClusterHost, softwareInstall)

	fmt.Println("正在为k8s集群节点启用iptables")
	sshd.LoopRemoteExec(K8sClusterHost, enableIptables)

	fmt.Println("正在为k8s集群节点开启ipvs")
	sshd.LoopRemoteExec(K8sClusterHost, fmt.Sprintf("sh -x %sEnableIpvs.sh", vars.InitScriptDir))

	fmt.Println("正在为k8s集群节点安装docker")
	sshd.LoopRemoteExec(K8sClusterHost, fmt.Sprintf("sh -x %sDockerInstall.sh", vars.InitScriptDir))

}
