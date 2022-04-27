package calico

import (
	myConst "erbiaoOS/const"
	"erbiaoOS/utils"
	"erbiaoOS/utils/file"
	"strings"
)

const (
	calicoYaml = myConst.TempDir + "calico.yaml"
)

// config 初始化calico网络组件编排文件
func config() {
	cfg := strings.ReplaceAll(yaml, "#   value: \"192.168.0.0/16\"", "  value: \"10.0.0.0/16\"")
	cfg = strings.ReplaceAll(cfg, "# - name: CALICO_IPV4POOL_CIDR", "- name: CALICO_IPV4POOL_CIDR")
	cfg = strings.ReplaceAll(cfg, "calico-yaml", "calico-config")
	file.Create(calicoYaml, cfg)
}

// Deploy 部署网络插件
func Deploy() {
	config()
	utils.ExecCmd("kubectl apply -f " + calicoYaml)
}
