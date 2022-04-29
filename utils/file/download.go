package file

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// Download 单文件下载
func Download(url, downloadDir string) (fileName string) {

	fmt.Println("开始下载：", url)

	temp := strings.Split(url, "/")
	fileName = temp[len(temp)-1]

	client := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}

	res, err := client.Get(url)

	if err != nil {
		log.Fatal("下载错误：", err)
	}
	defer res.Body.Close()
	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(res.Body, 32*1024)

	file, err := os.Create(downloadDir + fileName)
	if err != nil {
		panic(err)
	}
	// 获得文件的writer对象
	writer := bufio.NewWriter(file)

	_, err = io.Copy(writer, reader)
	if err != nil {
		panic(err)
	}
	return
}
