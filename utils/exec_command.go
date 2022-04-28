package utils

import (
	"bytes"
	"fmt"
	"os/exec"
)

// ExecCmd 单Linux指令执行
func ExecCmd(command string) {
	cmd := exec.Command("bash", "-c", command)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(out.String())
		panic(err)
	}
}

// MultiExecCmd 多Linux指令循环执行
func MultiExecCmd(cmds []string) {
	for _, cmd := range cmds {
		ExecCmd(cmd)
	}
}
