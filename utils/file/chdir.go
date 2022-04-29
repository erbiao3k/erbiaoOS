package file

import (
	"fmt"
	"log"
	"os"
)

// Chdir 切换目录
func Chdir(path string) {
	err := os.Chdir(path)
	if err != nil {
		log.Panicf("切换到目录%s失败：%s", path, err)
	}
	fmt.Sprintf("切换到目录：%s", path)
}
