package component

import (
	myConst "erbiaoOS/const"
	"erbiaoOS/pkg/config"
	"erbiaoOS/pkg/etcd"
	"erbiaoOS/utils/file"
	"erbiaoOS/utils/login/sshd"
	"strings"
)

// Init 初始化组件信息
func Init() {
	file.Chdir(myConst.SoftDir)

	// 下载并解压所有压缩包
	temp := func(pkgs []string) {
		for _, pkg := range pkgs {
			if strings.HasPrefix(pkg, "https://") || strings.HasPrefix(pkg, "http://") {
				pkg = file.Download(pkg, "./")
			}
			file.UnTargz(pkg, "./")
		}
	}

	temp([]string{myConst.K8sPkg, myConst.NginxPkg, myConst.EtcdPkg, myConst.DockerPkg})

	var etcdBinary = []string{"etcd-v3.5.2-linux-amd64/etcd", "etcd-v3.5.2-linux-amd64/etcdctl", "etcd-v3.5.2-linux-amd64/etcdutl"}
	var dockerBinary = []string{"docker/docker-proxy", "docker/dockerd", "docker/docker", "docker/containerd", "docker/containerd-shim-runc-v2", "docker/ctr", "docker/docker-init", "docker/runc", "docker/containerd-shim"}
	var k8sBinary = []string{"kubernetes/server/bin/kube-apiserver", "kubernetes/server/bin/kube-controller-manager", "kubernetes/server/bin/kubectl", "kubernetes/server/bin/kubelet", "kubernetes/server/bin/kube-proxy", "kubernetes/server/bin/kube-scheduler"}

	loopExec := func(hosts []config.HostInfo, binarys []string) {
		for _, host := range hosts {

			for _, b := range binarys {
				sshd.Upload(&host, b, myConst.BinaryDir)
			}

			sshd.RemoteExec(&host, "chmod +x -R "+myConst.BinaryDir)
		}
	}

	for _, ip := range etcd.ClusterIPs {
		host := config.GetHostInfo(ip)

		for _, binary := range etcdBinary {
			sshd.Upload(host, binary, myConst.BinaryDir)
		}
		sshd.RemoteExec(host, "chmod +x -R "+myConst.BinaryDir)
	}

	loopExec(config.K8sClusterHost, k8sBinary)
	loopExec(config.K8sClusterHost, dockerBinary)

}
