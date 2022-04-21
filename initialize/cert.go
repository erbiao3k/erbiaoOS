package initialize

import "erbiaoOS/pkg/cert"

func ClusterCertGenerate(masterIPs, nodeIPs []string) {
	etcdAltNames := cert.NewAltNames(EtcdHost(masterIPs, nodeIPs), []string{})
	cer := cert.NewCertInfo([]string{"k8s"}, "etcd", etcdAltNames.IPs, etcdAltNames.DNSNames)
	cert.Generate(cer, "etcd")

	apiserverClientIPs := append(masterIPs, "10.255.0.1")
	apiserverClientDnsNames := []string{"kubernetes", "kubernetes.default", "kubernetes.default.svc", "kubernetes.default.svc.cluster", "kubernetes.default.svc.cluster.local"}
	kubeApiserverAltNames := cert.NewAltNames(apiserverClientIPs, apiserverClientDnsNames)
	cer = cert.NewCertInfo([]string{"k8s"}, "kubernetes", kubeApiserverAltNames.IPs, kubeApiserverAltNames.DNSNames)
	cert.Generate(cer, "kube-apiserver")

	cer = cert.NewCertInfo([]string{"k8s"}, "system:kube-proxy", nil, nil)
	cert.Generate(cer, "kube-proxy")

	cer = cert.NewCertInfo([]string{"system:masters"}, "admin", nil, nil)
	cert.Generate(cer, "admin")

	ControllerManagerClientIPs := append(masterIPs, "10.255.0.1")
	ControllerManagerAltNames := cert.NewAltNames(ControllerManagerClientIPs, []string{})
	cer = cert.NewCertInfo([]string{"system:kube-controller-manager"}, "system:kube-controller-manager", ControllerManagerAltNames.IPs, ControllerManagerAltNames.DNSNames)
	cert.Generate(cer, "kube-controller-manager")

	kubeSchedulerClientIPs := append(masterIPs, "10.255.0.1")
	kubeSchedulerAltNames := cert.NewAltNames(kubeSchedulerClientIPs, []string{})
	cer = cert.NewCertInfo([]string{"system:kube-scheduler"}, "system:kube-scheduler", kubeSchedulerAltNames.IPs, kubeApiserverAltNames.DNSNames)
	cert.Generate(cer, "kube-scheduler")

}
