package initialize

import (
	"erbiaoOS/pkg/hostname"
	"erbiaoOS/setting"
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

	// 设置主机名(所有linux节点操作)
	linuxServer := [][]setting.HostInfo{k8sMasters, k8sNodes}
	k8sServer := [][]setting.HostInfo{k8sMasters, k8sNodes}

	initScriptDir := "tempData/initScript/"
	remoteTempData := "/opt/tempData"
	log.Println("正在为所有linux服务器上传系统初始化脚本")
	for _, temp := range linuxServer {
		for _, node := range temp {
			for _, f := range file.List(initScriptDir) {
				sshd2.SftpUploadFile(node.RemoteIp, node.User, node.Password, node.Port, initScriptDir+f, remoteTempData)
			}
		}
	}

	log.Println("正在为所有linux服务器设置主机名")
	for _, temp := range linuxServer {
		for _, node := range temp {
			hName := hostname.GenerateHostname(node.Role, node.LanIp)
			// 登陆到服务器，若服务器主机名包含localhost则按照Generate规则重命名主机名
			node := node
			cmd := fmt.Sprintf("mkdir %s -p && sh -x %s/SetHostname.sh %s", remoteTempData, remoteTempData, hName)
			sshd2.RemoteSshExec(node.RemoteIp, node.User, node.Password, node.Port, cmd)
		}
	}

	// 临时函数 for sshd.RemoteSshExec
	loopExec := func(severList [][]setting.HostInfo, cmd string) {
		for _, temp := range severList {
			for _, node := range temp {
				node := node
				sshd2.RemoteSshExec(node.RemoteIp, node.User, node.Password, node.Port, cmd)
			}
		}
	}

	log.Println("正在为所有linux服务器关闭SELinux")
	loopExec(linuxServer, fmt.Sprintf("sh -x %sDisableSELinux.sh", remoteTempData))

	log.Println("正在为所有linux服务器关闭firewalld服务")
	loopExec(linuxServer, fmt.Sprintf("sh -x %sDisableFirewalld.sh", remoteTempData))

	log.Println("正在为所有linux服务器卸载swap")
	loopExec(linuxServer, fmt.Sprintf("sh -x %sDisableSwap.sh", remoteTempData))

	log.Println("正在为所有linux服务器配置chrony服务")
	loopExec(linuxServer, fmt.Sprintf("sh -x %sEnableChrony.sh", remoteTempData))

	log.Println("正在为k8s集群节点linux服务器优化内核")
	loopExec(k8sServer, fmt.Sprintf("sh -x %sKernelOptimize.sh", remoteTempData))

	log.Println("正在为k8s集群节点基础软件安装")
	loopExec(k8sServer, fmt.Sprintf("sh -x %sSoftwareInstall.sh", remoteTempData))

	log.Println("正在为k8s集群节点启用iptables")
	loopExec(k8sServer, fmt.Sprintf("sh -x %sEnableIptables.sh", remoteTempData))

	log.Println("正在为k8s集群节点开启ipvs")
	loopExec(k8sServer, fmt.Sprintf("sh -x %sEnableIpvs.sh", remoteTempData))

	log.Println("正在为k8s集群节点安装docker")
	loopExec(k8sServer, fmt.Sprintf("sh -x %sDockerInstall.sh", remoteTempData))

}
