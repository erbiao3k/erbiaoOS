package myConst

import (
	"os"
	"strings"
)

var MasterIPs []string
var NodeIPs []string
var SshUser string
var SshPassword string
var SshPort string
var K8sPkg string
var EtcdPkg string
var NginxPkg string
var DockerPkg string

var ProgramName = program()

// 获取程序名
func program() string {
	temp := strings.Split(os.Args[0], "/")
	return temp[len(temp)-1]
}
