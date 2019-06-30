package main
import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
	"os"
	"path"
	"time"
)
func connect(user, password, host string, port int) (*sftp.Client, error) {
	var (
		auth     []ssh.AuthMethod
		addr string
		clientConfig *ssh.ClientConfig
		sshClient  *ssh.Client
		sftpClient  *sftp.Client
		err     error)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))
	clientConfig = &ssh.ClientConfig{
		User:  user,
		Auth:  auth,
		Timeout: 30 * time.Second,
		//需要验证服务端，不做验证返回nil就可以，点击HostKeyCallback看源码就知道了
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)
	if sshClient, err = ssh.Dial("tcp", addr, clientConfig);
	err != nil {   return nil, err  }

	// create sftp client
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		return nil, err
	}

	return sftpClient, nil
}

func main() {
	var (
		err error
		sftpClient *sftp.Client
	)
	// 这里换成实际的 SSH 连接的 用户名，密码，主机名或IP，SSH端口   
	sftpClient, err = connect("root", "312624", "192.168.1.7", 22)
	if err != nil {
		log.Fatal(err)
	}

	defer sftpClient.Close()
	// 用来测试的本地文件路径 和 远程机器上的文件夹   
	var localFilePath = "go.mod"
	var remoteDir = "/root/"

	srcFile, err := os.Open(localFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer srcFile.Close()

	var remoteFileName = path.Base(localFilePath)
	dstFile, err := sftpClient.Create(path.Join(remoteDir, remoteFileName))
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	buf := make([]byte, 1024)
	for {
		n, _ := srcFile.Read(buf)
		if n == 0 {
			break
		}
		dstFile.Write(buf)
	}

	fmt.Println("copy file to remote server finished!")
}
