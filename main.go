package main

import "erbiaoOS/cmd"

// 节点最大磁盘识别拆分到sysinit
// 系统初始化逻辑拆分与优化
// 资源预留
// 节点labels
// kubectl label node k8smaster008021 k8smaster008041 k8smaster008048 node-role.kubernetes.io/master=true --overwrite=true
// kubectl taint nodes -l node-role.kubernetes.io/master node-role.kubernetes.io/master=true:NoSchedule  --overwrite=true
// kubectl label nodes k8snode008022 node-role.kubernetes.io/worker=true --overwrite=true

func main() {
	cmd.Execute()
}
