package initialize

import (
	customConst "erbiaoOS/const"
	"erbiaoOS/utils/file"
)

// KubeSchedulerSystemdScript 生成kube-controller-manager systemd管理脚本
func KubeSchedulerSystemdScript() {
	file.Create(customConst.K8sMasterCfgDir+"kube-scheduler.service", customConst.KubeSchedulerSystemd)
}
