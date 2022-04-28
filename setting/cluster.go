package setting

import (
	"erbiaoOS/utils/login/sshd"
	"io"
	"log"
	"os"
	"strings"
)

const (
	// topDisk 获取当前本地文件系统最大的分区
	topDisk = "df -Tk|grep -Ev \"devtmpfs|tmpfs|overlay\"|grep -E \"ext4|ext3|xfs\"|awk '/\\//{print $5,$NF}'|sort -nr|awk '{print $2}'|head -1|tr '\\n' ' '|awk '{print $1}'"
)

// ClusterHost 集群节点初始化信息
type ClusterHost struct {
	K8sMaster   []HostInfo
	K8sNode     []HostInfo
	MysqlHost   []HostInfo
	FosHost     []HostInfo
	ConvertHost []HostInfo
}

// HostInfo 节点详细信息
type HostInfo struct {
	Role     string
	LanIp    string
	User     string
	Password string
	Port     string
	DataDir  string
	Mode     string
}

// fileContent 读取文件内容为,初始化集群host清单
func fileContent(path string) (string, error) {
	file, err := os.Open(path + "/hosts")
	if err != nil {
		log.Println("文件打开失败, err:", err)
		return "", err
	}
	defer file.Close()
	// 循环读取文件
	var content []byte
	var tmp = make([]byte, 128)
	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("文件读取失败, err:", err)
			return "", err
		}
		content = append(content, tmp[:n]...)
	}
	return string(content), nil
}

// contentAnalysis 文件内容分析
func contentAnalysis(content string) [][]string {
	var hostSlice [][]string
	// 按"\n"切分行
	temp := strings.Split(content, "\n")
	for _, line := range temp {
		// 去除首尾空格
		line = strings.TrimSpace(line)
		// 排除注释行
		if strings.HasPrefix(line, "#") {
			continue
		}
		for {
			// 强制将行内连续空格转换为单空格
			if strings.Contains(line, "  ") {
				line = strings.ReplaceAll(line, "  ", " ")
			} else {
				break
			}
		}
		// 行内按空格切分字段,获取每个节点的角色，服务器地址，服务器账号，服务器密码，服务器端口
		lineSlice := strings.Split(line, " ")

		// 空行剔除
		if len(lineSlice) != 1 {
			hostSlice = append(hostSlice, lineSlice)
		}
	}

	return hostSlice
}

// InitclusterHost 初始化集群节点信息
func InitclusterHost(path string) *ClusterHost {
	var ch ClusterHost
	var hi HostInfo
	content, _ := fileContent(path)
	analysis := contentAnalysis(content)

	for _, s := range analysis {
		hi.Role, hi.LanIp, hi.User, hi.Password, hi.Port, hi.Mode = s[0], s[1], s[2], s[3], s[4], s[5]

		hostInfo := &sshd.Info{
			LanIp:    hi.LanIp,
			User:     hi.User,
			Password: hi.Password,
			Port:     hi.Port,
		}

		hi.DataDir = sshd.RemoteSshExec(hostInfo, topDisk)
		hi.DataDir = strings.Split(hi.DataDir, "\n")[0]
		if hi.Role == "k8sMaster" {
			ch.K8sMaster = append(ch.K8sMaster, hi)
		}
		if hi.Role == "k8sNode" {
			ch.K8sNode = append(ch.K8sNode, hi)
		}
		if hi.Role == "mysqlHost" {
			ch.MysqlHost = append(ch.MysqlHost, hi)
		}
		if hi.Role == "fosHost" {
			ch.FosHost = append(ch.FosHost, hi)
		}
		if hi.Role == "convertHost" {
			ch.ConvertHost = append(ch.ConvertHost, hi)
		}
	}
	if len(ch.K8sMaster) < 1 || len(ch.K8sNode) < 1 {
		log.Fatal("请至少指定一个master节点和node节点")
	}
	return &ch
}
