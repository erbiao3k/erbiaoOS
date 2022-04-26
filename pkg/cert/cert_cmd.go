package cert

import (
	myConst "erbiaoOS/const"
	"erbiaoOS/pkg/etcd"
	"erbiaoOS/setting"
	"erbiaoOS/utils"
	"erbiaoOS/utils/login/sshd"
)

// certGenerate 生成k8s集群所需所有证书
func certGenerate(masterIPs []string) {

	etcdAltNames := NewAltNames(etcd.ClusterIPs, []string{})

	cer := newCertInfo([]string{"k8s"}, "etcd", etcdAltNames.IPs, etcdAltNames.DNSNames)

	generate(cer, myConst.EtcdSslDir+"etcd")

	apiserverClientIPs := append(masterIPs, "10.255.0.1")

	apiserverClientDnsNames := []string{"kubernetes", "kubernetes.default", "kubernetes.default.svc", "kubernetes.default.svc.cluster", "kubernetes.default.svc.cluster.local"}
	kubeApiserverAltNames := NewAltNames(apiserverClientIPs, apiserverClientDnsNames)
	cer = newCertInfo([]string{"k8s"}, "kubernetes", kubeApiserverAltNames.IPs, kubeApiserverAltNames.DNSNames)
	generate(cer, myConst.K8sSslDir+"kube-apiserver")

	cer = newCertInfo([]string{"k8s"}, "system:kube-proxy", nil, nil)
	generate(cer, myConst.K8sSslDir+"kube-proxy")

	cer = newCertInfo([]string{"system:masters"}, "admin", nil, nil)
	generate(cer, myConst.K8sSslDir+"admin")

	ControllerManagerClientIPs := append(masterIPs, "10.255.0.1")
	ControllerManagerAltNames := NewAltNames(ControllerManagerClientIPs, []string{})
	cer = newCertInfo([]string{"system:kube-controller-manager"}, "system:kube-controller-manager", ControllerManagerAltNames.IPs, ControllerManagerAltNames.DNSNames)
	generate(cer, myConst.K8sSslDir+"kube-controller-manager")

	kubeSchedulerClientIPs := append(masterIPs, "10.255.0.1")
	kubeSchedulerAltNames := NewAltNames(kubeSchedulerClientIPs, []string{})
	cer = newCertInfo([]string{"system:kube-scheduler"}, "system:kube-scheduler", kubeSchedulerAltNames.IPs, kubeApiserverAltNames.DNSNames)
	generate(cer, myConst.K8sSslDir+"kube-scheduler")

}

// InitCert 初始化各节点所需证书
func InitCert() {
	certGenerate(setting.K8sMasterIPs)

	for _, host := range setting.K8sMasterHost {
		if host.LanIp == utils.CurrentIP {
			continue
		}
		sshd.Upload(host.LanIp, host.User, host.Password, host.Port, myConst.K8sSslDir, myConst.K8sSslDir)
		sshd.Upload(host.LanIp, host.User, host.Password, host.Port, myConst.CaCenterDir, myConst.CaCenterDir)
	}

	for _, host := range setting.K8sNodeHost {
		if host.LanIp == utils.CurrentIP {
			continue
		}
		sshd.Upload(host.LanIp, host.User, host.Password, host.Port, myConst.K8sSslDir, myConst.K8sSslDir)
		sshd.Upload(host.LanIp, host.User, host.Password, host.Port, myConst.CaCenterDir, myConst.CaCenterDir)
	}

	for _, host := range etcd.ClusterIPs {
		if host == utils.CurrentIP {
			continue
		}
		hostInfo := setting.GetHostInfo(host)
		sshd.Upload(hostInfo.LanIp, hostInfo.User, hostInfo.Password, hostInfo.Port, myConst.EtcdSslDir, myConst.EtcdSslDir)
	}

}
