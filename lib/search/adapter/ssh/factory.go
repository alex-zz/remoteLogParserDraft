package ssh

import (
	"github.com/alex-zz/remoteLogParserDraft/lib/config"
	"github.com/alex-zz/remoteLogParserDraft/lib/search/adapter/pool"
)

type Factory struct {
	Settings *config.Settings
}

func (f *Factory) Create() (pool.Adapter, error){
	c, _ := CreateConnection(f.Settings)

	return c, nil
}