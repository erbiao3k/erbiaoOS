package cert

import (
	"erbiaoOS/pkg/etcd"
	"erbiaoOS/utils/file"
	"erbiaoOS/utils/net"
	"erbiaoOS/utils/sshd"
	"erbiaoOS/vars"
)

// certGenerate 生成k8s集群所需所有证书
func certGenerate(masterIPs []string) {

	// 生成etcd证书
	etcdAltNames := NewAltNames(etcd.Host(vars.K8sMasterIPs, vars.K8sNodeIPs), []string{})

	cer := newCertInfo([]string{"k8s"}, "etcd", etcdAltNames.IPs, etcdAltNames.DNSNames)

	generate(cer, vars.EtcdSslDir+"etcd")

	// 生成apiserver证书
	apiserverClientIPs := append(masterIPs, "10.255.0.1")

	apiserverClientDnsNames := []string{"kubernetes", "kubernetes.default", "kubernetes.default.svc", "kubernetes.default.svc.cluster", "kubernetes.default.svc.cluster.local"}
	kubeApiserverAltNames := NewAltNames(apiserverClientIPs, apiserverClientDnsNames)
	cer = newCertInfo([]string{"k8s"}, "kubernetes", kubeApiserverAltNames.IPs, kubeApiserverAltNames.DNSNames)
	generate(cer, vars.K8sSslDir+"kube-apiserver")

	// 生成kube-proxy证书
	cer = newCertInfo([]string{"k8s"}, "system:kube-proxy", nil, nil)
	generate(cer, vars.K8sSslDir+"kube-proxy")

	// 生成kubectl客户端管理证书
	cer = newCertInfo([]string{"system:masters"}, "admin", nil, nil)
	generate(cer, vars.K8sSslDir+"admin")

	// 生成kube-controller-manager客户端管理证书
	ControllerManagerClientIPs := append(masterIPs, "10.255.0.1")
	ControllerManagerAltNames := NewAltNames(ControllerManagerClientIPs, []string{})
	cer = newCertInfo([]string{"system:kube-controller-manager"}, "system:kube-controller-manager", ControllerManagerAltNames.IPs, ControllerManagerAltNames.DNSNames)
	generate(cer, vars.K8sSslDir+"kube-controller-manager")

	// 生成kube-scheduler客户端管理证书
	kubeSchedulerClientIPs := append(masterIPs, "10.255.0.1")
	kubeSchedulerAltNames := NewAltNames(kubeSchedulerClientIPs, []string{})
	cer = newCertInfo([]string{"system:kube-scheduler"}, "system:kube-scheduler", kubeSchedulerAltNames.IPs, kubeApiserverAltNames.DNSNames)
	generate(cer, vars.K8sSslDir+"kube-scheduler")

	// 生成API聚合proxy证书
	cer = newCertInfo([]string{"system:masters"}, "aggregator", nil, nil)
	generate(cer, vars.K8sSslDir+"proxy-client")
}

// InitCert 初始化各节点所需证书
func InitCert() {

	K8sMasterHost, K8sNodeHost, _ := vars.ClusterHostInfo()

	file.Create(CaPrivateKeyFile, caPrivateKey)
	file.Create(CaPubilcKeyFile, caPublicKey)

	certGenerate(vars.K8sMasterIPs)

	for _, host := range K8sMasterHost {
		if host.LanIp == net.CurrentIP {
			continue
		}

		sshd.Upload(&host, vars.K8sSslDir, vars.K8sSslDir)
		sshd.Upload(&host, vars.CaCenterDir, vars.CaCenterDir)
	}

	for _, host := range K8sNodeHost {
		if host.LanIp == net.CurrentIP {
			continue
		}

		sshd.Upload(&host, vars.K8sSslDir, vars.K8sSslDir)
		sshd.Upload(&host, vars.CaCenterDir, vars.CaCenterDir)
	}

	for _, host := range etcd.Host(vars.K8sMasterIPs, vars.K8sNodeIPs) {
		if host == net.CurrentIP {
			continue
		}
		info := vars.GetHostInfo(host)
		sshd.Upload(info, vars.EtcdSslDir, vars.EtcdSslDir)
	}

}
