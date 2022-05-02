package config

import (
	myConst "erbiaoOS/const"
	"os"
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
	ClusterApiserverEndpoint   = "127.0.0.1:16443"
	K8sMasterHost, K8sNodeHost = hostInfo()

	K8sClusterHost = append(K8sMasterHost, K8sNodeHost...)
)

func hostInfo() (masterHost, nodeHost []HostInfo) {
	for _, ip := range myConst.K8sMasterIPs {
		masterHost = append(masterHost, HostInfo{
			Role:     "k8sMaster",
			LanIp:    ip,
			User:     myConst.SshUser,
			Password: myConst.SshPassword,
			Port:     myConst.SshPort,
		})
	}

	for _, ip := range myConst.K8sMasterIPs {
		masterHost = append(nodeHost, HostInfo{
			Role:     "k8sNode",
			LanIp:    ip,
			User:     myConst.SshUser,
			Password: myConst.SshPassword,
			Port:     myConst.SshPort,
		})
	}
	return
}

// GetHostInfo 获取对应IP的节点信息
func GetHostInfo(ip string) *HostInfo {

	for _, host := range K8sClusterHost {
		if ip == host.LanIp {
			return &host
		}
	}
	return nil
}

func init() {
	// 初始化每节点临时目录
	for _, ip := range append(myConst.K8sMasterIPs, myConst.K8sNodeIPs...) {
		os.MkdirAll(myConst.TempDir+ip, 0777)
	}
}

// DeployMode 返回k8s的部署模式
func DeployMode() string {
	if len(myConst.K8sMasterIPs) < 2 {
		return "standalone"
	}
	return "cluster"
}

func EnterpointAddr() string {
	if DeployMode() == "standalone" {
		return myConst.K8sMasterIPs[0] + ":6443"
	}
	return ClusterApiserverEndpoint
}
