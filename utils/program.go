package utils

import (
	"os"
	"strings"
)

var ProgramName = Program()

// Program 获取程序名
func Program() string {
	temp := strings.Split(os.Args[0], "/")
	return temp[len(temp)-1]
}
