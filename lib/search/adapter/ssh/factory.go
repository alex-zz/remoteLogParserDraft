package ssh

import (
	"github.com/alex-zz/remoteLogParserDraft/lib/search/adapter/pool"
)

type Factory struct {
	Config *Config
}

func (f *Factory) Create() (pool.Adapter, error) {
	c, err := CreateConnection(f.Config)
	return c, err
}
