package config

import (
	myConst "erbiaoOS/const"
	"os"
)

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
	ApiserverEnterpoint = enterpointAddr()

	ClusterApiserverEndpoint   = "127.0.0.1:16443"
	K8sMasterHost, K8sNodeHost = hostInfo()

	K8sClusterHost = append(K8sMasterHost, K8sNodeHost...)
)

func hostInfo() (masterHost, nodeHost []HostInfo) {
	for _, ip := range K8sMasterIPs {
		masterHost = append(masterHost, HostInfo{
			Role:     "k8sMaster",
			LanIp:    ip,
			User:     SshUser,
			Password: SshPassword,
			Port:     SshPort,
			DataDir:  "",
		})
	}

	for _, ip := range K8sMasterIPs {
		masterHost = append(nodeHost, HostInfo{
			Role:     "k8sNode",
			LanIp:    ip,
			User:     SshUser,
			Password: SshPassword,
			Port:     SshPort,
			DataDir:  "",
		})
	}
	return
}

// DeployMode 返回k8s的部署模式
func DeployMode() string {
	if len(K8sMasterIPs) < 2 {
		return "standalone"
	}
	return "cluster"
}

func enterpointAddr() string {
	if DeployMode() == "standalone" {
		return K8sMasterIPs[0] + ":6443"
	}
	return ClusterApiserverEndpoint
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
	for _, ip := range append(K8sMasterIPs, K8sNodeIPs...) {
		os.MkdirAll(myConst.TempDir+ip, 0777)
	}
}
