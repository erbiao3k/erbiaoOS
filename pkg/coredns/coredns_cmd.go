package coredns

import (
	"erbiaoOS/utils"
	"erbiaoOS/utils/file"
	myConst "erbiaoOS/vars"
	"strings"
)

const (
	corednsYaml = myConst.TempDir + "coredns.yaml"
)

// conifg 初始化coredns组件编排文件
func conifg() {
	cfg := strings.ReplaceAll(yaml, "$DNS_SERVER_IP", "10.255.0.2")
	cfg = strings.ReplaceAll(cfg, "$DNS_DOMAIN", "cluster.local")
	cfg = strings.ReplaceAll(cfg, "$DNS_MEMORY_LIMIT", "170Mi")
	cfg = strings.ReplaceAll(cfg, "k8s.gcr.io/coredns/coredns:v1.8.6", "registry.aliyuncs.com/google_containers/coredns:v1.8.6")
	file.Create(corednsYaml, cfg)
}

func Deploy() {
	conifg()
	utils.ExecCmd("kubectl apply -f " + corednsYaml)
}
