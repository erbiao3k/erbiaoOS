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
	os.Mkdir(InitScriptDir, 0777)
	os.Mkdir(K8sMasterBinaryDir, 0777)
	os.Mkdir(K8sNodeBinaryDir, 0777)
	os.Mkdir(DockerTempData, 0777)
}
