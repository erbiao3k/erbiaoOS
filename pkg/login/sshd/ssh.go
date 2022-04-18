package sshd

import (
	gossh "golang.org/x/crypto/ssh"
	"log"
	"net"
	"strings"
	"time"
)

type Cli struct {
	user       string
	password   string
	host       string
	client     *gossh.Client
	session    *gossh.Session
	LastResult string
}

// Connect 初始化ssh客户端
func (c *Cli) sshConnect() (*Cli, error) {
	config := &gossh.ClientConfig{}
	config.SetDefaults()
	config.User = c.user
	config.Timeout = time.Second * 5
	config.Auth = []gossh.AuthMethod{gossh.Password(c.password)}
	config.HostKeyCallback = func(hostname string, remote net.Addr, key gossh.PublicKey) error { return nil }
	client, err := gossh.Dial("tcp", c.host, config)
	if nil != err {
		return c, err
	}
	c.client = client
	return c, nil
}

// runShell 执行shell
func (c Cli) runShell(shell string) (string, error) {
	if c.client == nil {
		if _, err := c.sshConnect(); err != nil {
			log.Fatal("初始化ssh连接失败：", err)
		}
	}
	session, err := c.client.NewSession()
	if err != nil {
		log.Fatal("初始化ssh会话失败：", err)
	}
	// 关闭会话
	defer session.Close()
	buf, err := session.CombinedOutput(shell)

	c.LastResult = string(buf)
	return c.LastResult, err
}

// RemoteSshExec 远程执行指令
func RemoteSshExec(host, user, password, port, command string) string {
	cli := Cli{
		host:     host + ":" + port,
		user:     user,
		password: password,
	}
	// 建立连接对象
	c, err := cli.sshConnect()
	if err != nil {
		panic("创建ssh连接失败，请确认ssh地址、端口、账号、密码正确：" + err.Error())
	}
	// 退出时关闭连接
	defer c.client.Close()
	exec, _ := c.runShell(command + "|| echo ErrorFlag:$?")

	if strings.Contains(exec, "ErrorFlag") {
		log.Fatalf("在节点【%s】执行指令【%s】失败，执行结果：\n ------------\n%s------------", host, command, exec)
	} else {
		log.Printf("在节点【%s】执行指令【%s】成功，执行结果：\n ------------\n%s------------", host, command, exec)
	}
	return exec
}
