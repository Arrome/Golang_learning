package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"time"
)

func sshconnet(user, password, host string, port int) (*ssh.Session, error) {
	var (
		auth []ssh.AuthMethod
		addr string
		clientConfig *ssh.ClientConfig
		client *ssh.Client
		session *ssh.Session
		err error
	)
	//获取认证
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User: user,
		Auth: auth,
		Timeout: 30 * time.Second,
	}

	//连接ssh
	addr = fmt.Sprintf("%s:%d", host, port)
	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	if session, err = client.NewSession(); err != nil {
		return nil, err
	}

	return session, nil
}


func scpListCopy(cinfos []CouponInfo) error {
	var localFilePath, remoteDir string

	sshftpClient := NewSftpClient()

	for _, v := range cinfos {
		for i:=0; i<len(v.Coupons); i++ {
			localFilePath = v.Coupons[i].Addr
			remoteDir = ""
		}
	}
}
func main() {
	
}
