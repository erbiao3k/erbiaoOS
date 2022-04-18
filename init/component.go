package init

import (
	customConst "erbiaoOS/const"
	"erbiaoOS/pkg/file"
	"erbiaoOS/setting"
	"fmt"
	"runtime"
)

func ComponentInit(component *setting.Component) {

	if !component.OfflineDeployment {

		CfsslUrl := fmt.Sprintf(customConst.CfsslUrl, runtime.GOOS, runtime.GOARCH)
		CfssljsonUrl := fmt.Sprintf(customConst.CfssljsonUrl, runtime.GOOS, runtime.GOARCH)
		cfsslCertinfoUrl := fmt.Sprintf(customConst.CfsslCertinfoUrl, runtime.GOOS, runtime.GOARCH)

		if runtime.GOOS == "windows" {
			CfsslUrl = CfsslUrl + ".exe"
			CfssljsonUrl = CfssljsonUrl + ".exe"
			cfsslCertinfoUrl = cfsslCertinfoUrl + ".exe"
		}

		file.Download(CfsslUrl, customConst.TempData)
		file.Download(CfssljsonUrl, customConst.TempData)
		file.Download(cfsslCertinfoUrl, customConst.TempData)
		file.Download(component.Kubernetes, customConst.TempData)
		file.Download(component.Docker, customConst.TempData)
		file.Download(component.Etcd, customConst.TempData)
	}

	k8sPackage := file.ListHasPrefix(customConst.TempData, []string{"kubernetes-server"})[0]
	dockerPackage := file.ListHasPrefix(customConst.TempData, []string{"docker-"})[0]
	etcdPackage := file.ListHasPrefix(customConst.TempData, []string{"etcd-v"})[0]

	// 解压所有压缩包
	for _, f := range []string{etcdPackage, dockerPackage, k8sPackage} {
		file.UnTargz(customConst.TempData+f, customConst.TempData)
	}

	var binary = map[string]string{
		customConst.K8sMasterBinaryDir + "cfssl":                   CfsslBinary(),
		customConst.K8sMasterBinaryDir + "cfssljson":               CfssljsonBinary(),
		customConst.K8sMasterBinaryDir + "cfssl-certinfo":          CfsslcertinfoBinary(),
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
		file.Copy(dist, src)
	}

	dockerBinary := file.List(customConst.DockerTempData)
	for _, Binary := range dockerBinary {
		file.Copy(customConst.K8sMasterBinaryDir+Binary, customConst.DockerTempData+Binary)
		file.Copy(customConst.K8sNodeBinaryDir+Binary, customConst.DockerTempData+Binary)
	}
}
