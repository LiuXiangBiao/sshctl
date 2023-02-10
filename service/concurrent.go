package service

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func SshExec(username, password string, hostlist []string, cmdlist []string, cipherList []string) {
	wg.Add(len(hostlist))

	for _, host := range hostlist {
		go Execmd(username, password, host, cmdlist, cipherList)
	}

	wg.Wait()
}

func Execmd(username, password, host string, cmdlist []string, cipherList []string) {
	var (
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		config       ssh.Config
		session      *ssh.Session
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
		User:            username,
		Auth:            []ssh.AuthMethod{ssh.Password(password)},
		Timeout:         time.Second,
		Config:          config,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:22", host)

	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	// create session
	if session, err = client.NewSession(); err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	if err := session.RequestPty("", 0, 0, modes); err != nil {
		log.Fatal(err)
	}

	cmdlist = append(cmdlist, "exit")
	newcmd := strings.Join(cmdlist, "&&")

	out, err := session.CombinedOutput(newcmd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(">>>>>>=======%s Result======<<<<<<\n", host)
	fmt.Println(string(out))

	wg.Done()
}
