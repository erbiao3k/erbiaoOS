package setting

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Component struct {
	OfflineDeployment bool   `json:"OfflineDeployment"`
	Cfssl             string `json:"Cfssl"`
	Cfsslcertinfo     string `json:"Cfssl-certinfo"`
	Cfssljson         string `json:"Cfssljson"`
	Docker            string `json:"Docker"`
	Etcd              string `json:"Etcd"`
	Kubernetes        string `json:"Kubernetes"`
}

// ComponentContent 组件版本，下载地址
func ComponentContent(path string) *Component {
	file, err := os.Open(path + "/component.json")

	if err != nil {
		log.Fatal("文件打开失败, err:", err)
	}
	defer file.Close()

	jsonData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("组件配置文件内容读取失败")
	}
	var post Component
	json.Unmarshal(jsonData, &post)

	return &post
}
