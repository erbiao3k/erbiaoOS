package component

import (
	customConst "erbiaoOS/const"
	"erbiaoOS/setting"
	file2 "erbiaoOS/utils/file"
)

// 非k8s核心组件固定版本，不影响使用

const (
	dockerUrl = "https://download.docker.com/linux/static/stable/x86_64/docker-20.10.8.tgz"
	etcdUrl   = "https://github.com/etcd-io/etcd/releases/download/v3.5.2/etcd-v3.5.2-linux-amd64.tar.gz"
)

// Init 初始化组件信息
func Init(component *setting.Component) {

	if !component.OfflineDeployment {
		file2.Download(dockerUrl, customConst.DeployDir)
		file2.Download(etcdUrl, customConst.DeployDir)
		file2.Download(component.Kubernetes, customConst.DeployDir)
	}

	k8sPackage := file2.ListHasPrefix(customConst.DeployDir, []string{"kubernetes-server"})[0]
	dockerPackage := file2.ListHasPrefix(customConst.DeployDir, []string{"docker-"})[0]
	etcdPackage := file2.ListHasPrefix(customConst.DeployDir, []string{"etcd-v"})[0]

	// 解压所有压缩包
	for _, f := range []string{etcdPackage, dockerPackage, k8sPackage} {
		file2.UnTargz(customConst.DeployDir+f, customConst.DeployDir)
	}

	var binary = map[string]string{
		customConst.K8sMasterBinaryDir + "etcd":                    customConst.DeployDir + "etcd-v3.5.2-linux-amd64/etcd",
		customConst.K8sMasterBinaryDir + "etcdctl":                 customConst.DeployDir + "etcd-v3.5.2-linux-amd64/etcdctl",
		customConst.K8sMasterBinaryDir + "etcdutl":                 customConst.DeployDir + "etcd-v3.5.2-linux-amd64/etcdutl",
		customConst.K8sMasterBinaryDir + "kube-apiserver":          customConst.DeployDir + "kubernetes/server/bin/kube-apiserver",
		customConst.K8sMasterBinaryDir + "kube-controller-manager": customConst.DeployDir + "kubernetes/server/bin/kube-controller-manager",
		customConst.K8sMasterBinaryDir + "kubelet":                 customConst.DeployDir + "kubernetes/server/bin/kubelet",
		customConst.K8sMasterBinaryDir + "kube-proxy":              customConst.DeployDir + "kubernetes/server/bin/kube-proxy",
		customConst.K8sMasterBinaryDir + "kube-scheduler":          customConst.DeployDir + "kubernetes/server/bin/kube-scheduler",
		customConst.K8sNodeBinaryDir + "kubelet":                   customConst.DeployDir + "kubernetes/server/bin/kubelet",
		customConst.K8sNodeBinaryDir + "kube-proxy":                customConst.DeployDir + "kubernetes/server/bin/kube-proxy",
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
