package nginx

import (
	myConst "erbiaoOS/const"
	"erbiaoOS/setting"
	"erbiaoOS/utils"
	"erbiaoOS/utils/file"
	"erbiaoOS/utils/login/sshd"
	"fmt"
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
	fmt.Println(fmt.Sprintf(nginxBuild, myConst.NginxDir))
	utils.ExecCmd(fmt.Sprintf(nginxBuild, myConst.NginxDir))
	file.Copy(myConst.TempDir+"nginx", myConst.NginxDir+"sbin/nginx")
}

func Start() {
	mainConfig()
	systemdScript()
	deploy()
	for _, host := range append(setting.K8sMasterHost, setting.K8sNodeHost...) {

		sshd.Upload(host.LanIp, host.User, host.Password, host.Port, nginxSystemd, myConst.SystemdServiceDir)
		//sshd.Upload(host.LanIp, host.User, host.Password, host.Port, myConst.NginxDir+"conf", myConst.NginxDir+"conf/")
		sshd.Upload(host.LanIp, host.User, host.Password, host.Port, mainConfigFile, myConst.NginxDir+"conf/")
		sshd.Upload(host.LanIp, host.User, host.Password, host.Port, myConst.TempDir+"nginx", myConst.NginxDir+"sbin/")

		sshd.RemoteSshExec(host.LanIp, host.User, host.Password, host.Port, restartCmd)
	}
}
