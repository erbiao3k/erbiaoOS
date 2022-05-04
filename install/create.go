package install

import (
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
	"erbiaoOS/utils/sshd"
	"erbiaoOS/vars"
	"fmt"
	"log"
	"os"
)

func K8sClusterInit() {

	_, _, K8sClusterHost := vars.ClusterHostInfo()
	// 初始化每节点临时目录
	for _, ip := range append(vars.K8sMasterIPs, vars.K8sNodeIPs...) {
		os.MkdirAll(vars.TempDir+ip, 0777)
	}

	log.Println("清理可能阻塞部署的软件包")
	sshd.LoopRemoteExec(K8sClusterHost, sysinit.RemoveSoft)

	log.Println("清理可能阻塞部署的进程")
	sshd.LoopRemoteExec(K8sClusterHost, sysinit.StopService)

	log.Println("清理可能阻塞部署的历史数据")
	sshd.LoopRemoteExec(K8sClusterHost, fmt.Sprintf("rm -rf %s %s %s", vars.EtcdDir, vars.CaCenterDir, vars.K8sSslDir))

	log.Println("初始化环境目录")
	sshd.LoopRemoteExec(K8sClusterHost, fmt.Sprintf("mkdir -p %s %s %s %s %s %s %s %s %s  ", vars.NginxDir+"/{logs,conf,sbin}", vars.InitScriptDir, vars.CaCenterDir, vars.EtcdSslDir, vars.EtcdDataDir, vars.K8sSslDir, vars.K8sCfgDir, vars.KubectlConfigDir, vars.KubernetesLogDir))

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
