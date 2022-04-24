package sysinit

import (
	customConst "erbiaoOS/const"
	"erbiaoOS/setting"
	"erbiaoOS/utils"
	"erbiaoOS/utils/file"
	sshd2 "erbiaoOS/utils/login/sshd"
	"fmt"
	"log"
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
func SysInit(clusterHost *setting.ClusterHost) {
	// 获取各服务节点清单
	k8sMasters := clusterHost.K8sMaster
	k8sNodes := clusterHost.K8sNode

	utils.Chdir(customConst.InitScriptDir)

	for f, cmd := range script {
		file.Create(f, cmd)
	}

	// 设置主机名(所有linux节点操作)
	linuxServer := [][]setting.HostInfo{k8sMasters, k8sNodes}
	k8sServer := [][]setting.HostInfo{k8sMasters, k8sNodes}

	log.Println("正在为所有linux服务器上传系统初始化脚本")
	for _, hosts := range linuxServer {
		for _, host := range hosts {
			sshd2.UploadDir(host.LanIp, host.User, host.Password, host.Port, customConst.InitScriptDir, customConst.DeployDir)
		}
	}

	log.Println("正在为所有linux服务器设置主机名")
	for _, hosts := range linuxServer {
		for _, host := range hosts {
			hName := utils.GenerateHostname(host.Role, host.LanIp)
			// 登陆到服务器，若服务器主机名包含localhost则按照Generate规则重命名主机名
			sshd2.RemoteSshExec(host.LanIp, host.User, host.Password, host.Port, setHostname+hName)
		}
	}

	// 临时函数 for sshd.RemoteSshExec
	loopExec := func(hostList [][]setting.HostInfo, cmd string) {
		for _, hosts := range hostList {
			for _, host := range hosts {
				sshd2.RemoteSshExec(host.LanIp, host.User, host.Password, host.Port, cmd)
			}
		}
	}

	log.Println("正在为所有linux服务器关闭SELinux")
	loopExec(linuxServer, disableSELinux)

	log.Println("正在为所有linux服务器关闭firewalld服务")
	loopExec(linuxServer, disableFirewalld)

	log.Println("正在为所有linux服务器卸载swap")
	loopExec(linuxServer, disableSwap)

	log.Println("正在为所有linux服务器配置chrony服务")
	loopExec(linuxServer, fmt.Sprintf("sh -x %sEnableChrony.sh", customConst.InitScriptDir))

	log.Println("正在为k8s集群节点linux服务器优化内核")
	loopExec(k8sServer, fmt.Sprintf("sh -x %sKernelOptimize.sh", customConst.InitScriptDir))

	log.Println("正在为k8s集群节点基础软件安装")
	loopExec(k8sServer, softwareInstall)

	log.Println("正在为k8s集群节点启用iptables")
	loopExec(k8sServer, enableIptables)

	log.Println("正在为k8s集群节点开启ipvs")
	loopExec(k8sServer, fmt.Sprintf("sh -x %sEnableIpvs.sh", customConst.InitScriptDir))

	log.Println("正在为k8s集群节点安装docker")
	loopExec(k8sServer, fmt.Sprintf("sh -x %sDockerInstall.sh", customConst.InitScriptDir))

}
