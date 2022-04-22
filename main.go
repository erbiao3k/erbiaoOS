package main

import (
	customConst "erbiaoOS/const"
	"erbiaoOS/pkg/calico"
	"erbiaoOS/pkg/cert"
	"erbiaoOS/pkg/component"
	"erbiaoOS/pkg/coredns"
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

	configDir := "config"

	clusterHostCfg := setting.InitclusterHost(configDir)
	componentCfg := setting.ComponentContent(configDir)

	IPs := func(clusterHost []setting.HostInfo) []string {
		var IPs []string
		for _, host := range clusterHostCfg.K8sMaster {
			IPs = append(IPs, host.LanIp)
		}
		return IPs
	}
	k8sMasterIPs := IPs(clusterHostCfg.K8sMaster)
	k8sNodeIPs := IPs(clusterHostCfg.K8sNode)

	var k8sClusterIPs []string
	k8sClusterIPs = append(k8sClusterIPs, k8sMasterIPs...)
	k8sClusterIPs = append(k8sClusterIPs, k8sNodeIPs...)

	log.Println("=+=+=+=+=+=+=+=+=+=+=+=+=+开始下载k8s必要组件,并分拣二进制文件=+=+=+=+=+=+=+=+=+")
	component.Init(componentCfg)

	log.Println("=+=+=+=+=+=+=+=+=+=+=+开始初始化k8s节点=+=+=+=+=+=+=+=+=+=+=+=+=+")
	sysinit.SysInit(clusterHostCfg)

	log.Println("=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+初始化各组件证书=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+")
	cert.ClusterCertGenerate(k8sMasterIPs, k8sNodeIPs)

	log.Println("=+=+=+=+=+=+=+==+=+=+=+=+=+=+初始化各组件配置文件=+=+=+=+=+=+=+=+=+=+=+=+=+")
	log.Println("初始化etcd集群")
	etcdIPs := etcd.HostLIst(k8sMasterIPs, k8sNodeIPs)
	etcd.systemdScript(etcdIPs)
	etcdServerUrls := etcd.clientCmd(etcdIPs)

	log.Println("初始化kube-apiserver服务")
	kube_apiserver.systemdScript(k8sMasterIPs, etcdServerUrls)

	log.Println("初始化kube-controller-manager服务")
	kube_controllermanager.systemdScript()

	log.Println("初始化kube-scheduler服务")
	kube_scheduler.systemdScript()

	log.Println("初始化kubelet服务")
	kubelet.cfg(customConst.K8sMasterCfgDir, k8sMasterIPs)
	kubelet.cfg(customConst.K8sNodeCfgDir, k8sNodeIPs)

	kubelet.systemdScript(customConst.K8sMasterCfgDir, clusterHostCfg.K8sMaster)
	kubelet.systemdScript(customConst.K8sNodeCfgDir, clusterHostCfg.K8sNode)

	log.Println("初始化kube-proxy服务")
	kube_proxy.cfg(customConst.K8sMasterCfgDir, k8sMasterIPs)
	kube_proxy.cfg(customConst.K8sNodeCfgDir, k8sNodeIPs)

	kube_proxy.systemdScript(customConst.K8sMasterCfgDir, clusterHostCfg.K8sMaster)
	kube_proxy.systemdScript(customConst.K8sNodeCfgDir, clusterHostCfg.K8sNode)

	log.Println("初始化calico网络组件")
	calico.Cfg()

	log.Println("初始化coreDNS组件")
	coredns.Cfg()

	log.Println("初始化nginx服务")
	nginx.MainCfg(k8sMasterIPs)
	nginx.systemd()

}
