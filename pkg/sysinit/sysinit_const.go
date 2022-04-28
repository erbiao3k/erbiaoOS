package sysinit

import "erbiaoOS/utils/file"

const (
	hostsFile      = "/etc/hosts"
	HostsFileBak   = "/etc/hosts.bak-fadada"
	BashProfile    = "/root/.bash_profile"
	BashProfileBak = "/root/.bash_profile.bak-fadada"
	SysConfigDir   = "/etc/"

	// RemoveSoft 清理可能阻塞部署的软件
	RemoveSoft = "yum -y remove docker* || echo 未安装软件包"

	// StopService 清理可能阻塞部署的进程
	StopService = "systemctl stop etcd docker kube-apiserver kube-controller-manager kube-proxy kube-scheduler kubelet nginx || echo 服务已停止"

	// setHostname 设置主机名的字符串
	setHostname = "hostnamectl set-hostname "

	// disableSELinux 关闭SELinux的字符串
	disableSELinux = "sed -i 's/\\=enforcing/\\=disabled/' /etc/selinux/config &&  setenforce 0 || echo SELinux已经是关闭状态"

	//disableFirewalld 停止firewalld服务，并取消开机自启
	disableFirewalld = "systemctl disable firewalld && systemctl stop firewalld"

	// disableSwap 停止使用swap
	disableSwap = "grep -v swap /etc/fstab  > /tmp/fstab && cat /tmp/fstab >/etc/fstab && mount -a && swapoff -a"

	//enableChrony 启用chrony时间同步服务
	enableChrony = "yum -y install chrony* \n" +
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

	// kernelOptimize 内核优化
	kernelOptimize = "# 将桥接的IPv4流量传递到iptables的链\n" +
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

	// softwareInstall 安装基础软件包
	softwareInstall = "yum install -y " +
		"yum-utils device-mapper-persistent-data lvm2 rpcbind device-mapper " +
		"conntrack socat telnet lsof wget vim make gcc gcc-c++ pcre* " +
		"ipvsadm net-tools libnl libnl-devel openssl openssl-devel bash-completion"

	// enableIptables 安装iptables，关闭即可，k8s自己初始化
	enableIptables = "yum install iptables-services -y &&" +
		"systemctl stop iptables &&" +
		"systemctl disable iptables &&" +
		"iptables -F"

	// enableIpvs 开启ipvs
	enableIpvs = "cat > /etc/sysconfig/modules/ipvs.modules << EOF\n" +
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

	// dockerInstall 安装docker
	dockerInstall = "systemctl stop docker\n" +
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
		"Restart=always\n" +
		"StartLimitBurst=3\n" +
		"LimitNOFILE=infinity\n" +
		"LimitNPROC=infinity\n" +
		"LimitCORE=infinity\n" +
		"TasksMax=infinity\n" +
		"Delegate=yes\n" +
		"KillMode=process\n" +
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

	// bashCompletion CentOS Linux bash shell 自动补全
	bashCompletion = "chmod a+x /usr/share/bash-completion/bash_completion && /usr/share/bash-completion/bash_completion"
)

var (
	script = map[string]string{
		"SetHostname.sh":      setHostname,
		"BashCompletion.sh":   bashCompletion,
		"DisableSELinux.sh":   disableSELinux,
		"DisableFirewalld.sh": disableFirewalld,
		"DisableSwap.sh":      disableSwap,
		"EnableChrony.sh":     enableChrony,
		"KernelOptimize.sh":   kernelOptimize,
		"SoftwareInstall.sh":  softwareInstall,
		"EnableIptables.sh":   enableIptables,
		"EnableIpvs.sh":       enableIpvs,
		"DockerInstall.sh":    dockerInstall,
	}

	SysHost                = hostFileInitContent()[0]
	CurrentUserBashProfile = hostFileInitContent()[1]
)

// 初始化的/etc/hosts信息和/root/.bash_profile信息
func hostFileInitContent() (initInfo []string) {

	if !file.Exist(HostsFileBak) {
		file.Copy(HostsFileBak, hostsFile)
		initInfo = append(initInfo, file.Read(hostsFile))
	} else {
		initInfo = append(initInfo, file.Read(HostsFileBak))
	}

	if !file.Exist(BashProfileBak) {
		file.Copy(BashProfileBak, BashProfile)
		initInfo = append(initInfo, file.Read(BashProfile))
	} else {
		initInfo = append(initInfo, file.Read(BashProfileBak))
	}
	return
}
