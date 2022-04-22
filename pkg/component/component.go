package component

import (
	customConst "erbiaoOS/const"
	"erbiaoOS/setting"
	"erbiaoOS/utils"
	file2 "erbiaoOS/utils/file"
	"erbiaoOS/utils/login/sshd"
)

// 非k8s核心组件固定版本，不影响使用

const (
	dockerUrl = "https://download.docker.com/linux/static/stable/x86_64/docker-20.10.8.tgz"
	etcdUrl   = "https://github.com/etcd-io/etcd/releases/download/v3.5.2/etcd-v3.5.2-linux-amd64.tar.gz"
)

// Init 初始化组件信息
func Init(component *setting.Component, clusterHost *setting.ClusterHost) {

	k8sMasterHost := clusterHost.K8sMaster
	k8sNodeHost := clusterHost.K8sNode

	utils.Chdir(customConst.SoftDir)

	if !component.OfflineDeployment {

		file2.Download(dockerUrl, ".")
		file2.Download(etcdUrl, ".")
		file2.Download(component.Kubernetes, ".")
	}

	k8sPackage := file2.ListHasPrefix(".", []string{"kubernetes-server"})[0]
	dockerPackage := file2.ListHasPrefix(".", []string{"docker-"})[0]
	etcdPackage := file2.ListHasPrefix(".", []string{"etcd-v"})[0]

	// 解压所有压缩包
	for _, f := range []string{etcdPackage, dockerPackage, k8sPackage} {
		file2.UnTargz(f, ".")
	}

	var etcdBinary = []string{"etcd-v3.5.2-linux-amd64/etcd", "etcd-v3.5.2-linux-amd64/etcdctl", "etcd-v3.5.2-linux-amd64/etcdutl"}
	var dockerBinary = []string{"docker/docker-proxy", "docker/dockerd", "docker/docker", "docker/containerd", "docker/containerd-shim-runc-v2", "docker/ctr", "docker/docker-init", "docker/runc", "docker/containerd-shim"}
	var k8sBinary = []string{"kubernetes/server/bin/kube-apiserver", "kubernetes/server/bin/kube-controller-manager", "kubernetes/server/bin/kubelet", "kubernetes/server/bin/kube-proxy", "kubernetes/server/bin/kube-scheduler"}

	loopExec := func(hosts []setting.HostInfo, binarys []string) {
		for _, h := range hosts {
			for _, b := range binarys {
				sshd.UploadFile(h.LanIp, h.User, h.Password, h.Port, b, customConst.BinaryDir)
			}
		}
	}

	loopExec(k8sMasterHost, etcdBinary)
	loopExec(k8sMasterHost, dockerBinary)
	loopExec(k8sMasterHost, k8sBinary)

	loopExec(k8sNodeHost, dockerBinary)
	loopExec(k8sNodeHost, k8sBinary)

}
