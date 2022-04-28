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

	user           = "system:kube-scheduler"
	context        = user
	kubeconfig     = myConst.TempDir + "kube-scheduler.kubeconfig"
	publicKeyFile  = myConst.K8sSslDir + "kube-scheduler.pem"
	privateKeyFile = myConst.K8sSslDir + "kube-scheduler-key.pem"

	restartCmd = "systemctl daemon-reload && systemctl enable kube-scheduler && systemctl restart kube-scheduler && sleep 1"
)

var (
	setClusterCmd     = fmt.Sprintf(myConst.SetClusterCmd, cert.CaPubilcKeyFile, setting.KubeApiserverEndpoint, kubeconfig)
	setCredentialsCmd = fmt.Sprintf(myConst.SetCredentialsCmd, user, publicKeyFile, privateKeyFile, kubeconfig)
	setContextCmd     = fmt.Sprintf(myConst.SetContextCmd, context, user, kubeconfig)
	useContextCmd     = fmt.Sprintf(myConst.UseContextCmd, context, kubeconfig)
)
