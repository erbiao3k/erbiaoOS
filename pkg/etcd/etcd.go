package etcd

import (
	"erbiaoOS/pkg/sysinit"
	"erbiaoOS/utils"
	"erbiaoOS/utils/file"
	"erbiaoOS/utils/num"
	sshd2 "erbiaoOS/utils/sshd"
	"erbiaoOS/vars"
	"strings"
)

// Host 按照逻辑选定etcd部署架构以及节点清单
//只要有集群高可用的规划，那么：
// 		1、master节点数一定是大于等于2的
//		2、当master节点数为2时，从node节点列表中选出一个节点，组成3节点etcd集群
//		3、当master节点数大于3，且为偶数个时，减少一个节点，组成n-1节点的etcd集群
//		4、etcd集群节点最多为9个

func Host(masterIPs []string, nodeIPs []string) []string {

	countHost := len(masterIPs)

	// 条件1 + 条件2
	if countHost == 2 {
		masterIPs = append(masterIPs, nodeIPs[0])
	}

	// 条件1 + 条件3
	if countHost > 3 && countHost <= 10 && num.Even(countHost) {
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
func systemdScript(IPs []string) {
	etcdCluster := ""
	for index, ip := range IPs {
		etcdName := utils.GenerateHostname("etcd", ip)
		etcdCluster = etcdCluster + etcdName + "=" + "https://" + ip + ":2380"
		if len(IPs)-1 != index {
			etcdCluster = etcdCluster + ","
		}
	}
	for _, ip := range IPs {
		currentEtcdName := utils.GenerateHostname("etcd", ip)
		systemd := strings.ReplaceAll(systemd, "currentEtcdName", currentEtcdName)
		systemd = strings.ReplaceAll(systemd, "currentEtcdIp", ip)
		systemd = strings.ReplaceAll(systemd, "etcdCluster", etcdCluster)
		file.Create(vars.TempDir+"/"+ip+"/etcd.service", systemd)
	}
}

// ClientCmd 生成etcdctl 客户端管理指令
func ClientCmd(IPs []string) (etcdServerUrls string) {
	for index, ip := range IPs {
		etcdServerUrls = etcdServerUrls + "https://" + ip + ":2379"
		if len(IPs)-1 != index {
			etcdServerUrls = etcdServerUrls + ","
		}
	}

	var currentUserBashProfile string

	if !file.Exist(BashProfileBak) {
		file.Copy(BashProfileBak, BashProfile)
		currentUserBashProfile = file.Read(BashProfile)
	} else {
		currentUserBashProfile = file.Read(BashProfileBak)
	}

	file.Create(sysinit.BashProfile, currentUserBashProfile+"\n "+strings.ReplaceAll(manageCmd, "clientUrls", etcdServerUrls))
	return etcdServerUrls
}

//Start 初始化etcd服务
func Start() {
	etcdIPs := Host(vars.K8sMasterIPs, vars.K8sNodeIPs)
	systemdScript(etcdIPs)
	ClientCmd(etcdIPs)
	for _, ip := range etcdIPs {
		info := vars.GetHostInfo(ip)

		sshd2.Upload(info, vars.TempDir+"/"+ip+"/etcd.service", vars.SystemdServiceDir)
		sshd2.Upload(info, sysinit.BashProfile, sysinit.SysConfigDir)
		sshd2.RemoteExec(info, etcdRestartCmd)
	}
}
