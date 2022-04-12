package file

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

// UnTargz tar.gz解压
func UnTargz(file string, dest string) {
	fmt.Println("正在解压文件：", file)
	fr, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer fr.Close()

	gr, err := gzip.NewReader(fr)
	if err != nil {
		panic(err)
	}
	defer gr.Close()

	tr := tar.NewReader(gr)
	for {
		h, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if h.FileInfo().IsDir() {
			os.MkdirAll(dest+h.Name, 0777)
		} else {
			fw, err := os.OpenFile(dest+h.Name, os.O_CREATE|os.O_WRONLY, 0777)
			if err != nil {
				panic(err)
			}
			// 写文件
			_, err = io.Copy(fw, tr)
			if err != nil {
				panic(err)
			}
		}

	}
}
