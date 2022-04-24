package main

import (
	"erbiaoOS/pkg/cert"
	"erbiaoOS/pkg/component"
	"erbiaoOS/pkg/etcd"
	"erbiaoOS/pkg/kube_apiserver"
	"erbiaoOS/pkg/kube_controllermanager"
	"erbiaoOS/pkg/kube_proxy"
	"erbiaoOS/pkg/kube_scheduler"
	"erbiaoOS/pkg/kubelet"
	"erbiaoOS/pkg/nginx"
	"erbiaoOS/pkg/sysinit"
	"erbiaoOS/setting"
	"log"
)

func main() {

	log.Println("=+=+=+=+=+=+=+=+=+=+=+=+=+下载k8s必要组件=+=+=+=+=+=+=+=+=+")
	component.Init(setting.ComponentCfg, setting.ClusterHostCfg)

	log.Println("=+=+=+=+=+=+=+=+=+=+=+初始化k8s节点环境=+=+=+=+=+=+=+=+=+=+=+=+=+")
	sysinit.SysInit(setting.ClusterHostCfg)

	log.Println("=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+初始化各组件证书=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+")
	cert.InitCert()

	log.Println("=+=+=+=+=+=+=+==+=+=+=+=+=+=+初始化各组件配置文件=+=+=+=+=+=+=+=+=+=+=+=+=+")
	log.Println("初始化etcd集群")
	etcd.InitEtcd()

	log.Println("初始化kube-apiserver服务")
	kube_apiserver.systemdScript(setting.K8sMasterIPs, etcdServerUrls)

	log.Println("初始化kube-controller-manager服务")
	kube_controllermanager.systemdScript()

	log.Println("初始化kube-scheduler服务")
	kube_scheduler.systemdScript()

	log.Println("初始化kubelet服务")
	kubelet.cfg(customConst.K8sMasterCfgDir, setting.K8sMasterIPs)
	kubelet.cfg(customConst.K8sNodeCfgDir, setting.K8sNodeIPs)

	kubelet.systemdScript(customConst.K8sMasterCfgDir, setting.K8sMasterHost)
	kubelet.systemdScript(customConst.K8sNodeCfgDir, setting.K8sNodeHost)

	log.Println("初始化kube-proxy服务")
	kube_proxy.cfg(customConst.K8sMasterCfgDir, setting.K8sMasterIPs)
	kube_proxy.cfg(customConst.K8sNodeCfgDir, setting.K8sNodeIPs)

	kube_proxy.systemdScript(customConst.K8sMasterCfgDir, setting.K8sMasterHost)
	kube_proxy.systemdScript(customConst.K8sNodeCfgDir, setting.K8sNodeHost)

	log.Println("初始化calico网络组件")
	calico.Cfg()

	log.Println("初始化coreDNS组件")
	coredns.Cfg()

	log.Println("初始化nginx服务")
	nginx.MainCfg(setting.K8sMasterIPs)
	nginx.systemd()

}
