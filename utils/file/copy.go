package file

import (
	"io"
	"os"
)

// Copy 单文件复制
func Copy(dstFileName string, srcFileName string) {

	srcFile, err := os.Open(srcFileName)

	if err != nil {
		panic("源文件打开失败：" + err.Error())
	}

	defer srcFile.Close()

	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		panic("目标文件打开失败：" + err.Error())
	}

	defer dstFile.Close()

	io.Copy(dstFile, srcFile)

}
