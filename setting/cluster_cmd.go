package setting

import (
	customConst "erbiaoOS/const"
	"os"
)

const (
	configDir = "config"
)

var (
	// ClusterHostCfg 集群初始化节点信息
	ClusterHostCfg = InitclusterHost(configDir)

	ComponentCfg = ComponentContent(configDir)

	K8sMasterIPs, K8sNodeIPs = ipList()
	K8sClusterIPs            = append(K8sMasterIPs, K8sNodeIPs...)

	K8sMasterHost = ClusterHostCfg.K8sMaster
	K8sNodeHost   = ClusterHostCfg.K8sNode
)

// ipList 返回集群节点IP清单
func ipList() (k8sMasterIPs []string, k8sNodeIPs []string) {
	for _, host := range K8sMasterHost {
		k8sMasterIPs = append(k8sMasterIPs, host.LanIp)
	}
	for _, host := range K8sNodeHost {
		k8sNodeIPs = append(k8sNodeIPs, host.LanIp)
	}

	return k8sMasterIPs, k8sNodeIPs
}

// GetHostInfo 获取对应IP的节点信息
func GetHostInfo(ip string) *HostInfo {

	for _, infos := range [][]HostInfo{K8sMasterHost, K8sNodeHost} {
		for _, host := range infos {
			if ip == host.LanIp {
				return &host
			}
		}
	}
	return nil
}

func init() {
	// 初始化每节点临时目录
	for _, ip := range K8sClusterIPs {
		os.MkdirAll(customConst.TempDir+ip, 0777)
	}
}
