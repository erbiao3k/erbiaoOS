package customConst

const (
	// CaConfigJson ca-config.json
	CaConfigJson = "{\n" +
		"    \"signing\": {\n" +
		"        \"default\": {\n" +
		"            \"expiry\": \"876000h\"\n" +
		"        },\n" +
		"        \"profiles\": {\n" +
		"            \"kubernetes\": {\n" +
		"                \"expiry\": \"876000h\",\n" +
		"                \"usages\": [\n" +
		"                    \"signing\",\n" +
		"                    \"key encipherment\",\n" +
		"                    \"server auth\",\n" +
		"                    \"client auth\"\n" +
		"                ]\n" +
		"            }\n" +
		"        }\n" +
		"    }\n" +
		"}"
	// CaCsrJson ca-csr.json
	CaCsrJson = "{\n" +
		"  \"CN\": \"kubernetes\",\n" +
		"  \"key\": {\n" +
		"      \"algo\": \"rsa\",\n" +
		"      \"size\": 2048\n" +
		"  },\n" +
		"  \"names\": [\n" +
		"    {\n" +
		"      \"C\": \"CN\",\n" +
		"      \"ST\": \"Beijing\",\n" +
		"      \"L\": \"Beijing\",\n" +
		"      \"O\": \"k8s\",\n" +
		"      \"OU\": \"system\"\n" +
		"    }\n" +
		"  ],\n" +
		"  \"ca\": {\n" +
		"          \"expiry\": \"876000h\"\n" +
		"  }\n" +
		"}"
	// EtcdCsrJson etcd-csr.json
	EtcdCsrJson = "{\n" +
		"  \"CN\": \"etcd\",\n" +
		"  \"hosts\": [\n" +
		"%s" +
		"    \"127.0.0.1\"\n" +
		"],\n" +
		"  \"key\": {\n" +
		"    \"algo\": \"rsa\",\n" +
		"    \"size\": 2048\n" +
		"  },\n" +
		"  \"names\": [{\n" +
		"    \"C\": \"CN\",\n" +
		"    \"ST\": \"Beijing\",\n" +
		"    \"L\": \"Beijing\",\n" +
		"    \"O\": \"k8s\",\n" +
		"    \"OU\": \"system\"\n" +
		"  }]\n" +
		"}"
	// KubeApiserverCsrJson kube-apiserver-csr.json
	KubeApiserverCsrJson = "{\n" +
		"    \"CN\": \"kubernetes\",\n" +
		"    \"hosts\": [\n" +
		"      \"10.255.0.1\",\n" +
		"      \"127.0.0.1\",\n" +
		"%s" +
		"      \"kubernetes\",\n" +
		"      \"kubernetes.default\",\n" +
		"      \"kubernetes.default.svc\",\n" +
		"      \"kubernetes.default.svc.cluster\",\n" +
		"      \"kubernetes.default.svc.cluster.local\"\n" +
		"    ],\n" +
		"    \"key\": {\n" +
		"        \"algo\": \"rsa\",\n" +
		"        \"size\": 2048\n" +
		"    },\n" +
		"    \"names\": [\n" +
		"        {\n" +
		"            \"C\": \"CN\",\n" +
		"            \"L\": \"BeiJing\",\n" +
		"            \"ST\": \"BeiJing\",\n" +
		"            \"O\": \"k8s\",\n" +
		"            \"OU\": \"System\"\n" +
		"        }\n" +
		"    ]\n" +
		"}"
	//KubeProxyCsrJson kube-proxy-csr.json
	KubeProxyCsrJson = "{\n" +
		"  \"CN\": \"system:kube-proxy\",\n" +
		"  \"hosts\": [],\n" +
		"  \"key\": {\n" +
		"    \"algo\": \"rsa\",\n" +
		"    \"size\": 2048\n" +
		"  },\n  \"names\": [\n" +
		"    {\n" +
		"      \"C\": \"CN\",\n" +
		"      \"L\": \"BeiJing\",\n" +
		"      \"ST\": \"BeiJing\",\n" +
		"      \"O\": \"k8s\",\n" +
		"      \"OU\": \"System\"\n" +
		"    }\n" +
		"  ]\n" +
		"}"
	// KubectlAdminCert kubectl指令使用的证书
	KubectlAdminCert = "{\n" +
		"  \"CN\": \"admin\",\n" +
		"  \"hosts\": [],\n" +
		"  \"key\": {\n" +
		"    \"algo\": \"rsa\",\n" +
		"    \"size\": 2048\n" +
		"  },\n" +
		"  \"names\": [\n" +
		"    {\n" +
		"      \"C\": \"CN\",\n" +
		"      \"L\": \"BeiJing\",\n" +
		"      \"ST\": \"BeiJing\",\n" +
		"      \"O\": \"system:masters\",\n" +
		"      \"OU\": \"System\"\n" +
		"    }\n" +
		"  ]\n" +
		"}"
	// KubeControllerManagerCsrJson kube-controller-manager-csr.json
	KubeControllerManagerCsrJson = "{\n" +
		"    \"CN\": \"system:kube-controller-manager\",\n" +
		"    \"key\": {\n" +
		"        \"algo\": \"rsa\",\n" +
		"        \"size\": 2048\n" +
		"    },\n" +
		"    \"hosts\": [\n" +
		"%s" +
		"      \"127.0.0.1\"\n" +
		"    ],\n" +
		"    \"names\": [\n" +
		"      {\n" +
		"        \"C\": \"CN\",\n" +
		"        \"ST\": \"Beijing\",\n" +
		"        \"L\": \"Beijing\",\n" +
		"        \"O\": \"system:kube-controller-manager\",\n" +
		"        \"OU\": \"system\"\n" +
		"      }\n" +
		"    ]\n" +
		"}"

	// KubeSchedulerCsrJson kube-scheduler-csr.json
	KubeSchedulerCsrJson = "{\n" +
		"    \"CN\": \"system:kube-scheduler\",\n" +
		"    \"key\": {\n" +
		"        \"algo\": \"rsa\",\n" +
		"        \"size\": 2048\n" +
		"    },\n" +
		"    \"hosts\": [\n" +
		"%s" +
		"      \"127.0.0.1\"\n" +
		"    ],\n" +
		"    \"names\": [\n" +
		"      {\n" +
		"        \"C\": \"CN\",\n" +
		"        \"ST\": \"Beijing\",\n" +
		"        \"L\": \"Beijing\",\n" +
		"        \"O\": \"system:kube-scheduler\",\n" +
		"        \"OU\": \"system\"\n" +
		"      }\n" +
		"    ]\n" +
		"}"
)
