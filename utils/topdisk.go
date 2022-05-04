package utils

import (
	"erbiaoOS/utils/sshd"
	"erbiaoOS/vars"
	"fmt"
	"strings"
)

var (
	// topDisk 获取当前本地文件系统最大的分区
	topDisk = "df -Tk|grep -Ev \"devtmpfs|tmpfs|overlay\"|grep -E \"ext4|ext3|xfs\"|awk '/\\//{print $5,$NF}'|sort -nr|awk '{print $2}'|head -1|tr '\\n' ' '|awk '{print $1}'"
)

func LargeDisk(host *vars.HostInfo, dir string) (dataDir string) {
	disk := sshd.RemoteExec(host, topDisk)
	disk = strings.ReplaceAll(disk, "\n", "")
	dataDir = disk + "/" + dir
	sshd.RemoteExec(host, fmt.Sprintf("mkdir -p %s ", dataDir))
	return
}
