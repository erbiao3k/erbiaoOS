package customConst

const (
	// EtcdSystemd etcd服务systemd管理脚本
	EtcdSystemd = "[Unit]\n" +
		"Description=Etcd Server\n" +
		"After=network.target\n" +
		"After=network-online.target\n" +
		"Wants=network-online.target\n\n" +
		"[Service]\n" +
		"Type=simple\n" +
		"ExecStart=/usr/local/bin/etcd \\\n" +
		"--name=currentEtcdName \\\n" +
		"--enable-v2=true \\\n" +
		"--data-dir=/opt/etcd/data/default.etcd \\\n" +
		"--listen-peer-urls=https://currentEtcdIp:2380 \\\n" +
		"--listen-client-urls=https://currentEtcdIp:2379,http://127.0.0.1:2379 \\\n" +
		"--advertise-client-urls=https://currentEtcdIp:2379 \\\n" +
		"--initial-advertise-peer-urls=https://currentEtcdIp:2380 \\\n" +
		"--initial-cluster=etcdCluster \\\n" +
		"--initial-cluster-token=etcd-cluster \\\n" +
		"--initial-cluster-state=new \\\n" +
		"--cert-file=/opt/etcd/ssl/etcd.pem \\\n" +
		"--key-file=/opt/etcd/ssl/etcd-key.pem \\\n" +
		"--peer-cert-file=/opt/etcd/ssl/etcd.pem \\\n" +
		"--peer-key-file=/opt/etcd/ssl/etcd-key.pem \\\n" +
		"--trusted-ca-file=/opt/caCenter/ca.pem \\\n" +
		"--peer-trusted-ca-file=/opt/caCenter/ca.pem\n" +
		"Restart=on-failure\n" +
		"LimitNOFILE=65536\n\n" +
		"[Install]\n" +
		"WantedBy=multi-user.target"

	// KubeApiserverSystemd kube-apiserver服务systemd管理脚本
	KubeApiserverSystemd = "[Unit]\n" +
		"Description=Kubernetes API Server\n" +
		"Documentation=https://github.com/kubernetes/kubernetes\n\n" +
		"After=etcd.service\n" +
		"Wants=etcd.service\n" +
		"[Service]\n" +
		"ExecStart=/usr/local/bin/kube-apiserver \\\n" +
		"  --enable-admission-plugins=NamespaceLifecycle,NodeRestriction,LimitRanger,ServiceAccount,DefaultStorageClass,ResourceQuota \\\n" +
		"  --anonymous-auth=false \\\n" +
		"  --bind-address=currentIPaddr \\\n" +
		"  --secure-port=6443 \\\n" +
		"  --advertise-address=currentIPaddr \\\n" +
		"  --insecure-port=0 \\\n" +
		"  --authorization-mode=Node,RBAC \\\n" +
		"  --runtime-config=api/all=true \\\n" +
		"  --enable-bootstrap-token-auth \\\n" +
		"  --service-cluster-ip-range=10.255.0.0/24 \\\n" +
		"  --token-auth-file=/opt/kubernetes/cfg/token.csv \\\n" +
		"  --service-node-port-range=30000-50000 \\\n" +
		"  --tls-cert-file=/opt/kubernetes/ssl/kube-apiserver.pem  \\\n" +
		"  --tls-private-key-file=/opt/kubernetes/ssl/kube-apiserver-key.pem \\\n" +
		"  --client-ca-file=/opt/caCenter/ca.pem  \\\n" +
		"  --kubelet-client-certificate=/opt/kubernetes/ssl/kube-apiserver.pem \\\n" +
		"  --kubelet-client-key=/opt/kubernetes/ssl/kube-apiserver-key.pem  \\\n" +
		"  --service-account-key-file=/opt/caCenter/ca.pem \\\n" +
		"  --service-account-signing-key-file=/opt/caCenter/ca-key.pem \\\n" +
		"  --service-account-issuer=https://kubernetes.default.svc.cluster.local \\\n" +
		"  --etcd-cafile=/opt/caCenter/ca.pem \\\n" +
		"  --etcd-certfile=/opt/etcd/ssl/etcd.pem \\\n" +
		"  --etcd-keyfile=/opt/etcd/ssl/etcd-key.pem \\\n" +
		"  --etcd-servers=etcdServerUrls \\\n" +
		"  --enable-swagger-ui=true \\\n" +
		"  --allow-privileged=true \\\n" +
		"  --apiserver-count=apiserverCount \\\n" +
		"  --audit-log-maxage=30 \\\n" +
		"  --audit-log-maxbackup=3 \\\n" +
		"  --audit-log-maxsize=100 \\\n" +
		"  --audit-log-path=/var/log/kube-apiserver-audit.log \\\n" +
		"  --event-ttl=1h \\\n" +
		"  --alsologtostderr=true \\\n" +
		"  --logtostderr=false \\\n" +
		"  --log-dir=/var/log/kubernetes \\\n" +
		"  --v=2\n" +
		"Restart=on-failure\n\n" +
		"[Install]\n" +
		"WantedBy=multi-user.target"

	// KubeControllerManagerSystemd kube-controller-manager服务systemd管理脚本
	KubeControllerManagerSystemd = "[Unit]\n" +
		"Description=Kubernetes Controller Manager\n" +
		"Documentation=https://github.com/kubernetes/kubernetes\n\n" +
		"[Service]\n" +
		"ExecStart=/usr/local/bin/kube-controller-manager \\\n" +
		"  --secure-port=10257 \\\n" +
		"  --bind-address=127.0.0.1 \\\n" +
		"  --kubeconfig=/opt/kubernetes/cfg/kube-controller-manager.kubeconfig \\\n" +
		"  --service-cluster-ip-range==10.255.0.0/24 \\\n" +
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

	// KubeSchedulerSystemd kube-controller-manager服务systemd管理脚本
	KubeSchedulerSystemd = "[Unit]\n" +
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

	//KubeletSystemd kubelet systemd管理脚本
	KubeletSystemd = "[Unit]\n" +
		"Description=Kubernetes Kubelet\n" +
		"Documentation=https://github.com/kubernetes/kubernetes\n" +
		"After=docker.service\n" +
		"Requires=docker.service\n" +
		"[Service]\n" +
		"WorkingDirectory=kubeletDataDir\n" +
		"ExecStart=/usr/local/bin/kubelet \\\n" +
		"  --bootstrap-kubeconfig=/opt/kubernetes/cfg/kubelet-bootstrap.kubeconfig \\\n" +
		"  --cert-dir=/opt/kubernetes/ssl \\\n" +
		"  --kubeconfig=/opt/kubernetes/cfg/kubelet.kubeconfig \\\n" +
		"  --config=/opt/kubernetes/cfg/kubelet \\\n" +
		"  --network-plugin=cni \\\n" +
		"  --pod-infra-container-image=registry.aliyuncs.com/google_containers/pause:3.6 \\\n" +
		"  --alsologtostderr=true \\\n" +
		"  --root-dir=kubeletDataDir\n" +
		"  --logtostderr=false \\\n" +
		"  --log-dir=/var/log/kubernetes \\\n" +
		"  --v=2\n" +
		"Restart=on-failure\n" +
		"RestartSec=5\n\n" +
		"[Install]\n" +
		"WantedBy=multi-user.target"
)
