package component

import (
	myConst "erbiaoOS/const"
	"erbiaoOS/pkg/etcd"
	"erbiaoOS/setting"
	"erbiaoOS/utils"
	"erbiaoOS/utils/file"
	"erbiaoOS/utils/login/sshd"
)

// 非k8s核心组件固定版本，不影响使用

const (
	dockerUrl = "https://download.docker.com/linux/static/stable/x86_64/docker-20.10.8.tgz"
	etcdUrl   = "https://github.com/etcd-io/etcd/releases/download/v3.5.2/etcd-v3.5.2-linux-amd64.tar.gz"
	nginxUrl  = "https://nginx.org/download/nginx-1.21.6.tar.gz"
)

// Init 初始化组件信息
func Init() {
	utils.Chdir(myConst.SoftDir)

	if !setting.ComponentCfg.OfflineDeployment {

		file.Download(dockerUrl, "./")
		file.Download(etcdUrl, "./")
		file.Download(nginxUrl, "./")
		file.Download(setting.ComponentCfg.Kubernetes, "./")
	}

	k8sPackage := file.ListHasPrefix("./", []string{"kubernetes-server"})[0]
	dockerPackage := "docker-20.10.8.tgz"
	etcdPackage := "etcd-v3.5.2-linux-amd64.tar.gz"
	nginxPackage := "nginx-1.21.6.tar.gz"

	// 解压所有压缩包
	for _, f := range []string{etcdPackage, dockerPackage, k8sPackage, nginxPackage} {
		file.UnTargz(f, "./")
	}

	var etcdBinary = []string{"etcd-v3.5.2-linux-amd64/etcd", "etcd-v3.5.2-linux-amd64/etcdctl", "etcd-v3.5.2-linux-amd64/etcdutl"}
	var dockerBinary = []string{"docker/docker-proxy", "docker/dockerd", "docker/docker", "docker/containerd", "docker/containerd-shim-runc-v2", "docker/ctr", "docker/docker-init", "docker/runc", "docker/containerd-shim"}
	var k8sBinary = []string{"kubernetes/server/bin/kube-apiserver", "kubernetes/server/bin/kube-controller-manager", "kubernetes/server/bin/kubectl", "kubernetes/server/bin/kubelet", "kubernetes/server/bin/kube-proxy", "kubernetes/server/bin/kube-scheduler"}

	loopExec := func(hosts []setting.HostInfo, binarys []string) {
		for _, host := range hosts {

			for _, b := range binarys {
				sshd.Upload(&host, b, myConst.BinaryDir)
			}

			sshd.RemoteExec(&host, "chmod +x -R "+myConst.BinaryDir)
		}
	}

	for _, ip := range etcd.ClusterIPs {
		host := setting.GetHostInfo(ip)

		for _, binary := range etcdBinary {
			sshd.Upload(host, binary, myConst.BinaryDir)
		}
		sshd.RemoteExec(host, "chmod +x -R "+myConst.BinaryDir)
	}

	loopExec(setting.K8sClusterHost, k8sBinary)
	loopExec(setting.K8sClusterHost, dockerBinary)

}
