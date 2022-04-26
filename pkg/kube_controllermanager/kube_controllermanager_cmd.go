package kube_controllermanager

import (
	"bytes"
	myConst "erbiaoOS/const"
	"erbiaoOS/setting"
	"erbiaoOS/utils/file"
	"erbiaoOS/utils/login/sshd"
	"fmt"
	"log"
	"os/exec"
)

// systemdScript 生成kube-controller-manager systemd管理脚本
func systemdScript() {
	file.Create(myConst.TempDir+"kube-controller-manager.service", systemd)
}

// credentials 初始化认证信息
func credentials() {
	cmd := exec.Command("bash", "-c", controllerManagerSetClusterCmd+controllerManagerSetCredentialsCmd+controllerManagerSetContextCmd+controllerManagerUseContextCmd)
	log.Println("----------------------")
	fmt.Println(cmd)
	log.Println("----------------------")

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal("初始化kubectl客户端异常：", err)
	}
}

func InitControllerManagerCluster() {
	systemdScript()
	credentials()

	for _, host := range setting.K8sMasterHost {
		sshd.Upload(host.LanIp, host.User, host.Password, host.Port, myConst.TempDir+"kube-controller-manager.service", myConst.SystemdServiceDir)
		sshd.Upload(host.LanIp, host.User, host.Password, host.Port, controllerManagerKubeConfig, myConst.K8sCfgDir)
		sshd.RemoteSshExec(host.LanIp, host.User, host.Password, host.Port, controllerManagerRestartCmd)
	}
}
