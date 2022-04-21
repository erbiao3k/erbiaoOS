package initialize

import (
	customConst "erbiaoOS/const"
	"erbiaoOS/utils/file"
	"strings"
)

// CorednsCfg 初始化coredns组件编排文件
func CorednsCfg() {
	cfg := strings.ReplaceAll(customConst.CoreDnsYaml, "$DNS_SERVER_IP", "10.255.0.2")
	cfg = strings.ReplaceAll(customConst.CoreDnsYaml, "$DNS_DOMAIN", "cluster.local")
	cfg = strings.ReplaceAll(customConst.CoreDnsYaml, "$DNS_MEMORY_LIMIT", "170Mi")
	cfg = strings.ReplaceAll(customConst.CoreDnsYaml, "k8s.gcr.io/coredns/coredns:v1.8.6", "registry.aliyuncs.com/google_containers/coredns:v1.8.6")
	file.Create(customConst.K8sMasterCfgDir+"coredns.yaml", cfg)
}
