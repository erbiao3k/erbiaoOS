package kube_controllermanager

import (
	customConst "erbiaoOS/const"
	"erbiaoOS/utils/file"
)

// systemdScript 生成kube-controller-manager systemd管理脚本
func systemdScript() {
	file.Create(customConst.K8sMasterCfgDir+"kube-controller-manager.service", systemd)
}
