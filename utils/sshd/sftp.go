package sshd

import (
	"erbiaoOS/utils/file"
	"erbiaoOS/vars"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"
)

// Connect 初始化sftp客户端
func conn(host *vars.HostInfo) *sftp.Client {
	var (
		auth         []ssh.AuthMethod
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
		sftpClient   *sftp.Client
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(host.Password))

	clientConfig = &ssh.ClientConfig{
		User:            host.User,
		Auth:            auth,
		Timeout:         time.Second * 5,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// 初始化ssh会话
	if sshClient, err = ssh.Dial("tcp", host.LanIp+":"+host.Port, clientConfig); err != nil {
		log.Fatal("初始化ssh会话失败：", err)
	}

	// 初始化sftp会话
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		log.Fatal("创建sftp连接失败，请确认ssh地址、端口、账号、密码正确：" + err.Error())
	}

	return sftpClient
}

// File sftp上传文件
func File(sftpClient *sftp.Client, localFile, remoteDir string) {

	srcFile, err := os.Open(localFile)
	if err != nil {
		log.Fatal(err)
	}

	fi, err := os.Stat(localFile)
	if err != nil {
		log.Fatal(err)
	}

	defer srcFile.Close()

	var remoteFileName = path.Base(localFile)
	dstFile, err := sftpClient.Create(path.Join(remoteDir, remoteFileName))

	defer dstFile.Close()

	buf := make([]byte, fi.Size())
	for {
		n, _ := srcFile.Read(buf)

		if n == 0 {
			break
		}
		dstFile.Write(buf)

	}

	if err != nil {
		log.Fatalf("为节点上传文件【%s】失败：【%s】\n", localFile, err)
	}
}

// Dir sftp上传目录
func Dir(sftpClient *sftp.Client, localPath, remotePath string) {

	localfiles, err := ioutil.ReadDir(localPath)
	if err != nil {
		log.Fatal("读取目录下文件失败：", err)
	}

	for _, info := range localfiles {
		localfilePath := path.Join(localPath, info.Name())
		remotefilePath := path.Join(remotePath, info.Name())
		if info.IsDir() {
			Dir(sftpClient, localfilePath, remotefilePath)
		} else {
			File(sftpClient, path.Join(localPath, info.Name()), remotePath)
		}
	}
}

// Upload 上传文件或目录总入口
func Upload(host *vars.HostInfo, localThing, remoteDir string) {
	sftpClient := conn(host)

	defer sftpClient.Close()

	if !file.Exist(localThing) {
		log.Panicf("上传的文件或目录%s不存在", localThing)
	}
	if file.IsDir(localThing) {
		RemoteExec(host, "mkdir -p "+remoteDir)
		Dir(sftpClient, localThing, remoteDir)
	} else if file.IsFile(localThing) {
		File(sftpClient, localThing, remoteDir)
	}
}
