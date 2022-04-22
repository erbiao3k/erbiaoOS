package setting

const (
	configDir = "config"
)

var (
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
