package utils

import (
	"bytes"
	"fmt"
	"os/exec"
)

// ExecCmd 当前Linux执行指令
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
