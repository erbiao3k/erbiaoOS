package main

import (
	customConst "erbiaoOS/const"
	"erbiaoOS/initialize"
	"erbiaoOS/setting"
	"log"
)

func main() {

	configDir := "config"

	clusterHost := setting.InitclusterHost(configDir)
	component := setting.ComponentContent(configDir)

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
	customConst.InitScript()

	log.Println("=+=+=+=+=+=+=+=+=+=+=+=+=+开始下载k8s必要组件,并分拣二进制文件=+=+=+=+=+=+=+=+=+")
	initialize.ComponentInit(component)

	log.Println("=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+初始化各组件证书=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+")
	initialize.ClusterCertGenerate(k8sMasterHost, k8sNodeHost)

	log.Println("=+=+=+=+=+=+=+==+=+=+=+=+=+=+初始化各组件配置文件=+=+=+=+=+=+=+=+=+=+=+=+=+")
	log.Println("初始化etcd systemd管理脚本，etcdctl客户端指令")
	etcdHost := initialize.EtcdHost(k8sMasterHost, k8sNodeHost)
	initialize.EtcdSystemdScript(etcdHost)
	etcdServerUrls := initialize.EtcdCtl(etcdHost)

	log.Println("初始化kube-apiserver systemd管理脚本")
	initialize.KubeApiserverSystemdScript(k8sMasterHost, etcdServerUrls)

	log.Println("初始化kube-controller-manager systemd管理脚本")
	initialize.KubeControllerManagerSystemdScript()

	log.Println("初始化kube-scheduler systemd管理脚本")
	initialize.KubeSchedulerSystemdScript()

	log.Println("初始化kubelet配置文件，以及kubelet systemd管理脚本")
	initialize.KubeletCfg(customConst.K8sMasterCfgDir, k8sMasterHost)
	initialize.KubeletCfg(customConst.K8sNodeCfgDir, k8sNodeHost)

	initialize.KubeletSystemdScript(customConst.K8sMasterCfgDir, clusterHost.K8sMaster)
	initialize.KubeletSystemdScript(customConst.K8sNodeCfgDir, clusterHost.K8sNode)

	log.Println("初始化kube-proxy配置文件，以及kube-proxy systemd管理脚本")
	initialize.KubeProxyCfg(customConst.K8sMasterCfgDir, k8sMasterHost)
	initialize.KubeProxyCfg(customConst.K8sNodeCfgDir, k8sNodeHost)

	initialize.KubeProxySystemdScript(customConst.K8sMasterCfgDir, clusterHost.K8sMaster)
	initialize.KubeProxySystemdScript(customConst.K8sNodeCfgDir, clusterHost.K8sNode)

	log.Println("初始化calico网络组件编排文件")
	initialize.CalicoCfg()

	log.Println("初始化coreDNS组件编排文件")
	initialize.CorednsCfg()

	log.Println("初始化nginx主配置文件，以及nginx systemd管理脚本")
	initialize.NginxMainCfg(k8sMasterHost)
	initialize.NginxSystemd()

}
