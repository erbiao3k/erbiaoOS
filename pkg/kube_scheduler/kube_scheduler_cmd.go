package kube_scheduler

import (
	myConst "erbiaoOS/const"
	"erbiaoOS/utils/file"
)

// systemdScript 生成kube-controller-manager systemd管理脚本
func systemdScript() {
	file.Create(myConst.K8sMasterCfgDir+"kube-scheduler.service", systemd)
}
