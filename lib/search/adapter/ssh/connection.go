package ssh

import (
	"fmt"
	"github.com/alex-zz/remoteLogParserDraft/lib/config"
	"golang.org/x/crypto/ssh"
	"strconv"
	"strings"
	"time"
	"github.com/alex-zz/remoteLogParserDraft/lib/search"
)

type Connection struct {
	conn *ssh.Client
}

func CreateConnection(connectionConfig *config.Connection) (*Connection, error) {
	c := &Connection{}

	//todo handle auth keys
	sshConfig := &ssh.ClientConfig{
		User: connectionConfig.Settings.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(connectionConfig.Settings.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         time.Second * 5,
	}

	var err error

	addr := connectionConfig.Settings.Host + ":" + strconv.Itoa(connectionConfig.Settings.Port)
	c.conn, err = ssh.Dial("tcp", addr, sshConfig)
	if err != nil {
		return nil, err
	}

	//todo keep alive
	/*
		    go func() {
	        t := time.NewTicker(2 * time.Second)
	        defer t.Stop()
	        for {
	            <-t.C
	            _, _, err := client.Conn.SendRequest("keepalive@golang.org", true, nil)
	            if err != nil {
	                return
	            }
	        }
	    }()

	*/

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

func (c *Connection) IsActive() bool {
	return true
}

func (c *Connection) Find(criteria *search.Criteria) (*search.Result, error) {
	res, _ := c.RunCommand("echo Test From Server")
	fmt.Printf(res)
	r := &search.Result{}
	return r, nil
}
