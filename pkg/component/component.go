package component

import (
	"erbiaoOS/pkg/etcd"
	"erbiaoOS/utils/file"
	sshd2 "erbiaoOS/utils/sshd"
	"erbiaoOS/vars"
	"fmt"
	"strings"
)

// Init 初始化组件信息
func Init() {

	_, _, K8sClusterHost := vars.ClusterHostInfo()

	file.Chdir(vars.SoftDir)

	// 下载并解压所有压缩包
	temp := func(pkgs []string) {
		for _, pkg := range pkgs {
			if strings.HasPrefix(pkg, "https://") || strings.HasPrefix(pkg, "http://") {
				pkg = file.Download(pkg, "./")
			}
			file.UnTargz(pkg, "./")
		}
	}

	temp([]string{vars.K8sPkg, vars.NginxPkg, vars.EtcdPkg, vars.DockerPkg})

	etcdDir := file.ListHasPrefix(vars.SoftDir, []string{"etcd-v", "."}, file.DIR)[0]
	var etcdBinary = []string{"%s/etcd", "%s/etcdctl", "%s/etcdutl"}
	var dockerBinary = []string{"docker/docker-proxy", "docker/dockerd", "docker/docker", "docker/containerd", "docker/containerd-shim-runc-v2", "docker/ctr", "docker/docker-init", "docker/runc", "docker/containerd-shim"}
	var k8sBinary = []string{"kubernetes/server/bin/kube-apiserver", "kubernetes/server/bin/kube-controller-manager", "kubernetes/server/bin/kubectl", "kubernetes/server/bin/kubelet", "kubernetes/server/bin/kube-proxy", "kubernetes/server/bin/kube-scheduler"}

	loopExec := func(hosts []vars.HostInfo, binarys []string) {
		for _, host := range hosts {
			for _, b := range binarys {
				sshd2.Upload(&host, b, vars.BinaryDir)
			}
			sshd2.RemoteExec(&host, "chmod +x -R "+vars.BinaryDir)
		}
	}

	for _, ip := range etcd.Host(vars.K8sMasterIPs, vars.K8sNodeIPs) {
		host := vars.GetHostInfo(ip)

		for _, binary := range etcdBinary {
			sshd2.Upload(host, fmt.Sprintf(binary, etcdDir), vars.BinaryDir)
		}
		sshd2.RemoteExec(host, "chmod +x -R "+vars.BinaryDir)
	}

	loopExec(K8sClusterHost, k8sBinary)
	loopExec(K8sClusterHost, dockerBinary)

}
