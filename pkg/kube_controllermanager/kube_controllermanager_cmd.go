package kube_controllermanager

import (
	"bytes"
	myConst "erbiaoOS/const"
	"erbiaoOS/pkg/cert"
	"erbiaoOS/setting"
	"erbiaoOS/utils/file"
	"erbiaoOS/utils/login/sshd"
	"fmt"
	"log"
	"os/exec"
)

const (
	controllerManagerUser           = "system:kube-controller-manager"
	controllerManagerContext        = controllerManagerUser
	controllerManagerKubeConfig     = myConst.TempDir + "kube-controller-manager.kubeconfig"
	controllerManagerPublicKeyFile  = myConst.K8sSslDir + "kube-controller-manager.pem"
	controllerManagerPrivateKeyFile = myConst.K8sSslDir + "kube-controller-manager-key.pem"
)

var (
	controllerManagerSetClusterCmd     = fmt.Sprintf(myConst.SetClusterCmd, cert.CaPubilcKeyFile, setting.RandMasterIP, controllerManagerKubeConfig)
	controllerManagerSetCredentialsCmd = fmt.Sprintf(myConst.SetCredentialsCmd, controllerManagerUser, controllerManagerPublicKeyFile, controllerManagerPrivateKeyFile, controllerManagerKubeConfig)
	controllerManagerSetContextCmd     = fmt.Sprintf(myConst.SetContextCmd, controllerManagerContext, controllerManagerUser, controllerManagerKubeConfig)
	controllerManagerUseContextCmd     = fmt.Sprintf(myConst.UseContextCmd, controllerManagerContext, controllerManagerKubeConfig)
)

// systemdScript 生成kube-controller-manager systemd管理脚本
func systemdScript() {
	file.Create(myConst.TempDir+"kube-controller-manager.service", systemd)
}

// controllerManagerCredentials 初始化controllerManager认证信息
func controllerManagerCredentials() {
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
	controllerManagerCredentials()

	for _, host := range setting.K8sMasterHost {
		sshd.Upload(host.LanIp, host.User, host.Password, host.Port, myConst.TempDir+"kube-controller-manager.service", myConst.SystemdServiceDir)
		sshd.Upload(host.LanIp, host.User, host.Password, host.Port, controllerManagerKubeConfig, myConst.K8sCfgDir)
		sshd.RemoteSshExec(host.LanIp, host.User, host.Password, host.Port, controllerManagerRestartCmd)
	}
}
