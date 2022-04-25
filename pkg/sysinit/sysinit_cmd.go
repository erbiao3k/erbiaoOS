package sysinit

import (
	myConst "erbiaoOS/const"
	"erbiaoOS/setting"
	"erbiaoOS/utils"
	"erbiaoOS/utils/file"
	"erbiaoOS/utils/login/sshd"
	"fmt"
)

// LoopExec 临时函数 for sshd.RemoteSshExec
var LoopExec = func(hostList [][]setting.HostInfo, cmd string) {
	for _, hosts := range hostList {
		for _, host := range hosts {
			sshd.RemoteSshExec(host.LanIp, host.User, host.Password, host.Port, cmd)
		}
	}
}

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

	utils.Chdir(myConst.InitScriptDir)

	for f, cmd := range script {
		file.Create(f, cmd)
	}

	// 设置主机名(所有linux节点操作)

	fmt.Println("正在为所有linux服务器上传系统初始化脚本")
	for _, hosts := range setting.LinuxServer {
		for _, host := range hosts {
			if host.LanIp == utils.CurrentIP {
				continue
			}
			sshd.Upload(host.LanIp, host.User, host.Password, host.Port, myConst.InitScriptDir, myConst.InitScriptDir)
		}
	}

	fmt.Println("正在为所有linux服务器设置主机名")
	for _, hosts := range setting.LinuxServer {
		for _, host := range hosts {
			hName := utils.GenerateHostname(host.Role, host.LanIp)
			// 登陆到服务器，若服务器主机名包含localhost则按照Generate规则重命名主机名
			sshd.RemoteSshExec(host.LanIp, host.User, host.Password, host.Port, setHostname+hName)
		}
	}

	fmt.Println("初始化集群/etc/hosts文件")
	initHostfile()
	for _, hosts := range setting.LinuxServer {
		for _, host := range hosts {
			if host.LanIp == utils.CurrentIP {
				continue
			}
			sshd.Upload(host.LanIp, host.User, host.Password, host.Port, hostsFile, SysConfigDir)
		}
	}

	fmt.Println("正在为所有linux服务器关闭SELinux")
	LoopExec(setting.LinuxServer, disableSELinux)

	fmt.Println("正在为所有linux服务器关闭firewalld服务")
	LoopExec(setting.LinuxServer, disableFirewalld)

	fmt.Println("正在为所有linux服务器卸载swap")
	LoopExec(setting.LinuxServer, disableSwap)

	fmt.Println("正在为所有linux服务器配置chrony服务")
	LoopExec(setting.LinuxServer, fmt.Sprintf("sh -x %sEnableChrony.sh", myConst.InitScriptDir))

	fmt.Println("正在为k8s集群节点linux服务器优化内核")
	LoopExec(setting.K8sServer, fmt.Sprintf("sh -x %sKernelOptimize.sh", myConst.InitScriptDir))

	fmt.Println("正在为k8s集群节点安装基础软件")
	LoopExec(setting.K8sServer, softwareInstall)

	fmt.Println("正在为k8s集群节点启用iptables")
	LoopExec(setting.K8sServer, enableIptables)

	fmt.Println("正在为k8s集群节点开启ipvs")
	LoopExec(setting.K8sServer, fmt.Sprintf("sh -x %sEnableIpvs.sh", myConst.InitScriptDir))

	fmt.Println("正在为k8s集群节点安装docker")
	LoopExec(setting.K8sServer, fmt.Sprintf("sh -x %sDockerInstall.sh", myConst.InitScriptDir))

}
