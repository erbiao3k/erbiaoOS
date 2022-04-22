package customConst

import (
	"os"
)

const (
	// LocalTemp 部署程序临时目录
	LocalTemp = "localTemp/"

	CaCenterDir             = LocalTemp + "caCenter/"
	EtcdSslDir              = LocalTemp + "etcd/ssl/"
	EtcdDataDir             = LocalTemp + "etcd/data/"
	EtcdSystemdDir          = LocalTemp + "etcd/systemd/"
	KubernetesMasterSslDir  = LocalTemp + "kubernetes/masterSsl/"
	KubernetesNodeSslDir    = LocalTemp + "kubernetes/nodeSsl/"
	KubernetesClusterSslDir = LocalTemp + "kubernetes/clusterSsl/"
	InitScriptDir           = LocalTemp + "initScript/"
	K8sMasterCfgDir         = LocalTemp + "k8sMasterCfg/"
	K8sNodeCfgDir           = LocalTemp + "k8sNodeCfg/"
	K8sMasterBinaryDir      = LocalTemp + "k8sMasterBinary/"
	K8sNodeBinaryDir        = LocalTemp + "k8sNodeBinaryDir/"
	DockerTempData          = LocalTemp + "docker/"
)

func init() {
	os.MkdirAll(LocalTemp, 0777)
	os.MkdirAll(CaCenterDir, 0777)
	os.MkdirAll(EtcdDataDir, 0777)
	os.MkdirAll(EtcdSslDir, 0777)
	os.MkdirAll(EtcdSystemdDir, 0777)
	os.MkdirAll(KubernetesMasterSslDir, 0777)
	os.MkdirAll(KubernetesNodeSslDir, 0777)
	os.MkdirAll(KubernetesClusterSslDir, 0777)
	os.MkdirAll(InitScriptDir, 0777)
	os.MkdirAll(K8sMasterCfgDir, 0777)
	os.MkdirAll(K8sNodeCfgDir, 0777)
	os.MkdirAll(K8sMasterBinaryDir, 0777)
	os.MkdirAll(K8sNodeBinaryDir, 0777)
	os.MkdirAll(DockerTempData, 0777)
}
