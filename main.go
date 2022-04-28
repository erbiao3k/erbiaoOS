package main

import (
	myConst "erbiaoOS/const"
	"erbiaoOS/pkg/calico"
	"erbiaoOS/pkg/cert"
	"erbiaoOS/pkg/component"
	"erbiaoOS/pkg/coredns"
	"erbiaoOS/pkg/etcd"
	"erbiaoOS/pkg/kube_apiserver"
	"erbiaoOS/pkg/kube_controllermanager"
	"erbiaoOS/pkg/kube_proxy"
	"erbiaoOS/pkg/kube_scheduler"
	"erbiaoOS/pkg/kubectl"
	"erbiaoOS/pkg/kubelet"
	"erbiaoOS/pkg/nginx"
	"erbiaoOS/pkg/sysinit"
	"erbiaoOS/setting"
	"erbiaoOS/utils/login/sshd"
	"fmt"
	"log"
)

func main() {

	log.Println("【main】清理可能阻塞部署的软件包")
	sshd.LoopRemoteExec(setting.K8sClusterHost, sysinit.RemoveSoft)

	log.Println("【main】清理可能阻塞部署的进程")
	sshd.LoopRemoteExec(setting.K8sClusterHost, sysinit.StopService)

	log.Println("【main】清理可能阻塞部署的历史数据")
	sshd.LoopRemoteExec(setting.K8sClusterHost, fmt.Sprintf("rm -rf %s %s %s", myConst.EtcdDir, myConst.CaCenterDir, myConst.K8sSslDir))

	log.Println("【main】初始化环境目录")
	sshd.LoopRemoteExec(setting.K8sClusterHost, fmt.Sprintf("mkdir -p %s %s %s %s %s %s %s %s %s", myConst.NginxDir+"/{logs,conf,sbin}", myConst.InitScriptDir, myConst.CaCenterDir, myConst.EtcdSslDir, myConst.EtcdDataDir, myConst.K8sSslDir, myConst.K8sCfgDir, myConst.KubectlConfigDir, myConst.KubernetesLogDir))

	log.Println("【main】下载k8s必要组件")
	component.Init()

	log.Println("【main】初始化k8s节点环境")
	sysinit.SysInit()

	log.Println("【main】初始化各组件证书")
	cert.InitCert()

	log.Println("【main】初始化etcd集群")
	etcd.Start()

	log.Println("初始化高可用nginx服务")
	nginx.Start()

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

	log.Println("初始化kube-proxy服务")
	kube_proxy.Start()

	log.Println("初始化calico网络组件")
	calico.Deploy()

	log.Println("初始化coreDNS组件")
	coredns.Deploy()

}

// 节点最大磁盘识别拆分到sysinit
// 系统初始化逻辑拆分与优化
// 资源预留
// 节点labels
//
