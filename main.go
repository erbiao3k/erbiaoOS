package main

import (
	myConst "erbiaoOS/const"
	"erbiaoOS/pkg/cert"
	"erbiaoOS/pkg/component"
	"erbiaoOS/pkg/etcd"
	"erbiaoOS/pkg/kube_apiserver"
	"erbiaoOS/pkg/kube_controllermanager"
	"erbiaoOS/pkg/kube_scheduler"
	"erbiaoOS/pkg/kubectl"
	"erbiaoOS/pkg/kubelet"
	"erbiaoOS/pkg/sysinit"
	"erbiaoOS/setting"
	"fmt"
	"log"
)

func main() {

	log.Println("【main】清理可能阻塞部署的进程")
	sysinit.LoopExec(setting.K8sClusterHost, sysinit.StopService)

	log.Println("【main】清理可能阻塞部署的历史数据")
	sysinit.LoopExec(setting.K8sClusterHost, fmt.Sprintf("rm -rf %s", myConst.EtcdDataDir))

	log.Println("【main】初始化环境目录")
	sysinit.LoopExec(setting.K8sClusterHost, fmt.Sprintf("mkdir -p %s %s %s %s %s %s %s %s", myConst.InitScriptDir, myConst.CaCenterDir, myConst.EtcdSslDir, myConst.EtcdDataDir, myConst.K8sSslDir, myConst.K8sCfgDir, myConst.KubectlConfigDir, myConst.KubernetesLogDir))

	log.Println("【main】下载k8s必要组件")
	component.Init(setting.ComponentCfg, setting.ClusterHostCfg)

	log.Println("【main】初始化k8s节点环境")
	sysinit.SysInit()

	log.Println("【main】初始化各组件证书")
	cert.InitCert()

	log.Println("【main】初始化etcd集群")
	etcd.Start()

	log.Println("初始化kube-apiserver服务")
	kube_apiserver.Start()

	log.Println("初始化kubectl客户端管理工具")
	kubectl.InitKubectl()

	log.Println("初始化kube-controller-manager服务")
	kube_controllermanager.Start()

	log.Println("初始化kube-scheduler服务")
	kube_scheduler.Start()

	log.Println("初始化kubelet服务")
	kubelet.Start()

	//kubelet.systemdScript(myConst.K8sMasterCfgDir, setting.K8sMasterHost)
	//kubelet.systemdScript(myConst.K8sNodeCfgDir, setting.K8sNodeHost)
	//
	//log.Println("初始化kube-proxy服务")
	//kube_proxy.cfg(myConst.K8sMasterCfgDir, setting.K8sMasterIPs)
	//kube_proxy.cfg(myConst.K8sNodeCfgDir, setting.K8sNodeIPs)
	//
	//kube_proxy.systemdScript(myConst.K8sMasterCfgDir, setting.K8sMasterHost)
	//kube_proxy.systemdScript(myConst.K8sNodeCfgDir, setting.K8sNodeHost)
	//
	//log.Println("初始化calico网络组件")
	//calico.Cfg()
	//
	//log.Println("初始化coreDNS组件")
	//coredns.Cfg()
	//
	//log.Println("初始化nginx服务")
	//nginx.MainCfg(setting.K8sMasterIPs)
	//nginx.systemd()

}
