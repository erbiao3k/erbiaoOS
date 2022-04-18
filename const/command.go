package customConst

const (
	// SetHostname 设置主机名的字符串
	SetHostname = `echo $(hostname) |grep localhost || hostnamectl set-hostname $1`

	// DisableSELinux 关闭SELinux的字符串
	DisableSELinux = "sed -i 's/\\=enforcing/\\=disabled/' /etc/selinux/config &&  setenforce 0 || echo SELinux已经是关闭状态"

	//DisableFirewalld 停止firewalld服务，并取消开机自启
	DisableFirewalld = "systemctl disable firewalld && systemctl stop firewalld"

	// DisableSwap 停止使用swap
	DisableSwap = "grep -v swap /etc/fstab  > /opt/tempData/fstab && cat /opt/tempData/fstab >/etc/fstab && mount -a && swapoff -a"

	//EnableChrony 启用chrony时间同步服务
	EnableChrony = "yum -y initialize chrony* \n" +
		"cat > /etc/chrony.conf <<EOF\n" +
		"pool ntp1.aliyun.com iburst\n" +
		"driftfile /var/lib/chrony/drift\n" +
		"makestep 1.0 3\n" +
		"rtcsync\n" +
		"keyfile /etc/chrony.keys\n" +
		"leapsectz right/UTC\n" +
		"logdir /var/log/chrony\n" +
		"EOF\n" +
		"systemctl enable --now chronyd && \n" +
		"systemctl restart chronyd"

	// KernelOptimize 内核优化
	KernelOptimize = "# 将桥接的IPv4流量传递到iptables的链\n" +
		"cat > /etc/sysctl.d/k8s.conf << EOF\n" +
		"net.ipv6.conf.all.disable_ipv6 = 1\n" +
		"net.ipv6.conf.default.disable_ipv6 = 1\n" +
		"net.ipv6.conf.lo.disable_ipv6 = 1\n" +
		"net.ipv4.ip_forward = 1\n" +
		"#iptables透明网桥的实现\n" +
		"# NOTE: kube-proxy 要求 NODE 节点操作系统中要具备 /sys/module/br_netfilter 文件，而且还要设置 bridge-nf-call-iptables=1，如果不满足要求，那么 kube-proxy 只是将检查信息记录到日志中，kube-proxy 仍然会正常运行，但是这样通过 Kube-proxy 设置的某些 iptables 规则就不会工作。\n" +
		"net.bridge.bridge-nf-call-ip6tables = 1\n" +
		"net.bridge.bridge-nf-call-iptables = 1\n" +
		"net.bridge.bridge-nf-call-arptables = 1\n" +
		"EOF\n" +
		"sysctl --system"

	// SoftwareInstall 安装基础软件包
	SoftwareInstall = "yum initialize -y " +
		"yum-utils device-mapper-persistent-data lvm2 rpcbind device-mapper " +
		"conntrack socat telnet lsof wget vim make gcc gcc-c++ pcre* " +
		"ipvsadm net-tools libnl libnl-devel openssl openssl-devel bash-completion"

	// EnableIptables 安装iptables，关闭即可，k8s自己初始化
	EnableIptables = "yum initialize iptables-services -y &&" +
		"systemctl stop iptables &&" +
		"systemctl disable iptables &&" +
		"iptables -F"

	// EnableIpvs 开启ipvs
	EnableIpvs = "cat > /etc/sysconfig/modules/ipvs.modules << EOF\n" +
		"modprobe -- ip_vs\n" +
		"modprobe -- ip_vs_lc\n" +
		"modprobe -- ip_vs_wlc\n" +
		"modprobe -- ip_vs_rr\n" +
		"modprobe -- ip_vs_wrr\n" +
		"modprobe -- ip_vs_lblc\n" +
		"modprobe -- ip_vs_lblcr\n" +
		"modprobe -- ip_vs_dh\n" +
		"modprobe -- ip_vs_sh\n" +
		"modprobe -- ip_vs_nq\n" +
		"modprobe -- ip_vs_sed\n" +
		"modprobe -- ip_vs_ftp\n" +
		"modprobe -- nf_conntrack\n" +
		"EOF\n" +
		"chmod 755 /etc/sysconfig/modules/ipvs.modules &&" +
		"bash /etc/sysconfig/modules/ipvs.modules"

	// DockerInstall 安装docker
	DockerInstall = "systemctl stop docker\n" +
		"tar xf /opt/tempData/docker-20.10.8.tgz -C /opt/tempData && \n" +
		"cp /opt/tempData/docker/* /usr/local/bin &&\n" +
		"dockerDisk=`df -Tk|grep -Ev \"devtmpfs|tmpfs|overlay\"|grep -E \"ext4|ext3|xfs\"|awk '/\\//{print \\$5,\\$NF}'|sort -nr|awk '{print \\$2}'|head -1|tr '\\n' ' '|awk '{print \\$1}'`\n" +
		"dockerDataDir=`echo ${dockerDisk}/dockerData|sed s'/\\/\\//\\//g'`\n" +
		"mkdir ${dockerDataDir} 2> /dev/null \n" +
		"cat > /etc/systemd/system/docker.service << EOF\n" +
		"[Unit]\n" +
		"Description=Docker Application Container Engine\n" +
		"Documentation=https://docs.docker.com\n" +
		"After=network-online.target firewalld.service\n" +
		"Wants=network-online.target\n\n" +
		"[Service]\n" +
		"Type=notify\n" +
		"ExecStart=/usr/local/bin/dockerd -H unix:///var/run/docker.sock -H tcp://0.0.0.0:2375 --data-root=${dockerDataDir} --config-file=/opt/docker/daemon.json\n" +
		"ExecReload=/bin/kill -s HUP \\$MAINPID\n" +
		"TimeoutSec=0\n" +
		"RestartSec=2\n" +
		"Restart=always\n\n" +
		"StartLimitBurst=3\n\n" +
		"tartLimitInterval=60s\n\n" +
		"LimitNOFILE=infinity\n" +
		"LimitNPROC=infinity\n" +
		"LimitCORE=infinity\n\n" +
		"TasksMax=infinity\n\n" +
		"Delegate=yes\n\n" +
		"KillMode=process\n\n" +
		"[Install]\n" +
		"WantedBy=multi-user.target\n" +
		"EOF\n" +
		"mkdir /opt/docker 2> /dev/null\n" +
		"cat > /opt/docker/daemon.json << EOF\n" +
		"{\n" +
		"    \"exec-opts\": [\"native.cgroupdriver=systemd\"],\n" +
		"    \"registry-mirrors\": [\n" +
		"        \"https://1nj0zren.mirror.aliyuncs.com\",\n" +
		"        \"https://kfwkfulq.mirror.aliyuncs.com\",\n" +
		"        \"https://2lqq34jg.mirror.aliyuncs.com\",\n" +
		"        \"https://pee6w651.mirror.aliyuncs.com\",\n" +
		"        \"http://hub-mirror.c.163.com\",\n" +
		"        \"https://docker.mirrors.ustc.edu.cn\",\n" +
		"        \"http://f1361db2.m.daocloud.io\",\n" +
		"        \"https://registry.docker-cn.com\"\n" +
		"    ]\n" +
		"}\n" +
		"EOF\n" +
		"systemctl daemon-reload && systemctl enable docker && systemctl restart docker\n"

	// BashCompletion CentOS Linux bash shell 自动补全
	BashCompletion = "chmod a+x /usr/share/bash-completion/bash_completion && /usr/share/bash-completion/bash_completion"

	// EtcdctlManagerCommand 生成etcdctl指令
	EtcdctlManagerCommand = "alias etcdctl3='ETCDCTL_API=3 etcdctl --cacert=/opt/caCenter/ca.pem --cert=/opt/etcd/ssl/etcd.pem --key=/opt/etcd/ssl/etcd-key.pem --endpoints=clientUrls'\n" +
		"alias etcdctl2='ETCDCTL_API=2 etcdctl --ca-file=/opt/caCenter/ca.pem --cert-file=/opt/etcd/ssl/etcd.pem --key-file=/opt/etcd/ssl/etcd-key.pem --endpoints=clientUrls'"
)

var InitScript = map[string]string{
	"SetHostname.sh":      SetHostname,
	"BashCompletion.sh":   BashCompletion,
	"DisableSELinux.sh":   DisableSELinux,
	"DisableFirewalld.sh": DisableFirewalld,
	"DisableSwap.sh":      DisableSwap,
	"EnableChrony.sh":     EnableChrony,
	"KernelOptimize.sh":   KernelOptimize,
	"SoftwareInstall.sh":  SoftwareInstall,
	"EnableIptables.sh":   EnableIptables,
	"EnableIpvs.sh":       EnableIpvs,
	"DockerInstall.sh":    DockerInstall,
}
