package ssh

import (
	"github.com/alex-zz/remoteLogParserDraft/lib/config"
	"github.com/alex-zz/remoteLogParserDraft/lib/search/adapter/pool"
)

type Factory struct {
	ConnectionConfig *config.Connection
}

func (f *Factory) Create() (pool.Adapter, error) {
	c, err := CreateConnection(f.ConnectionConfig)
	return c, err
}
