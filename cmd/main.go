package main

import (
	customConst "Op2K8sDeploy/const"
	"Op2K8sDeploy/install"
	"Op2K8sDeploy/pkg/file"
	"Op2K8sDeploy/setting"
	"log"
)

func main() {

	configDir := "../config"

	log.Println("=+=+=+=+=+=+=+=+=+=+==+=++=+=+初始化配置文件==+=++=+=+=+=+=+=+=+=+=+=+=+=+=+=+")
	clusterHost := setting.InitclusterHost(configDir)
	component := setting.ComponentContent(configDir)

	// 初始化信息
	var k8sMasterHost []string
	for _, master := range clusterHost.K8sMaster {
		k8sMasterHost = append(k8sMasterHost, master.LanIp)
	}

	var k8sNodeHost []string
	for _, master := range clusterHost.K8sNode {
		k8sNodeHost = append(k8sNodeHost, master.LanIp)
	}

	var k8sClusterHost []string
	k8sClusterHost = append(k8sClusterHost, k8sMasterHost...)
	k8sClusterHost = append(k8sClusterHost, k8sNodeHost...)

	log.Println("=+=+=+=+=+=+=+=+=+=+=+创建k8s集群节点环境初始化脚本=+=+=+=+=+=+=+=+=+=+=+=+=+")

	for script, cmd := range customConst.InitScript {
		file.Create(customConst.InitScriptDir+script, cmd)
	}

	log.Println("=+=+=+=+=+=+=+=+=+=+=+=+=+开始下载k8s必要组件,并分拣二进制文件=+=+=+=+=+=+=+=+=+")
	install.ComponentInit(component)

	log.Println("=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+初始化各组件证书=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+")

	log.Println("初始化CA机构证书")
	install.InitCaCert()

	log.Println("初始化ETCD服务证书")
	install.InitEtcdCert(k8sMasterHost)

	log.Println("依据CA机构证书生成kube-apiserver证书")
	install.InitKubeApiserverCert(k8sMasterHost)

	log.Println("依据CA机构证书生成kube-proxy证书")
	install.InitKubeProxyCert()

	log.Println("依据CA机构证书生成kubectl管理证书")
	install.InitKubectlCert()

	log.Println("依据CA机构证书生成kube-controller-manager证书")
	install.InitKubeControllerManagerCert(k8sMasterHost)

	log.Println("依据CA机构证书生成kube-scheduler证书")
	install.InitKubeSchedulerCert(k8sMasterHost)

	log.Println("=+=+=+=+=+=+=+==+=+=+=+=+=+=+=+初始化各组件配置文件=+=+=+=+=+=+=+=+=+=+=+=+=+")
}
