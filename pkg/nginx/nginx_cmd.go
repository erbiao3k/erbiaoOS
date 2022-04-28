package nginx

import (
	myConst "erbiaoOS/const"
	"erbiaoOS/setting"
	"erbiaoOS/utils"
	"erbiaoOS/utils/file"
	"erbiaoOS/utils/login/sshd"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	mainConfigFile = myConst.TempDir + "nginx.conf"
	nginxSystemd   = myConst.TempDir + "nginx.service"
)

// mainConfig 生成nginx主配置文件
func mainConfig() {
	upstreamConf := ""
	for _, ip := range setting.K8sMasterIPs {
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
	os.Chdir(myConst.SoftDir + "nginx-1.21.6")
	utils.ExecCmd(fmt.Sprintf(nginxBuild, myConst.NginxDir))
	file.Copy(myConst.TempDir+"nginx", myConst.NginxDir+"sbin/nginx")
}

func Start() {
	if len(setting.K8sMasterIPs) == 1 {
		log.Println("k8s控制平面为单节点，跳过高可用架构部署")
		return
	}
	mainConfig()
	systemdScript()
	deploy()
	for _, host := range setting.K8sClusterHost {

		sshd.Upload(&host, nginxSystemd, myConst.SystemdServiceDir)
		sshd.Upload(&host, mainConfigFile, myConst.NginxDir+"conf/")
		sshd.Upload(&host, myConst.TempDir+"nginx", myConst.NginxDir+"sbin/")

		sshd.RemoteExec(&host, restartCmd)
	}
}
