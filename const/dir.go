package customConst

import (
	"os"
)

const (
	TempData                = "tempData/"
	CaCenterDir             = TempData + "caCenter/"
	EtcdSslDir              = TempData + "etcd/ssl/"
	EtcdDataDir             = TempData + "etcd/data/"
	EtcdSystemdDir          = TempData + "etcd/systemd/"
	KubernetesMasterSslDir  = TempData + "kubernetes/masterSsl/"
	KubernetesNodeSslDir    = TempData + "kubernetes/nodeSsl/"
	KubernetesClusterSslDir = TempData + "kubernetes/clusterSsl/"
	InitScriptDir           = TempData + "initScript/"
	K8sMasterCfgDir         = TempData + "k8sMasterCfg/"
	K8sNodeCfgDir           = TempData + "k8sNodeCfg/"
	K8sMasterBinaryDir      = TempData + "k8sMasterBinary/"
	K8sNodeBinaryDir        = TempData + "k8sNodeBinaryDir/"
	DockerTempData          = TempData + "docker/"
)

func init() {
	os.MkdirAll(TempData, 0777)
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
