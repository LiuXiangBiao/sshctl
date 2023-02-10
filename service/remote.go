package service

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"os"
	"time"
)

func CopyExec(RemoteFilePath, LocalFilePath string, ip, user, password string, cipherList []string) {
	var (
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		config       ssh.Config
		sftpclient   *sftp.Client
		err          error
	)

	if len(cipherList) == 0 {
		config = ssh.Config{
			Ciphers: []string{"aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-gcm@openssh.com", "arcfour256", "arcfour128", "aes128-cbc", "3des-cbc", "aes192-cbc", "aes256-cbc"},
		}
	} else {
		config = ssh.Config{
			Ciphers: cipherList,
		}
	}

	clientConfig = &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.Password(password)},
		Timeout:         30 * time.Second,
		Config:          config,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:22", ip)

	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		log.Fatal("client dail faield")
	}
	defer client.Close()

	sftpclient, err = sftp.NewClient(client)
	if err != nil {
		log.Fatal("sftp client faield")
	}
	defer sftpclient.Close()

	source, err := sftpclient.Open(RemoteFilePath)
	if err != nil {
		log.Fatal("Failed to open remote file: ", err)
	}
	defer source.Close()

	// 4. 创建本地文件
	target, err := os.OpenFile(LocalFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Failed to open local file: ", err)
	}
	defer target.Close()

	// 5. 数据复制
	n, err := io.Copy(target, source)
	if err != nil {
		log.Fatal("Failed to copy file: ", err)
	}
	log.Println("Succeed to copy file: ", n)

}
