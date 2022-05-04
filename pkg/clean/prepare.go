package clean

import (
	"erbiaoOS/pkg/sysinit"
	"erbiaoOS/utils/sshd"
	"erbiaoOS/vars"
	"fmt"
	"log"
	"os"
)

var (
	// 部署前需要清理的数据
	oldData = " " + vars.EtcdDir + " " + vars.CaCenterDir + " " + vars.K8sSslDir + " "

	// 需要初始化的目录清
	initDirs = vars.DeployDir + "{soft,nginx/{logs,conf,sbin},initScript,caCenter,etcd/{ssl,data},kubernetes/{ssl,cfg}} /var/log/kubernetes"
)

// Prepare 清理可能阻塞部署的软件包，进程，历史数据
func Prepare() {
	_, _, K8sClusterHost := vars.ClusterHostInfo()

	log.Println("清理可能阻塞部署的软件包")
	sshd.LoopRemoteExec(K8sClusterHost, sysinit.RemoveSoft)

	log.Println("清理可能阻塞部署的进程")
	sshd.LoopRemoteExec(K8sClusterHost, sysinit.StopService)

	log.Println("清理可能阻塞部署的历史数据")
	sshd.LoopRemoteExec(K8sClusterHost, fmt.Sprintf("rm -rf %s ", oldData))

	// 初始化每节点临时目录
	for _, ip := range append(vars.K8sMasterIPs, vars.K8sNodeIPs...) {
		os.MkdirAll(vars.TempDir+ip, 0777)
	}

	log.Println("初始化环境目录")
	sshd.LoopRemoteExec(K8sClusterHost, fmt.Sprintf("mkdir -p %s", initDirs))

}
