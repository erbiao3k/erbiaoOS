package cert

import (
	"erbiaoOS/pkg/etcd"
	"erbiaoOS/utils/file"
	"erbiaoOS/utils/login/sshd"
	"erbiaoOS/utils/net"
	"erbiaoOS/vars"
)

// certGenerate 生成k8s集群所需所有证书
func certGenerate(masterIPs []string) {

	etcdAltNames := NewAltNames(etcd.ClusterIPs, []string{})

	cer := newCertInfo([]string{"k8s"}, "etcd", etcdAltNames.IPs, etcdAltNames.DNSNames)

	generate(cer, vars.EtcdSslDir+"etcd")

	apiserverClientIPs := append(masterIPs, "10.255.0.1")

	apiserverClientDnsNames := []string{"kubernetes", "kubernetes.default", "kubernetes.default.svc", "kubernetes.default.svc.cluster", "kubernetes.default.svc.cluster.local"}
	kubeApiserverAltNames := NewAltNames(apiserverClientIPs, apiserverClientDnsNames)
	cer = newCertInfo([]string{"k8s"}, "kubernetes", kubeApiserverAltNames.IPs, kubeApiserverAltNames.DNSNames)
	generate(cer, vars.K8sSslDir+"kube-apiserver")

	cer = newCertInfo([]string{"k8s"}, "system:kube-proxy", nil, nil)
	generate(cer, vars.K8sSslDir+"kube-proxy")

	cer = newCertInfo([]string{"system:masters"}, "admin", nil, nil)
	generate(cer, vars.K8sSslDir+"admin")

	ControllerManagerClientIPs := append(masterIPs, "10.255.0.1")
	ControllerManagerAltNames := NewAltNames(ControllerManagerClientIPs, []string{})
	cer = newCertInfo([]string{"system:kube-controller-manager"}, "system:kube-controller-manager", ControllerManagerAltNames.IPs, ControllerManagerAltNames.DNSNames)
	generate(cer, vars.K8sSslDir+"kube-controller-manager")

	kubeSchedulerClientIPs := append(masterIPs, "10.255.0.1")
	kubeSchedulerAltNames := NewAltNames(kubeSchedulerClientIPs, []string{})
	cer = newCertInfo([]string{"system:kube-scheduler"}, "system:kube-scheduler", kubeSchedulerAltNames.IPs, kubeApiserverAltNames.DNSNames)
	generate(cer, vars.K8sSslDir+"kube-scheduler")

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

	for _, host := range etcd.ClusterIPs {
		if host == net.CurrentIP {
			continue
		}
		info := vars.GetHostInfo(host)
		sshd.Upload(info, vars.EtcdSslDir, vars.EtcdSslDir)
	}

}
