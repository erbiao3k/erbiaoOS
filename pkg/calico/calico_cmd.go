package calico

import (
	myConst "erbiaoOS/const"
	"erbiaoOS/utils/file"
	"strings"
)

// cfg 初始化calico网络组件编排文件
func cfg() {
	cfg := strings.ReplaceAll(yaml, "192.168.0.0/16", "10.0.0.0/16")
	file.Create(myConst.K8sMasterCfgDir+"calico.yaml", cfg)
}
