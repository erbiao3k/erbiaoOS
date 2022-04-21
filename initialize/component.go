package initialize

import (
	customConst "erbiaoOS/const"
	"erbiaoOS/setting"
	file2 "erbiaoOS/utils/file"
)

func ComponentInit(component *setting.Component) {

	//if !component.OfflineDeployment {

	//CfsslUrl := fmt.Sprintf(customConst.CfsslUrl, runtime.GOOS, runtime.GOARCH)
	//CfssljsonUrl := fmt.Sprintf(customConst.CfssljsonUrl, runtime.GOOS, runtime.GOARCH)
	//cfsslCertinfoUrl := fmt.Sprintf(customConst.CfsslCertinfoUrl, runtime.GOOS, runtime.GOARCH)
	//
	//if runtime.GOOS == "windows" {
	//	CfsslUrl = CfsslUrl + ".exe"
	//	CfssljsonUrl = CfssljsonUrl + ".exe"
	//	cfsslCertinfoUrl = cfsslCertinfoUrl + ".exe"
	//}

	//file2.Download(CfsslUrl, customConst.TempData)
	//file2.Download(CfssljsonUrl, customConst.TempData)
	//file2.Download(cfsslCertinfoUrl, customConst.TempData)
	file2.Download(component.Kubernetes, customConst.TempData)
	file2.Download(customConst.DockerUrl, customConst.TempData)
	file2.Download(customConst.EtcdUrl, customConst.TempData)
	//}

	k8sPackage := file2.ListHasPrefix(customConst.TempData, []string{"kubernetes-server"})[0]
	dockerPackage := file2.ListHasPrefix(customConst.TempData, []string{"docker-"})[0]
	etcdPackage := file2.ListHasPrefix(customConst.TempData, []string{"etcd-v"})[0]

	// 解压所有压缩包
	for _, f := range []string{etcdPackage, dockerPackage, k8sPackage} {
		file2.UnTargz(customConst.TempData+f, customConst.TempData)
	}

	var binary = map[string]string{
		//customConst.K8sMasterBinaryDir + "cfssl":                   CfsslBinary(),
		//customConst.K8sMasterBinaryDir + "cfssljson":               CfssljsonBinary(),
		//customConst.K8sMasterBinaryDir + "cfssl-certinfo":          CfsslcertinfoBinary(),
		customConst.K8sMasterBinaryDir + "etcd":                    customConst.TempData + "etcd-v3.5.2-linux-amd64/etcd",
		customConst.K8sMasterBinaryDir + "etcdctl":                 customConst.TempData + "etcd-v3.5.2-linux-amd64/etcdctl",
		customConst.K8sMasterBinaryDir + "etcdutl":                 customConst.TempData + "etcd-v3.5.2-linux-amd64/etcdutl",
		customConst.K8sMasterBinaryDir + "kube-apiserver":          customConst.TempData + "kubernetes/server/bin/kube-apiserver",
		customConst.K8sMasterBinaryDir + "kube-controller-manager": customConst.TempData + "kubernetes/server/bin/kube-controller-manager",
		customConst.K8sMasterBinaryDir + "kubelet":                 customConst.TempData + "kubernetes/server/bin/kubelet",
		customConst.K8sMasterBinaryDir + "kube-proxy":              customConst.TempData + "kubernetes/server/bin/kube-proxy",
		customConst.K8sMasterBinaryDir + "kube-scheduler":          customConst.TempData + "kubernetes/server/bin/kube-scheduler",
		customConst.K8sNodeBinaryDir + "kubelet":                   customConst.TempData + "kubernetes/server/bin/kubelet",
		customConst.K8sNodeBinaryDir + "kube-proxy":                customConst.TempData + "kubernetes/server/bin/kube-proxy",
	}

	for dist, src := range binary {
		file2.Copy(dist, src)
	}

	dockerBinary := file2.List(customConst.DockerTempData)
	for _, Binary := range dockerBinary {
		file2.Copy(customConst.K8sMasterBinaryDir+Binary, customConst.DockerTempData+Binary)
		file2.Copy(customConst.K8sNodeBinaryDir+Binary, customConst.DockerTempData+Binary)
	}
}
