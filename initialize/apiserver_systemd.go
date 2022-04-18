package initialize

import (
	customConst "erbiaoOS/const"
	"erbiaoOS/pkg/file"
)

// KubeApiserverSystemdScript 生成kube-apiserver systemd管理脚本
func KubeApiserverSystemdScript() {
	file.Create(customConst.K8sMasterCfgDir+"kube-apiserver.service", customConst.KubeApiserverSystemd)
}
