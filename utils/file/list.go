package file

import (
	"io/ioutil"
	"log"
	"strings"
)

const (
	FILE = "file"
	DIR  = "dir"
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

// List 列举目录下所有文件
func List(path string) (files, dirs []string) {

	infoList, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for info := range infoList {
		if !infoList[info].IsDir() {
			files = append(files, infoList[info].Name())
		} else {
			dirs = append(dirs, infoList[info].Name())
		}
	}
	return
}

// ListContains 匹配包含关键词列举目录下所有文件， 不包含子目录
func ListContains(path string, searchWords []string, obj string) (searchList []string) {
	if len(searchWords) == 0 {
		log.Fatal("匹配所需的关键字为空")
	}

	var info []string

	if obj == FILE {
		info, _ = List(path)
	} else if obj == DIR {
		_, info = List(path)
	}

	for _, f := range info {
		for _, s := range searchWords {
			if strings.Contains(f, s) {
				searchList = append(searchList, f)
			}
		}
	}
	return searchList
}

// ListHasPrefix 匹配以关键词开头目录下所有文件， 不包含子目录
func ListHasPrefix(path string, searchWords []string, obj string) (searchList []string) {
	if len(searchWords) == 0 {
		log.Fatal("匹配所需的关键字为空")
	}

	var info []string

	if obj == FILE {
		info, _ = List(path)
	} else if obj == DIR {
		_, info = List(path)
	}

	for _, f := range info {
		for _, s := range searchWords {
			if strings.HasPrefix(f, s) {
				searchList = append(searchList, f)
			}
		}
	}
	return searchList
}
