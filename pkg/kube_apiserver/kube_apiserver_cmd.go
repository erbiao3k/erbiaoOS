package kube_apiserver

import (
	myConst "erbiaoOS/const"
	"erbiaoOS/pkg/etcd"
	"erbiaoOS/setting"
	"erbiaoOS/utils"
	"erbiaoOS/utils/file"
	"erbiaoOS/utils/login/sshd"
	"fmt"
	"strconv"
	"strings"
)

// systemdScript 生成kube-apiserver systemd管理脚本
func systemdScript() {
	apiserverCount := len(setting.K8sMasterIPs)
	for _, ip := range setting.K8sMasterIPs {
		cfg := strings.ReplaceAll(systemd, "currentIPaddr", ip)
		cfg = strings.ReplaceAll(cfg, "etcdServerUrls", etcd.EtcdServerUrls)
		cfg = strings.ReplaceAll(cfg, "apiserverCount", strconv.Itoa(apiserverCount))
		file.Create(myConst.TempDir+ip+"/kube-apiserver.service", cfg)
	}
}

// bootstrapToken 生成集群启动引导令牌
func bootstrapToken() {
	tokenCsv := fmt.Sprintf("%s,kubelet-bootstrap,10001,\"system:kubelet-bootstrap\"", utils.RandomString)
	file.Create(myConst.TempDir+"token.csv", tokenCsv)
}

func Start() {
	bootstrapToken()
	systemdScript()
	for _, host := range setting.K8sMasterHost {

		hostInfo := &sshd.Info{
			LanIp:    host.LanIp,
			User:     host.User,
			Password: host.Password,
			Port:     host.Port,
		}
		sshd.Upload(hostInfo, myConst.TempDir+"token.csv", myConst.K8sCfgDir)
		sshd.Upload(hostInfo, myConst.TempDir+"/"+host.LanIp+"/kube-apiserver.service", myConst.SystemdServiceDir)
		sshd.RemoteSshExec(hostInfo, apiserverRestartCmd)
	}
}
