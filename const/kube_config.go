package customConst

const (
	// KubeletCfg kubelet 配置文件
	KubeletCfg = "{\n" +
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

	// KubeProxyCfg kube-proxy配置文件
	KubeProxyCfg = "apiVersion: kubeproxy.config.k8s.io/v1alpha1\n" +
		"bindAddress: currentKubeproxyIP\n" +
		"clientConnection:\n" +
		"  kubeconfig: /opt/kubernetes/cfg/kube-proxy.kubeconfig\n" +
		"clusterCIDR: 10.0.0.0/16\n" +
		"healthzBindAddress: currentKubeproxyIP:10256\n" +
		"kind: KubeProxyConfiguration\n" +
		"metricsBindAddress: currentKubeproxyIP:10249\n" +
		"mode: \"ipvs\""
)
