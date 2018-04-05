package adapter

import (
	"github.com/alex-zz/remoteLogParserDraft/lib/search/adapter/pool"
	"github.com/alex-zz/remoteLogParserDraft/lib/search/adapter/ssh"
	"github.com/alex-zz/remoteLogParserDraft/lib/config"
	"errors"
)

const (
	AdapterSSH = "{{adapter.ssh}}"
)

func GetAvailableAdapters() []string {
	return []string{AdapterSSH}
}

func GetAdapterFactory(config *config.Connection) (pool.Creator, error) {
	var c pool.Creator
	var err error

	switch config.Adapter {
	case AdapterSSH:
		factory := &ssh.Factory{}
		factory.Config = &ssh.Config{
			Host: config.Settings.Host,
			Port: config.Settings.Port,
			User: config.Settings.User,
			Password: config.Settings.Password,
			KeyPath: config.Settings.KeyPath,
		}
		c = factory
	default:
		err = errors.New("unsupported adapter")
	}

	return c, err
}