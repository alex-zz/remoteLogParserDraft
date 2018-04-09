package adapter

import (
	"github.com/alex-zz/remoteLogParserDraft/lib/search/adapter/pool"
	"github.com/alex-zz/remoteLogParserDraft/lib/search/adapter/ssh"
	"errors"
)

const (
	AdapterSSH = "{{adapter.ssh}}"
)

func GetAvailableAdapters() []string {
	return []string{AdapterSSH}
}

func GetAdapterFactory(adapterName string, config *Config) (pool.Creator, error) {
	var c pool.Creator
	var err error

	switch adapterName {
	case AdapterSSH:
		factory := &ssh.Factory{}
		factory.Config = &ssh.Config{
			Host: config.Host,
			Port: config.Port,
			User: config.User,
			Password: config.Password,
			KeyPath: config.KeyPath,
		}
		c = factory
	default:
		err = errors.New("unsupported adapter")
	}

	return c, err
}