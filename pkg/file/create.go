package file

import (
	"fmt"
	"os"
)

// Create 创建文件
func Create(filename string, content string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Printf("文件打开失败%s, err：%s", filename, err)
		return
	}
	defer file.Close()
	fmt.Println("生成文件：", filename)
	file.Write([]byte(content + "\n")) //写入字节切片数据
	//file.WriteString(content + "\n") //直接写入字符串数据
}
