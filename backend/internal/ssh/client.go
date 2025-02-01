package ssh

import (
	"bytes"
	"fmt"
	"time"

	"golang.org/x/crypto/ssh"
)

type Client struct {
	config *ssh.ClientConfig
}

func NewClient(user, password string) *Client {
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	}

	return &Client{config: config}
}

func (c *Client) ExecuteCommand(host string, port int, command string) (string, error) {
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", host, port), c.config)
	if err != nil {
		return "", fmt.Errorf("failed to dial: %v", err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		return "", fmt.Errorf("failed to create session: %v", err)
	}
	defer session.Close()

	var stdout, stderr bytes.Buffer
	session.Stdout = &stdout
	session.Stderr = &stderr

	if err := session.Run(command); err != nil {
		return "", fmt.Errorf("failed to run command: %v, stderr: %s", err, stderr.String())
	}

	return stdout.String(), nil
}
