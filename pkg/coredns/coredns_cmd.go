package coredns

import (
	myConst "erbiaoOS/const"
	"erbiaoOS/utils/file"
	"strings"
)

// Cfg 初始化coredns组件编排文件
func cfg() {
	cfg := strings.ReplaceAll(yaml, "$DNS_SERVER_IP", "10.255.0.2")
	cfg = strings.ReplaceAll(yaml, "$DNS_DOMAIN", "cluster.local")
	cfg = strings.ReplaceAll(yaml, "$DNS_MEMORY_LIMIT", "170Mi")
	cfg = strings.ReplaceAll(yaml, "k8s.gcr.io/coredns/coredns:v1.8.6", "registry.aliyuncs.com/google_containers/coredns:v1.8.6")
	file.Create(myConst.K8sMasterCfgDir+"coredns.yaml", cfg)
}
