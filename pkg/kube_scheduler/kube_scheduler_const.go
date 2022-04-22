package kube_scheduler

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
		"--v=2" +
		"Restart=on-failure\n\n" +
		"[Install]\n" +
		"WantedBy=multi-user.target"
)
