package kube_proxy

import (
	myConst "erbiaoOS/const"
)

const (
	// cfgContent kube-proxy配置文件
	cfgContent = "apiVersion: kubeproxy.config.k8s.io/v1alpha1\n" +
		"bindAddress: currentKubeproxyIP\n" +
		"clientConnection:\n" +
		"  kubeconfig: /opt/kubernetes/cfg/kube-proxy.kubeconfig\n" +
		"clusterCIDR: 10.0.0.0/16\n" +
		"healthzBindAddress: currentKubeproxyIP:10256\n" +
		"kind: KubeProxyConfiguration\n" +
		"metricsBindAddress: currentKubeproxyIP:10249\n" +
		"mode: \"ipvs\""

	// systemd kube-proxy systemd管理脚本
	systemd = "[Unit]\n" +
		"Description=Kubernetes Kube-Proxy Server\n" +
		"Documentation=https://github.com/kubernetes/kubernetes\n" +
		"After=network.target\n\n" +
		"[Service]\n" +
		"WorkingDirectory=kubeProxyDataDir\n" +
		"ExecStart=/usr/local/bin/kube-proxy \\\n" +
		"  --config=/opt/kubernetes/cfg/kube-proxy \\\n" +
		"  --alsologtostderr=true \\\n" +
		"  --logtostderr=false \\\n" +
		"  --log-dir=/var/log/kubernetes \\\n" +
		"  --v=2\n" +
		"Restart=on-failure\n" +
		"RestartSec=5\n" +
		"LimitNOFILE=65536\n\n" +
		"[Install]\n" +
		"WantedBy=multi-user.target"

	user           = "kube-proxy"
	context        = "default"
	kubeconfig     = myConst.TempDir + "kube-proxy.kubeconfig"
	publicKeyFile  = myConst.K8sSslDir + "kube-proxy.pem"
	privateKeyFile = myConst.K8sSslDir + "kube-proxy-key.pem"

	// restartCmd 重启指令
	restartCmd = "systemctl daemon-reload && systemctl enable kube-proxy && systemctl restart kube-proxy && sleep 1"
)
