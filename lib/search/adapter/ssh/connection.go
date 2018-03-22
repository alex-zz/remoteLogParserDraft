package ssh

import (
	"golang.org/x/crypto/ssh"
	"time"
	"strconv"
	"fmt"
	"strings"
	"github.com/alex-zz/remoteLogParserDraft/lib/config"
)

//todo keep alive

type Connection struct {
	conn *ssh.Client
}

func CreateConnection(settings *config.Settings) (*Connection, error) {
	c := &Connection{}

	//todo handle auth keys
	sshConfig := &ssh.ClientConfig{
		User: settings.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(settings.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout: time.Second * 5,
	}

	var err error

	c.conn, err = ssh.Dial("tcp", settings.Host+":"+strconv.Itoa(settings.Port), sshConfig)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Connection) RunCommand(command string) (string, int) {
	//todo handle session error
	session, err := c.conn.NewSession()
	if err != nil {
		c.Destroy()
		panic(fmt.Sprintf("Error executing command: %v", err))
	}
	defer session.Close()

	output, err := session.Output(command)
	exitCode := 0
	if err != nil && strings.HasPrefix(err.Error(), "Process exited with status ") {
		exitCodeStr := string(err.Error()[len("Process exited with status ")])
		exitCode, err = strconv.Atoi(exitCodeStr)
	} else if err != nil {
		c.Destroy()
		panic(fmt.Sprintf("Failed to run: %s", err))
	}

	return string(output), exitCode
}

func (c *Connection) Destroy() {
	if c.conn != nil {
		c.conn.Close()
	}
}

func (c * Connection) Find() {
	res, _ := c.RunCommand("echo Test From Server")
	fmt.Printf(res)
}