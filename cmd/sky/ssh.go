package main

import (
	"code.google.com/p/go.crypto/ssh"
	"code.google.com/p/gopass"
	"errors"
)

type SSHConn struct {
	host   string
	user   string
	client *ssh.ClientConn
}

// SSH logic
func (c *SSHConn) Connect(host, user string) error {
	c.host = host
	c.user = user

	config := &ssh.ClientConfig{
		User: c.user,
		Auth: []ssh.ClientAuth{
			ssh.ClientAuthPassword(c),
		},
	}

	var err error
	c.client, err = ssh.Dial("tcp", c.host, config)
	if err != nil {
		return err
	}

	return nil
}

func (c *SSHConn) Exec(cmd string) (out []byte, err error) {
	var session *ssh.Session

	session, err = c.client.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	defer session.Close()

	return session.CombinedOutput(cmd)
}

func (c *SSHConn) Password(user string) (string, error) {
	pass, err := gopass.GetPass("Password for " + user + ": ")

	if err != nil {
		return "", errors.New("Failed to collect password: " + err.Error())
	}

	return pass, err
}

func (c *SSHConn) Close() {
	c.client.Close()
}

func (t *SSHConn) SetEnv(name, value string) {
	// TODO:
	return
}
