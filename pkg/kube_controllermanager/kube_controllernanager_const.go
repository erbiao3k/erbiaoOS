package kube_controllermanager

import (
	myConst "erbiaoOS/vars"
)

const (
	// systemd kube-controller-manager服务systemd管理脚本
	systemd = "[Unit]\n" +
		"Description=Kubernetes Controller Manager\n" +
		"Documentation=https://github.com/kubernetes/kubernetes\n\n" +
		"[Service]\n" +
		"ExecStart=/usr/local/bin/kube-controller-manager \\\n" +
		"  --secure-port=10257 \\\n" +
		"  --bind-address=127.0.0.1 \\\n" +
		"  --kubeconfig=/opt/kubernetes/cfg/kube-controller-manager.kubeconfig \\\n" +
		"  --service-cluster-ip-range=10.255.0.0/24 \\\n" +
		"  --cluster-name=kubernetes \\\n" +
		"  --cluster-signing-cert-file=/opt/caCenter/ca.pem \\\n" +
		"  --cluster-signing-key-file=/opt/caCenter/ca-key.pem \\\n" +
		"  --allocate-node-cidrs=true \\\n" +
		"  --cluster-cidr=10.0.0.0/16 \\\n" +
		"  --experimental-cluster-signing-duration=876000h \\\n" +
		"  --root-ca-file=/opt/caCenter/ca.pem \\\n" +
		"  --service-account-private-key-file=/opt/caCenter/ca-key.pem \\\n" +
		"  --leader-elect=true \\\n" +
		"  --feature-gates=RotateKubeletServerCertificate=true \\\n" +
		"  --controllers=*,bootstrapsigner,tokencleaner \\\n" +
		"  --horizontal-pod-autoscaler-sync-period=10s \\\n" +
		"  --tls-cert-file=/opt/kubernetes/ssl/kube-controller-manager.pem \\\n" +
		"  --tls-private-key-file=/opt/kubernetes/ssl/kube-controller-manager-key.pem \\\n" +
		"  --use-service-account-credentials=true \\\n" +
		"  --alsologtostderr=true \\\n" +
		"  --logtostderr=false \\\n" +
		"  --log-dir=/var/log/kubernetes \\\n" +
		"  --v=2\n" +
		"Restart=on-failure\n\n" +
		"[Install]\n" +
		"WantedBy=multi-user.target"

	user           = "system:kube-controller-manager"
	context        = user
	kubeconfig     = myConst.TempDir + "kube-controller-manager.kubeconfig"
	publicKeyFile  = myConst.K8sSslDir + "kube-controller-manager.pem"
	privateKeyFile = myConst.K8sSslDir + "kube-controller-manager-key.pem"

	// restartCmd 重启指令
	restartCmd = "systemctl daemon-reload && systemctl enable kube-controller-manager && systemctl restart kube-controller-manager && sleep 1"
)
