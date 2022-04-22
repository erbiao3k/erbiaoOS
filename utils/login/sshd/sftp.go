package sshd

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"
)

// Connect 初始化sftp客户端
func conn(host, user, password, port string) (*sftp.Client, error) {
	var (
		auth         []ssh.AuthMethod
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
		sftpClient   *sftp.Client
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User:            user,
		Auth:            auth,
		Timeout:         time.Second * 5,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// 初始化ssh会话
	if sshClient, err = ssh.Dial("tcp", host+":"+port, clientConfig); err != nil {
		log.Fatal("初始化ssh会话失败：", err)
	}

	// 初始化sftp会话
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		log.Fatal("初始化sftp会话失败：", err)
	}

	return sftpClient, nil
}

// UploadFile sftp上传文件
func UploadFile(host, user, password, port, localFile, remoteDir string) {
	sftpClient, err := conn(host, user, password, port)
	if err != nil {
		panic("创建sftp连接失败，请确认ssh地址、端口、账号、密码正确：" + err.Error())
	}
	defer sftpClient.Close()

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
		log.Fatalf("为节点【%s】上传文件【%s】失败：【%s】", host, localFile, err)
	}
	fmt.Printf("节点【%s】上传文件【%s】完成", host, localFile)
}

// UploadDir sftp上传目录
func UploadDir(host, user, password, port, localPath, remotePath string) {
	sftpClient, err := conn(host, user, password, port)
	if err != nil {
		panic("创建sftp连接失败，请确认ssh地址、端口、账号、密码正确：" + err.Error())
	}
	defer sftpClient.Close()

	localfiles, err := ioutil.ReadDir(localPath)
	if err != nil {
		log.Fatal("读取目录下文件失败：", err)
	}

	for _, info := range localfiles {
		localfilePath := path.Join(localPath, info.Name())
		remotefilePath := path.Join(remotePath, info.Name())
		if info.IsDir() {
			sftpClient.Mkdir(remotefilePath)
			UploadDir(host, user, password, port, localfilePath, remotefilePath)
		} else {
			UploadFile(host, user, password, port, path.Join(localPath, info.Name()), remotePath)
		}
	}

	fmt.Printf("节点【%s】上传目录【%s】完成", host, localPath)
}
