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

// EtcdSystemd 只要有集群高可用的规划，那么：
// 		1、master节点数一定是大于等于2的
//		2、当master节点数为2时，从node节点列表中选出一个节点，组成3节点etcd集群
//		3、当master节点数大于3，且为偶数个时，减少一个节点，组成n-1节点的etcd集群
//		4、etcd集群节点最多为9个
func EtcdSystemd(masterHost []string, nodeHost []string) {

	countHost := len(masterHost)

	// 条件1 + 条件2
	if countHost == 2 {
		masterHost = append(masterHost, nodeHost[0])
	}

	// 条件1 + 条件3
	if countHost > 3 && isEven(countHost) {
		masterHost = append(masterHost[:0], masterHost[1:]...)
	}

	etcdCluster := ""
	for index, ip := range masterHost {
		etcdName := hostname.GenerateHostname("etcd", ip)
		etcdCluster = etcdCluster + etcdName + "=" + "https://" + ip + ":2380"
		if countHost-1 != index {
			etcdCluster = etcdCluster + ","
		}
		// 条件1 + 条件4
		if index == 9 {
			break
		}
	}

	for _, ip := range masterHost {
		currentEtcdName := hostname.GenerateHostname("etcd", ip)
		systemd := strings.ReplaceAll(customConst.EtcdSystemd, "currentEtcdName", currentEtcdName)
		systemd = strings.ReplaceAll(systemd, "currentEtcdIp", ip)
		systemd = strings.ReplaceAll(systemd, "etcdCluster", etcdCluster)
		file.Create(customConst.TempData+ip+"-etcd.service", systemd)
	}

}
