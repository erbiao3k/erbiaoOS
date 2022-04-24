package etcd

import (
	customConst "erbiaoOS/const"
	"erbiaoOS/setting"
	"erbiaoOS/utils"
	"erbiaoOS/utils/file"
	"erbiaoOS/utils/login/sshd"
	"strings"
)

// Hosts 按照逻辑选定etcd部署架构以及节点清单
//只要有集群高可用的规划，那么：
// 		1、master节点数一定是大于等于2的
//		2、当master节点数为2时，从node节点列表中选出一个节点，组成3节点etcd集群
//		3、当master节点数大于3，且为偶数个时，减少一个节点，组成n-1节点的etcd集群
//		4、etcd集群节点最多为9个
func Hosts(masterIPs []string, nodeIPs []string) []string {

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
func systemdScript(etcdIPs []string) {
	etcdCluster := ""
	for index, ip := range etcdIPs {
		etcdName := utils.GenerateHostname("etcd", ip)
		etcdCluster = etcdCluster + etcdName + "=" + "https://" + ip + ":2380"
		if len(etcdIPs)-1 != index {
			etcdCluster = etcdCluster + ","
		}
	}
	for _, ip := range etcdIPs {
		currentEtcdName := utils.GenerateHostname("etcd", ip)
		systemd := strings.ReplaceAll(systemd, "currentEtcdName", currentEtcdName)
		systemd = strings.ReplaceAll(systemd, "currentEtcdIp", ip)
		systemd = strings.ReplaceAll(systemd, "etcdCluster", etcdCluster)
		file.Create(customConst.TempDir+"/"+ip+"/etcd.service", systemd)
	}
}

// ClientCmd 生成etcdctl 客户端管理指令
func ClientCmd(etcdIPs []string) (etcdServerUrls string) {
	for index, ip := range etcdIPs {
		etcdServerUrls = etcdServerUrls + "https://" + ip + ":2379"
		if len(etcdIPs)-1 != index {
			etcdServerUrls = etcdServerUrls + ","
		}
	}
	file.Create(customConst.TempDir+"etcdctl", strings.ReplaceAll(manageCommand, "clientUrls", etcdServerUrls))
	return etcdServerUrls
}

//InitEtcd 初始化etcd服务
func InitEtcd() {
	etcdIPs := Hosts(setting.K8sMasterIPs, setting.K8sNodeIPs)
	systemdScript(etcdIPs)
	ClientCmd(etcdIPs)
	for _, ip := range etcdIPs {
		hostInfo := setting.GetHostInfo(ip)
		sshd.UploadFile(hostInfo.LanIp, hostInfo.User, hostInfo.Password, hostInfo.Port, customConst.TempDir+"/"+ip+"/etcd.service", customConst.SystemdServiceDir)
	}
}
