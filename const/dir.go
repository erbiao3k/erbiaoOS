package customConst

import (
	"os"
)

const (
	// DeployDir 程序部署目录
	DeployDir = "/opt/"

	// InitScriptDir 系统初始化脚本目录
	InitScriptDir = DeployDir + "initScript/"

	// CaCenterDir CA机构证书目录
	CaCenterDir = DeployDir + "caCenter/"

	// EtcdDir etcd部署目录
	EtcdDir = DeployDir + "etcd/"

	// EtcdSslDir etcd ssl部署目录
	EtcdSslDir = EtcdDir + "ssl/"

	// EtcdDataDir etcd数据目录
	EtcdDataDir = EtcdDir + "data/"

	// K8sDir k8s部署目录
	K8sDir = DeployDir + "kubernetes/"

	// K8sSslDir k8s ssl部署目录
	K8sSslDir = K8sDir + "ssl/"

	// K8sCfgDir k8s配置目录
	K8sCfgDir = K8sDir + "cfg/"

	// BinaryDir 二进制目录
	BinaryDir = "/usr/local/bin/"

	// SystemdServiceDir systemd管理脚本目录
	SystemdServiceDir = "/etc/systemd/system/"
)

func init() {
	os.MkdirAll(DeployDir, 0777)
	os.MkdirAll(CaCenterDir, 0777)
	os.MkdirAll(EtcdDataDir, 0777)
	os.MkdirAll(EtcdSslDir, 0777)
	os.MkdirAll(K8sSslDir, 0777)
	os.MkdirAll(K8sCfgDir, 0777)
	os.MkdirAll(InitScriptDir, 0777)
}
