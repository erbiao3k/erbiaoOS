package install

import (
	"erbiaoOS/pkg/calico"
	"erbiaoOS/pkg/cert"
	"erbiaoOS/pkg/clean"
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
	"log"
)

func K8sClusterInit() {

	log.Println("部署前置准备")
	clean.Prepare()

	log.Println("准备k8s组件")
	component.Init()

	log.Println("初始化k8s节点环境")
	sysinit.SysInit()

	log.Println("初始化各组件证书")
	cert.InitCert()

	log.Println("初始化etcd集群")
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
