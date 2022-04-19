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
)
