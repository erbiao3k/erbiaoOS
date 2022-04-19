package initialize

import (
	customConst "erbiaoOS/const"
	"erbiaoOS/pkg/file"
	"erbiaoOS/pkg/hostname"
	"strings"
)

// isOdd 判断是奇数
func isOdd(num int) bool {
	if num%2 == 0 {
		return false
	}
	return true
}

// isEven 判断是偶数
func isEven(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}

// EtcdHost 只要有集群高可用的规划，那么：
// 		1、master节点数一定是大于等于2的
//		2、当master节点数为2时，从node节点列表中选出一个节点，组成3节点etcd集群
//		3、当master节点数大于3，且为偶数个时，减少一个节点，组成n-1节点的etcd集群
//		4、etcd集群节点最多为9个
func EtcdHost(masterIPs []string, nodeIPs []string) []string {

	countHost := len(masterIPs)

	// 条件1 + 条件2
	if countHost == 2 {
		masterIPs = append(masterIPs, nodeIPs[0])
	}

	// 条件1 + 条件3
	if countHost > 3 && countHost <= 10 && isEven(countHost) {
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

// EtcdSystemdScript 生成etcd systemd管理脚本
func EtcdSystemdScript(etcdIPs []string) {
	etcdCluster := ""
	for index, ip := range etcdIPs {
		etcdName := hostname.GenerateHostname("etcd", ip)
		etcdCluster = etcdCluster + etcdName + "=" + "https://" + ip + ":2380"
		if len(etcdIPs)-1 != index {
			etcdCluster = etcdCluster + ","
		}
	}
	for _, ip := range etcdIPs {
		currentEtcdName := hostname.GenerateHostname("etcd", ip)
		systemd := strings.ReplaceAll(customConst.EtcdSystemd, "currentEtcdName", currentEtcdName)
		systemd = strings.ReplaceAll(systemd, "currentEtcdIp", ip)
		systemd = strings.ReplaceAll(systemd, "etcdCluster", etcdCluster)
		file.Create(customConst.EtcdSystemdDir+ip+"-etcd.service", systemd)
	}
}

// EtcdCtl 生成etcdctl 客户端管理指令
func EtcdCtl(etcdIPs []string) string {
	etcdServerUrls := ""
	for index, ip := range etcdIPs {
		etcdServerUrls = etcdServerUrls + "https://" + ip + ":2379"
		if len(etcdIPs)-1 != index {
			etcdServerUrls = etcdServerUrls + ","
		}
	}
	file.Create(customConst.EtcdSystemdDir+"etcdctl", strings.ReplaceAll(customConst.EtcdctlManagerCommand, "clientUrls", etcdServerUrls))
	return etcdServerUrls
}