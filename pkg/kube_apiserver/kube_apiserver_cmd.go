package kube_apiserver

import (
	myConst "erbiaoOS/const"
	"erbiaoOS/pkg/config"
	"erbiaoOS/pkg/etcd"
	"erbiaoOS/utils"
	"erbiaoOS/utils/file"
	"erbiaoOS/utils/login/sshd"
	"fmt"
	"strconv"
	"strings"
)

// systemdScript 生成kube-apiserver systemd管理脚本
func systemdScript() {
	apiserverCount := len(config.K8sMasterIPs)
	for _, ip := range config.K8sMasterIPs {
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
	for _, host := range config.K8sMasterHost {

		sshd.Upload(&host, myConst.TempDir+"token.csv", myConst.K8sCfgDir)
		sshd.Upload(&host, myConst.TempDir+"/"+host.LanIp+"/kube-apiserver.service", myConst.SystemdServiceDir)
		sshd.RemoteExec(&host, apiserverRestartCmd)
	}
}
