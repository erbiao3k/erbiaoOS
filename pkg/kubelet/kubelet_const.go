package kubelet

import (
	myConst "erbiaoOS/const"
	"erbiaoOS/pkg/cert"
	"erbiaoOS/setting"
	"fmt"
)

const (
	// cfgContent kubelet 配置文件
	cfgContent = "{\n" +
		"  \"kind\": \"KubeletConfiguration\",\n" +
		"  \"apiVersion\": \"kubelet.config.k8s.io/v1beta1\",\n" +
		"  \"authentication\": {\n" +
		"    \"x509\": {\n" +
		"      \"clientCAFile\": \"/opt/caCenter/ca.pem\"\n" +
		"    },\n" +
		"    \"webhook\": {\n" +
		"      \"enabled\": true,\n" +
		"      \"cacheTTL\": \"2m0s\"\n" +
		"    },\n" +
		"    \"anonymous\": {\n" +
		"      \"enabled\": false\n" +
		"    }\n" +
		"  },\n" +
		"  \"authorization\": {\n" +
		"    \"mode\": \"Webhook\",\n" +
		"    \"webhook\": {\n" +
		"      \"cacheAuthorizedTTL\": \"5m0s\",\n" +
		"      \"cacheUnauthorizedTTL\": \"30s\"\n" +
		"    }\n" +
		"  },\n" +
		"  \"address\": \"currentKubeletIP\",\n" +
		"  \"port\": 10250,\n" +
		"  \"readOnlyPort\": 10255,\n" +
		"  \"cgroupDriver\": \"systemd\",\n" +
		"  \"hairpinMode\": \"promiscuous-bridge\",\n" +
		"  \"serializeImagePulls\": false,\n" +
		"  \"clusterDomain\": \"cluster.local.\",\n" +
		"  \"clusterDNS\": [\"10.255.0.2\"]\n" +
		"}"

	//systemd kubelet systemd管理脚本
	systemd = "[Unit]\n" +
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

	user           = "system:kube-scheduler"
	context        = user
	kubeconfig     = myConst.TempDir + "kubelet-bootstrap.kubeconfig"
	publicKeyFile  = myConst.K8sSslDir + "kube-scheduler.pem"
	privateKeyFile = myConst.K8sSslDir + "kube-scheduler-key.pem"

	restartCmd = "systemctl daemon-reload && systemctl enable kube-scheduler && systemctl restart kube-scheduler && sleep 1"
)

var (
	setClusterCmd     = fmt.Sprintf(myConst.SetClusterCmd, cert.CaPubilcKeyFile, setting.RandMasterIP, kubeconfig)
	setCredentialsCmd = fmt.Sprintf(myConst.SetCredentialsCmd, user, publicKeyFile, privateKeyFile, kubeconfig)
	setContextCmd     = fmt.Sprintf(myConst.SetContextCmd, context, user, kubeconfig)
	useContextCmd     = fmt.Sprintf(myConst.UseContextCmd, context, kubeconfig)
)
