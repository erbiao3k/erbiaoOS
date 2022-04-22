package etcd

const (
	// systemd etcd服务systemd管理脚本
	systemd = "[Unit]\n" +
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

	// manageCommand 生成etcdctl指令
	manageCommand = "alias etcdctl3='ETCDCTL_API=3 etcdctl --cacert=/opt/caCenter/ca.pem --cert=/opt/etcd/ssl/etcd.pem --key=/opt/etcd/ssl/etcd-key.pem --endpoints=clientUrls'\n" +
		"alias etcdctl2='ETCDCTL_API=2 etcdctl --ca-file=/opt/caCenter/ca.pem --cert-file=/opt/etcd/ssl/etcd.pem --key-file=/opt/etcd/ssl/etcd-key.pem --endpoints=clientUrls'"
)
