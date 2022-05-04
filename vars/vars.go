package vars

var (
	K8sMasterIPs []string
	K8sNodeIPs   []string
	SshUser      string
	SshPassword  string
	SshPort      string
	K8sPkg       string
	EtcdPkg      string
	NginxPkg     string
	DockerPkg    string
)

// ClusterHost 集群节点初始化信息
type ClusterHost struct {
	K8sMaster []HostInfo
	K8sNode   []HostInfo
}

// HostInfo 节点详细信息
type HostInfo struct {
	Role     string
	LanIp    string
	User     string
	Password string
	Port     string
	DataDir  string
}

var (
	ClusterApiserverEndpoint = "127.0.0.1:16443"
)

// GetHostInfo 获取对应IP的节点信息
func GetHostInfo(ip string) *HostInfo {

	_, _, K8sClusterHost := ClusterHostInfo()

	for _, host := range K8sClusterHost {
		if ip == host.LanIp {
			return &host
		}
	}
	return nil
}

// DeployMode 返回k8s的部署模式
func DeployMode() string {
	if len(K8sMasterIPs) < 2 {
		return "standalone"
	}
	return "cluster"
}

func EnterpointAddr() string {
	if DeployMode() == "standalone" {
		return K8sMasterIPs[0] + ":6443"
	}
	return ClusterApiserverEndpoint
}

func ClusterHostInfo() (masterHost, nodeHost, clusterHost []HostInfo) {
	for _, ip := range K8sMasterIPs {
		info := HostInfo{
			Role:     "k8sMaster",
			LanIp:    ip,
			User:     SshUser,
			Password: SshPassword,
			Port:     SshPort,
		}
		masterHost = append(masterHost, info)
	}

	for _, ip := range K8sNodeIPs {
		info := HostInfo{
			Role:     "k8sNode",
			LanIp:    ip,
			User:     SshUser,
			Password: SshPassword,
			Port:     SshPort,
		}
		nodeHost = append(nodeHost, info)
	}

	clusterHost = append(masterHost, nodeHost...)
	return
}
