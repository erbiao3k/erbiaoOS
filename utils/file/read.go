package file

import (
	"io/ioutil"
	"log"
)

// Read 读取文件内容
func Read(file string) string {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		log.Panicf("文件*%s读取错误：%s", f, err)
	}
	return string(f)
}
