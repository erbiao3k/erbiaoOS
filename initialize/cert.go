package initialize

import (
	customConst "erbiaoOS/const"
	"erbiaoOS/pkg/file"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

var (
	caPrivateKeyfile = customConst.CaCenterDir + "ca-key.pem"
	caPublicKeyfile  = customConst.CaCenterDir + "ca.pem"
	caCsrJsonfile    = customConst.CaCenterDir + "ca-csr.json"
	caConfigJsonfile = customConst.CaCenterDir + "ca-config.json"
)

// slice2String 切片转换为格式化的字符串
func slice2String(host []string) string {
	var str string
	for _, ip := range host {
		str = fmt.Sprintf("    \"%s\",\n", ip) + str
	}
	return str
}

// CaCert 初始化CA机构证书
func CaCert() {

	file.Create(caConfigJsonfile, customConst.CaConfigJson)
	file.Create(caCsrJsonfile, customConst.CaCsrJson)

	cmd := fmt.Sprintf("%s gencert -initca %s | %s -bare %sca -", CfsslBinary(), caCsrJsonfile, CfssljsonBinary(), customConst.CaCenterDir)

	if runtime.GOOS == "windows" {
		cmd = strings.ReplaceAll(cmd, "/", "\\")
		file.Create("createCaCert.bat", cmd)
		exec.Command("createCaCert.bat").Run()
	} else {
		file.Create("createCaCert.sh", cmd)
		exec.Command("bash createCaCert.sh").Run()
	}

}

// EtcdCert 初始化Etcd服务证书
func EtcdCert(masterHost []string) {
	etcdCsrJsonfile := customConst.TempData + "etcd-csr.json"

	EtcdCsrJson := fmt.Sprintf(customConst.EtcdCsrJson, slice2String(masterHost))

	file.Create(etcdCsrJsonfile, EtcdCsrJson)

	cmd := fmt.Sprintf("%s gencert -ca=%s -ca-key=%s -config=%s -profile=kubernetes %s | %s  -bare %s",
		CfsslBinary(), caPublicKeyfile, caPrivateKeyfile, caConfigJsonfile, etcdCsrJsonfile, CfssljsonBinary(), customConst.EtcdSslDir+"etcd")
	if runtime.GOOS == "windows" {
		cmd = strings.ReplaceAll(cmd, "/", "\\")
		file.Create("createEtcdCert.bat", cmd)
		exec.Command("createEtcdCert.bat").Run()
	} else {
		file.Create("createEtcdCert.sh", cmd)
		exec.Command("bash createEtcdCert.sh").Run()
	}
}

// KubeApiserverCert 初始化组件KubeApiser证书
func KubeApiserverCert(masterHost []string) {
	kubeApiserverCsrJsonfile := customConst.TempData + "kube-apiserver-csr.json"
	kubeApiserverCsrJson := fmt.Sprintf(customConst.KubeApiserverCsrJson, slice2String(masterHost))
	file.Create(kubeApiserverCsrJsonfile, kubeApiserverCsrJson)
	cmd := fmt.Sprintf("%s gencert -ca=%s -ca-key=%s -config=%s -profile=kubernetes %s | %s  -bare %s",
		CfsslBinary(), caPublicKeyfile, caPrivateKeyfile, caConfigJsonfile, kubeApiserverCsrJsonfile, CfssljsonBinary(), customConst.KubernetesMasterSslDir+"kube-apiserver")

	if runtime.GOOS == "windows" {
		cmd = strings.ReplaceAll(cmd, "/", "\\")
		file.Create("createApiserverCert.bat", cmd)
		exec.Command("createApiserverCert.bat").Run()
	} else {
		file.Create("createApiserverCert.sh", cmd)
		exec.Command("bash createApiserverCert.sh").Run()
	}
}

// KubeProxyCert 初始化组件KubeProxy证书
func KubeProxyCert() {
	kubeProxyCsrJsonfile := customConst.TempData + "kube-proxy-csr.json"
	file.Create(kubeProxyCsrJsonfile, customConst.KubeProxyCsrJson)
	cmd := fmt.Sprintf("%s gencert -ca=%s -ca-key=%s -config=%s -profile=kubernetes %s | %s  -bare %s",
		CfsslBinary(), caPublicKeyfile, caPrivateKeyfile, caConfigJsonfile, kubeProxyCsrJsonfile, CfssljsonBinary(), customConst.KubernetesClusterSslDir+"kube-proxy")

	if runtime.GOOS == "windows" {
		cmd = strings.ReplaceAll(cmd, "/", "\\")
		file.Create("createKubProxyCert.bat", cmd)
		exec.Command("createKubProxyCert.bat").Run()
	} else {
		file.Create("createKubProxyCert.sh", cmd)
		exec.Command("bash createKubProxyCert.sh").Run()
	}
}

// KubectlCert 初始化Kubectl管理证书
func KubectlCert() {

	kubectlAdminCertfile := customConst.TempData + "admin-csr.json"
	file.Create(kubectlAdminCertfile, customConst.KubectlAdminCert)

	cmd := fmt.Sprintf("%s gencert -ca=%s -ca-key=%s -config=%s -profile=kubernetes %s | %s  -bare %s",
		CfsslBinary(), caPublicKeyfile, caPrivateKeyfile, caConfigJsonfile, kubectlAdminCertfile, CfssljsonBinary(), customConst.KubernetesMasterSslDir+"admin")

	if runtime.GOOS == "windows" {
		cmd = strings.ReplaceAll(cmd, "/", "\\")
		file.Create("createKubctlAdminCert.bat", cmd)
		exec.Command("createKubctlAdminCert.bat").Run()
	} else {
		file.Create("createKubctlAdminCert.sh", cmd)
		exec.Command("bash createKubctlAdminCert.sh").Run()
	}
}

// KubeControllerManagerCert 初始化组件KubeControllerManager证书
func KubeControllerManagerCert(masterHost []string) {
	kubeControllerManagerCsrJson := fmt.Sprintf(customConst.KubeControllerManagerCsrJson, slice2String(masterHost))
	kubeControllerManagerCsrJsonfile := customConst.TempData + "kube-controller-manager-csr.json"
	file.Create(kubeControllerManagerCsrJsonfile, kubeControllerManagerCsrJson)

	cmd := fmt.Sprintf("%s gencert -ca=%s -ca-key=%s -config=%s -profile=kubernetes %s | %s  -bare %s",
		CfsslBinary(), caPublicKeyfile, caPrivateKeyfile, caConfigJsonfile, kubeControllerManagerCsrJsonfile, CfssljsonBinary(), customConst.KubernetesMasterSslDir+"kube-controller-manager")

	if runtime.GOOS == "windows" {
		cmd = strings.ReplaceAll(cmd, "/", "\\")
		file.Create("createkubeControllerManagerCert.bat", cmd)
		exec.Command("createkubeControllerManagerCert.bat").Run()
	} else {
		file.Create("createkubeControllerManagerCert.sh", cmd)
		exec.Command("bash createkubeControllerManagerCert.sh").Run()
	}
}

// KubeSchedulerCert 初始化组件KubeScheduler证书
func KubeSchedulerCert(masterHost []string) {
	kubeSchedulerCsrJson := fmt.Sprintf(customConst.KubeSchedulerCsrJson, slice2String(masterHost))
	kubeSchedulerCsrJsonfile := customConst.TempData + "kube-scheduler-csr.json"
	file.Create(kubeSchedulerCsrJsonfile, kubeSchedulerCsrJson)

	cmd := fmt.Sprintf("%s gencert -ca=%s -ca-key=%s -config=%s -profile=kubernetes %s | %s  -bare %s",
		CfsslBinary(), caPublicKeyfile, caPrivateKeyfile, caConfigJsonfile, kubeSchedulerCsrJsonfile, CfssljsonBinary(), customConst.KubernetesMasterSslDir+"kube-scheduler")

	if runtime.GOOS == "windows" {
		cmd = strings.ReplaceAll(cmd, "/", "\\")
		file.Create("createkubeSchedulerCert.bat", cmd)
		exec.Command("createkubeSchedulerCert.bat").Run()
	} else {
		file.Create("createkubeSchedulerCert.sh", cmd)
		exec.Command("bash createkubeSchedulerCert.sh").Run()
	}
}
