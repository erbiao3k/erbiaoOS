package nginx

import (
	"erbiaoOS/utils"
	"erbiaoOS/utils/file"
	"erbiaoOS/utils/login/sshd"
	"erbiaoOS/vars"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	mainConfigFile = vars.TempDir + "nginx.conf"
	nginxSystemd   = vars.TempDir + "nginx.service"
)

// mainConfig 生成nginx主配置文件
func mainConfig() {
	upstreamConf := ""
	for _, ip := range vars.K8sMasterIPs {
		upstreamConf = upstreamConf + "        server " + ip + ":6443 max_fails=3 fail_timeout=30s;\n"
	}

	file.Create(mainConfigFile, strings.ReplaceAll(mainConf, "upstreamConf", upstreamConf))
}

// systemdScript 生成nginx主配置文件
func systemdScript() {
	file.Create(nginxSystemd, systemd)
}

// deploy 部署nginx
func deploy() {
	nginxCodeDir := file.ListHasPrefix(vars.SoftDir, []string{"nginx-", "."}, file.DIR)[0]
	os.Chdir(vars.SoftDir + nginxCodeDir)
	utils.ExecCmd(fmt.Sprintf(nginxBuild, vars.NginxDir))
	file.Copy(vars.TempDir+"nginx", vars.NginxDir+"sbin/nginx")
}

func Start() {
	_, _, K8sClusterHost := vars.ClusterHostInfo()

	if len(vars.K8sMasterIPs) == 1 {
		log.Println("k8s控制平面为单节点，跳过高可用架构部署")
		return
	}
	mainConfig()
	systemdScript()
	deploy()
	for _, host := range K8sClusterHost {

		sshd.Upload(&host, nginxSystemd, vars.SystemdServiceDir)
		sshd.Upload(&host, mainConfigFile, vars.NginxDir+"conf/")
		sshd.Upload(&host, vars.TempDir+"nginx", vars.NginxDir+"sbin/")

		sshd.RemoteExec(&host, restartCmd)
	}
}
