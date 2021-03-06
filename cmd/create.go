package cmd

import (
	"erbiaoOS/install"
	"erbiaoOS/utils"
	"erbiaoOS/utils/net"
	"erbiaoOS/vars"
	"fmt"
	"github.com/spf13/cobra"
)

var localCmdStringVarP = createCmd.Flags().StringVarP

var localCmdStringSliceVarP = createCmd.Flags().StringSliceVarP

var createShow = `
创建k8s集群
	当指定master节点数 = 1时，判定standalone模式，etcd为单节点
	当指定master节点数 > 1时，判定cluster HA模式
		1、当master节点数为2时，从node节点列表中选出一个节点，组成3节点etcd集群
		2、当master节点数大于3，且为偶数个时，减少一个节点，组成n-1节点的etcd集群
		3、etcd集群节点最多为9个
`

var createExample = fmt.Sprintf(`
创建一个三master的k8s集群
./%s create -m 10.21.8.21 -m 10.21.8.41 -m 10.21.8.48 -n 10.21.8.22 -u root -p password --k8s-pkg /tmp/soft/kubernetes-server-linux-amd64.tar.gz --nginx-pkg /tmp/soft/nginx-1.21.6.tar.gz --etcd-pkg /tmp/soft/etcd-v3.5.2-linux-amd64.tar.gz --docker-pkg /tmp/soft/docker-20.10.8.tgz

`, utils.ProgramName)

var createCmd = &cobra.Command{
	Use:     "create",
	Short:   "创建k8s集群",
	Long:    createShow,
	Run:     runCreate,
	Example: createExample,
}

func init() {
	rootCmd.AddCommand(createCmd)

	localCmdStringSliceVarP(&vars.K8sMasterIPs, "master-ip", "m", []string{}, "[必填]master节点IP")
	localCmdStringSliceVarP(&vars.K8sNodeIPs, "node-ip", "n", []string{}, "[必填]node节点IP")

	localCmdStringVarP(&vars.SshUser, "ssh-user", "u", "root", "master节点和node节点部署使用的账号，必须具备root权限")
	localCmdStringVarP(&vars.SshPort, "ssh-port", "P", "22", "master节点和node节点ssh端口")
	localCmdStringVarP(&vars.SshPassword, "ssh-password", "p", "", "[必填]master节点和node节点部署使用账号的密码")

	localCmdStringVarP(&vars.K8sPkg, "k8s-pkg", "", "https://dl.k8s.io/v1.23.4/kubernetes-server-linux-amd64.tar.gz", "k8s二进制包位置，可以是url地址或文件系统位置(如：./kubernetes-server-linux-amd64.tar.gz)")

	localCmdStringVarP(&vars.NginxPkg, "nginx-pkg", "", "https://nginx.org/download/nginx-1.21.6.tar.gz", "nginx源码位置，可以是url地址或文件系统位置(如：./nginx-1.21.6.tar.gz)")

	localCmdStringVarP(&vars.DockerPkg, "docker-pkg", "", "https://download.docker.com/linux/static/stable/x86_64/docker-20.10.8.tgz", "docker二进制包位置，可以是url地址或文件系统位置(如：./docker-20.10.8.tgz)")

	localCmdStringVarP(&vars.EtcdPkg, "etcd-pkg", "", "https://github.com/etcd-io/etcd/releases/download/v3.5.2/etcd-v3.5.2-linux-amd64.tar.gz", "etcd二进制包位置，可以是url地址或文件系统位置(如：./etcd-v3.5.2-linux-amd64.tar.gz)")

	createCmd.MarkFlagRequired("master-ip")
	createCmd.MarkFlagRequired("node-ip")
	createCmd.MarkFlagRequired("ssh-password")
}

func runCreate(*cobra.Command, []string) {
	net.ParseMultiIPv4(append(vars.K8sMasterIPs, vars.K8sNodeIPs...))
	install.K8sClusterInit()
}
