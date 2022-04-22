package kube_apiserver

const (
	// systemd kube-apiserver服务systemd管理脚本
	systemd = "[Unit]\n" +
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
)
