package file

import (
	"io/ioutil"
	"log"
	"strings"
)

// isValueInSlice 判断字符串切片中是否包含某个value
func isValueInSlice(value string, slice []string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// List 列举目录下所有文件， 不包含子目录
func List(path string) []string {
	var fileList []string

	fileInfoList, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for f := range fileInfoList {
		if !fileInfoList[f].IsDir() {
			fileList = append(fileList, fileInfoList[f].Name())
		}
	}
	return fileList
}

// ListContains 匹配包含关键词列举目录下所有文件， 不包含子目录
func ListContains(path string, searchWords []string) []string {
	if len(searchWords) == 0 {
		log.Fatal("匹配所需的关键字为空")
	}

	fileList := List(path)
	var searchList []string

	for _, f := range fileList {
		for _, s := range searchWords {
			if strings.Contains(f, s) {
				searchList = append(searchList, f)
			}
		}
	}
	return searchList
}

// ListHasPrefix 匹配以关键词开头目录下所有文件， 不包含子目录
func ListHasPrefix(path string, searchWords []string) []string {
	if len(searchWords) == 0 {
		log.Fatal("匹配所需的关键字为空")
	}

	fileList := List(path)
	var searchList []string

	for _, f := range fileList {
		for _, s := range searchWords {
			if strings.HasPrefix(f, s) {
				searchList = append(searchList, f)
			}
		}
	}
	return searchList
}
