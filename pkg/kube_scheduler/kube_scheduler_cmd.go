package kube_scheduler

import (
	"bytes"
	myConst "erbiaoOS/const"
	"erbiaoOS/setting"
	"erbiaoOS/utils/file"
	"erbiaoOS/utils/login/sshd"
	"log"
	"os/exec"
)

// systemdScript 生成kube-scheduler systemd管理脚本
func systemdScript() {
	file.Create(myConst.TempDir+"kube-scheduler.service", systemd)
}

// credentials 初始化kube-scheduler认证信息
func credentials() {
	cmd := exec.Command("bash", "-c", schedulerSetClusterCmd+schedulerSetCredentialsCmd+schedulerSetContextCmd+schedulerUseContextCmd)

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal("初始化kubectl客户端异常：", err)
	}
}

// InitSchedulerCluster 初始化scheduler集群
func InitSchedulerCluster() {
	systemdScript()
	credentials()

	for _, host := range setting.K8sMasterHost {
		sshd.Upload(host.LanIp, host.User, host.Password, host.Port, myConst.TempDir+"kube-scheduler.service", myConst.SystemdServiceDir)
		sshd.Upload(host.LanIp, host.User, host.Password, host.Port, schedulerKubeConfig, myConst.K8sCfgDir)
		sshd.RemoteSshExec(host.LanIp, host.User, host.Password, host.Port, schedulerRestartCmd)
	}
}
