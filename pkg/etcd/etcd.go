package etcd

import (
	myConst "erbiaoOS/const"
	"erbiaoOS/pkg/config"
	"erbiaoOS/pkg/sysinit"
	"erbiaoOS/utils"
	"erbiaoOS/utils/file"
	"erbiaoOS/utils/login/sshd"
	"strings"
)

var (
	ClusterIPs     = hosts(config.K8sMasterIPs, config.K8sNodeIPs)
	EtcdServerUrls = ClientCmd()
)

// hosts 按照逻辑选定etcd部署架构以及节点清单
//只要有集群高可用的规划，那么：
// 		1、master节点数一定是大于等于2的
//		2、当master节点数为2时，从node节点列表中选出一个节点，组成3节点etcd集群
//		3、当master节点数大于3，且为偶数个时，减少一个节点，组成n-1节点的etcd集群
//		4、etcd集群节点最多为9个

func hosts(masterIPs []string, nodeIPs []string) []string {

	countHost := len(masterIPs)

	// 条件1 + 条件2
	if countHost == 2 {
		masterIPs = append(masterIPs, nodeIPs[0])
	}

	// 条件1 + 条件3
	if countHost > 3 && countHost <= 10 && utils.Even(countHost) {
		masterIPs = append(masterIPs[:0], masterIPs[1:]...)
	}

	// 条件1 + 条件4
	if countHost > 10 {
		n := countHost - 9
		for i := 0; i < n; i++ {
			masterIPs = append(masterIPs[:0], masterIPs[1:]...)
		}
	}

	etcdHost := masterIPs
	return etcdHost
}

// systemdScript 生成etcd systemd管理脚本
func systemdScript() {
	etcdCluster := ""
	for index, ip := range ClusterIPs {
		etcdName := utils.GenerateHostname("etcd", ip)
		etcdCluster = etcdCluster + etcdName + "=" + "https://" + ip + ":2380"
		if len(ClusterIPs)-1 != index {
			etcdCluster = etcdCluster + ","
		}
	}
	for _, ip := range ClusterIPs {
		currentEtcdName := utils.GenerateHostname("etcd", ip)
		systemd := strings.ReplaceAll(systemd, "currentEtcdName", currentEtcdName)
		systemd = strings.ReplaceAll(systemd, "currentEtcdIp", ip)
		systemd = strings.ReplaceAll(systemd, "etcdCluster", etcdCluster)
		file.Create(myConst.TempDir+"/"+ip+"/etcd.service", systemd)
	}
}

// ClientCmd 生成etcdctl 客户端管理指令
func ClientCmd() (etcdServerUrls string) {
	for index, ip := range ClusterIPs {
		etcdServerUrls = etcdServerUrls + "https://" + ip + ":2379"
		if len(ClusterIPs)-1 != index {
			etcdServerUrls = etcdServerUrls + ","
		}
	}
	file.Create(sysinit.BashProfile, sysinit.CurrentUserBashProfile+"\n "+strings.ReplaceAll(manageCmd, "clientUrls", etcdServerUrls))
	return etcdServerUrls
}

//Start 初始化etcd服务
func Start() {
	systemdScript()
	ClientCmd()
	for _, ip := range ClusterIPs {
		info := config.GetHostInfo(ip)

		sshd.Upload(info, myConst.TempDir+"/"+ip+"/etcd.service", myConst.SystemdServiceDir)
		sshd.Upload(info, sysinit.BashProfile, sysinit.SysConfigDir)
		sshd.RemoteExec(info, etcdRestartCmd)
	}
}
