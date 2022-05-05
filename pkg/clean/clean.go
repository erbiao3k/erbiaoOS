package clean

import (
	"erbiaoOS/pkg/sysinit"
	"erbiaoOS/utils"
	"erbiaoOS/utils/sshd"
	"erbiaoOS/vars"
	"fmt"
	"log"
	"os"
)

var (
	// 部署开始前.需要清理的数据
	oldData = " " + vars.EtcdDir + " " + vars.CaCenterDir + " " + vars.K8sSslDir + " "

	// 部署前.需要初始化的目录
	initDirs = vars.DeployDir + "{soft,nginx/{logs,conf,sbin},initScript,caCenter,etcd/{ssl,data},kubernetes/{ssl,cfg}} /var/log/kubernetes"

	// 部署结束后，需要清理数据
	dirtyData = " " + vars.TempDir + "*.tar " + vars.SoftDir + "{etcd,nginx,kubernetes,offlineImage}"

	RoleCmd = "kubectl label node %s node-role.kubernetes.io/%s=true --overwrite=true"

	// master节点不允许调度pod
	noScheduleCmd = "kubectl taint nodes -l node-role.kubernetes.io/master node-role.kubernetes.io/master=true:NoSchedule  --overwrite=true"
)

// PostStart 部署前，清理可能阻塞部署的软件包，进程，历史数据
func PostStart() {
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

// PreStop 结束部署前，清理垃圾，设置节点属性
func PreStop() {
	_, _, K8sClusterHost := vars.ClusterHostInfo()

	for _, host := range K8sClusterHost {
		sshd.RemoteExec(&host, fmt.Sprintf("rm -rf %s", dirtyData))

		hostname := utils.GenerateHostname(host.Role, host.LanIp)

		masterRole := fmt.Sprintf(RoleCmd, hostname, "master")
		nodeRole := fmt.Sprintf(RoleCmd, hostname, "worker")

		// allinone模式下，不设置节点角色
		if len(K8sClusterHost) == 1 {
			return
		}

		if host.Role == vars.MasterRole {
			utils.ExecCmd(masterRole)
			utils.ExecCmd(noScheduleCmd)
		}
		if host.Role == vars.NodeRole {
			utils.ExecCmd(nodeRole)
		}

	}
}
