package cert

import (
	"erbiaoOS/pkg/etcd"
)

func ClusterCertGenerate(masterIPs, nodeIPs []string) {
	etcdAltNames := NewAltNames(etcd.HostLIst(masterIPs, nodeIPs), []string{})
	cer := newCertInfo([]string{"k8s"}, "etcd", etcdAltNames.IPs, etcdAltNames.DNSNames)
	generate(cer, "etcd")

	apiserverClientIPs := append(masterIPs, "10.255.0.1")
	apiserverClientDnsNames := []string{"kubernetes", "kubernetes.default", "kubernetes.default.svc", "kubernetes.default.svc.cluster", "kubernetes.default.svc.cluster.local"}
	kubeApiserverAltNames := NewAltNames(apiserverClientIPs, apiserverClientDnsNames)
	cer = newCertInfo([]string{"k8s"}, "kubernetes", kubeApiserverAltNames.IPs, kubeApiserverAltNames.DNSNames)
	generate(cer, "kube-apiserver")

	cer = newCertInfo([]string{"k8s"}, "system:kube-proxy", nil, nil)
	generate(cer, "kube-proxy")

	cer = newCertInfo([]string{"system:masters"}, "admin", nil, nil)
	generate(cer, "admin")

	ControllerManagerClientIPs := append(masterIPs, "10.255.0.1")
	ControllerManagerAltNames := NewAltNames(ControllerManagerClientIPs, []string{})
	cer = newCertInfo([]string{"system:kube-controller-manager"}, "system:kube-controller-manager", ControllerManagerAltNames.IPs, ControllerManagerAltNames.DNSNames)
	generate(cer, "kube-controller-manager")

	kubeSchedulerClientIPs := append(masterIPs, "10.255.0.1")
	kubeSchedulerAltNames := NewAltNames(kubeSchedulerClientIPs, []string{})
	cer = newCertInfo([]string{"system:kube-scheduler"}, "system:kube-scheduler", kubeSchedulerAltNames.IPs, kubeApiserverAltNames.DNSNames)
	generate(cer, "kube-scheduler")

}
