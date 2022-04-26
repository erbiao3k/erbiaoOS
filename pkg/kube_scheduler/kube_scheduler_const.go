package kube_scheduler

import (
	myConst "erbiaoOS/const"
	"erbiaoOS/pkg/cert"
	"erbiaoOS/setting"
	"fmt"
)

const (
	// systemd kube-controller-manager服务systemd管理脚本
	systemd = "[Unit]\n" +
		"Description=Kubernetes Scheduler\n" +
		"Documentation=https://github.com/kubernetes/kubernetes\n\n" +
		"[Service]\n" +
		"ExecStart=/usr/local/bin/kube-scheduler \\\n" +
		"--address=127.0.0.1 \\\n" +
		"--kubeconfig=/opt/kubernetes/cfg/kube-scheduler.kubeconfig \\\n" +
		"--leader-elect=true \\\n" +
		"--alsologtostderr=true \\\n" +
		"--logtostderr=false \\\n" +
		"--log-dir=/var/log/kubernetes \\\n" +
		"--v=2\n" +
		"Restart=on-failure\n\n" +
		"[Install]\n" +
		"WantedBy=multi-user.target"

	schedulerUser           = "system:kube-scheduler"
	schedulerContext        = schedulerUser
	schedulerKubeConfig     = myConst.TempDir + "kube-scheduler.kubeconfig"
	schedulerPublicKeyFile  = myConst.K8sSslDir + "kube-scheduler.pem"
	schedulerPrivateKeyFile = myConst.K8sSslDir + "kube-scheduler-key.pem"

	schedulerRestartCmd = "systemctl daemon-reload && systemctl enable kube-scheduler && systemctl restart kube-scheduler && sleep 1"
)

var (
	schedulerSetClusterCmd     = fmt.Sprintf(myConst.SetClusterCmd, cert.CaPubilcKeyFile, setting.RandMasterIP, schedulerKubeConfig)
	schedulerSetCredentialsCmd = fmt.Sprintf(myConst.SetCredentialsCmd, schedulerUser, schedulerPublicKeyFile, schedulerPrivateKeyFile, schedulerKubeConfig)
	schedulerSetContextCmd     = fmt.Sprintf(myConst.SetContextCmd, schedulerContext, schedulerUser, schedulerKubeConfig)
	schedulerUseContextCmd     = fmt.Sprintf(myConst.UseContextCmd, schedulerContext, schedulerKubeConfig)
)
