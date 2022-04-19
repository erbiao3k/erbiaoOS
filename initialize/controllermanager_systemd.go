package initialize

import (
	customConst "erbiaoOS/const"
	"erbiaoOS/pkg/file"
)

// KubeControllerManagerSystemdScript 生成kube-controller-manager systemd管理脚本
func KubeControllerManagerSystemdScript() {
	file.Create(customConst.K8sMasterCfgDir+"kube-controller-manager.service", customConst.KubeControllerManagerSystemd)
}
