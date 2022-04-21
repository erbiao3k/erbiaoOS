package initialize

import (
	customConst "erbiaoOS/const"
	"erbiaoOS/utils/file"
	"strings"
)

// CalicoCfg 初始化calico网络组件编排文件
func CalicoCfg() {
	cfg := strings.ReplaceAll(customConst.CalicoYaml, "192.168.0.0/16", "10.0.0.0/16")
	file.Create(customConst.K8sMasterCfgDir+"calico.yaml", cfg)
}
