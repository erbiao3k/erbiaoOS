package vars

const (

	// TempDir 临时目录
	TempDir = "/tmp/"

	// DeployDir 程序部署目录
	DeployDir = "/opt/"

	// InitScriptDir 系统初始化脚本目录
	InitScriptDir = DeployDir + "initScript/"

	// SoftDir 系统初始化脚本目录
	SoftDir = DeployDir + "soft/"

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

	// NginxDir nginx目录
	NginxDir = DeployDir + "nginx/"

	// BinaryDir 二进制目录
	BinaryDir = "/usr/local/bin/"

	// SystemdServiceDir systemd管理脚本目录
	SystemdServiceDir = "/etc/systemd/system/"

	// KubectlConfigDir kubectl 客户端工具配置文件路径
	KubectlConfigDir = "/root/.kube/"
)
