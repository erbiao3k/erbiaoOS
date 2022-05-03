package kube_apiserver

import (
	"erbiaoOS/pkg/etcd"
	"erbiaoOS/utils"
	"erbiaoOS/utils/file"
	"erbiaoOS/utils/login/sshd"
	"erbiaoOS/vars"
	"fmt"
	"strconv"
	"strings"
)

// systemdScript 生成kube-apiserver systemd管理脚本
func systemdScript() {
	apiserverCount := len(vars.K8sMasterIPs)
	for _, ip := range vars.K8sMasterIPs {
		cfg := strings.ReplaceAll(systemd, "currentIPaddr", ip)
		cfg = strings.ReplaceAll(cfg, "etcdServerUrls", etcd.ClientCmd(etcd.Host(vars.K8sMasterIPs, vars.K8sNodeIPs)))
		cfg = strings.ReplaceAll(cfg, "apiserverCount", strconv.Itoa(apiserverCount))
		file.Create(vars.TempDir+ip+"/kube-apiserver.service", cfg)
	}
}

// bootstrapToken 生成集群启动引导令牌
func bootstrapToken() {
	tokenCsv := fmt.Sprintf("%s,kubelet-bootstrap,10001,\"system:kubelet-bootstrap\"", utils.RandomString)
	file.Create(vars.TempDir+"token.csv", tokenCsv)
}

func Start() {
	K8sMasterHost, _, _ := vars.ClusterHostInfo()

	bootstrapToken()
	systemdScript()

	for _, host := range K8sMasterHost {

		sshd.Upload(&host, vars.TempDir+"token.csv", vars.K8sCfgDir)
		sshd.Upload(&host, vars.TempDir+"/"+host.LanIp+"/kube-apiserver.service", vars.SystemdServiceDir)
		sshd.RemoteExec(&host, apiserverRestartCmd)
	}
}
