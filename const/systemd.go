package customConst

const (
	// EtcdService etcd服务systemd管理脚本
	EtcdService = "[Unit]\n" +
		"Description=Etcd Server\n" +
		"After=network.target\n" +
		"After=network-online.target\n" +
		"Wants=network-online.target\n\n" +
		"[Service]\n" +
		"Type=simple\n" +
		"ExecStart=/usr/local/bin/etcd \\\n" +
		"--name=\\${ETCD_NAME} \\\n" +
		"--enable-v2=true \\\n" +
		"--data-dir=/opt/etcd/data/default.etcd \\\n" +
		"--listen-peer-urls=https://\\${ETCD_IP}:2380 \\\n" +
		"--listen-client-urls=https://\\${ETCD_IP}:2379,http://127.0.0.1:2379 \\\n" +
		"--advertise-client-urls=https://\\${ETCD_IP}:2379 \\\n" +
		"--initial-advertise-peer-urls=https://\\${ETCD_IP}:2380 \\\n" +
		"--initial-cluster=\\${ETCD_NAME}=https://\\${ETCD_IP}:2380,\\${ETCD_CLUSTER} \\\n" +
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
)
