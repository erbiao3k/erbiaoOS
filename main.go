package main

import (
	"erbiaoOS/initialize"
	"erbiaoOS/setting"
	"log"
)

func main() {

	configDir := "config"

	log.Println("=+=+=+=+=+=+=+=+=+=+==+=++=+=+初始化配置文件==+=++=+=+=+=+=+=+=+=+=+=+=+=+=+=+")
	clusterHost := setting.InitclusterHost(configDir)
	//component := setting.ComponentContent(configDir)

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
	//
	//log.Println("=+=+=+=+=+=+=+=+=+=+=+创建k8s集群节点环境初始化脚本=+=+=+=+=+=+=+=+=+=+=+=+=+")
	//for script, cmd := range customConst.InitScript {
	//	file.Create(customConst.InitScriptDir+script, cmd)
	//}
	//
	//log.Println("=+=+=+=+=+=+=+=+=+=+=+=+=+开始下载k8s必要组件,并分拣二进制文件=+=+=+=+=+=+=+=+=+")
	//initialize.ComponentInit(component)
	//
	//log.Println("=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+初始化各组件证书=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+")
	//
	//log.Println("初始化CA机构证书")
	//initialize.CaCert()
	//
	//log.Println("依据CA机构证书生成ETCD服务证书")
	//initialize.EtcdCert(k8sMasterHost)
	//
	//log.Println("依据CA机构证书生成kube-apiserver证书")
	//initialize.KubeApiserverCert(k8sMasterHost)
	//
	//log.Println("依据CA机构证书生成kube-proxy证书")
	//initialize.KubeProxyCert()
	//
	//log.Println("依据CA机构证书生成kubectl管理证书")
	//initialize.KubectlCert()
	//
	//log.Println("依据CA机构证书生成kube-controller-manager证书")
	//initialize.KubeControllerManagerCert(k8sMasterHost)
	//
	//log.Println("依据CA机构证书生成kube-scheduler证书")
	//initialize.KubeSchedulerCert(k8sMasterHost)

	log.Println("=+=+=+=+=+=+=+==+=+=+=+=+=+=+初始化各组件配置文件=+=+=+=+=+=+=+=+=+=+=+=+=+")
	log.Println("初始化etcd systemd管理脚本，etcdctl客户端指令")
	etcdHost := initialize.EtcdHost(k8sMasterHost, k8sNodeHost)
	initialize.EtcdSystemdScript(etcdHost)
	etcdServerUrls := initialize.EtcdCtl(etcdHost)

	log.Println("初始化 kube-apiserver配置文件，systemd管理脚本")
	initialize.KubeApiserverCfg(k8sMasterHost, etcdServerUrls)
	initialize.KubeApiserverSystemdScript()
}
