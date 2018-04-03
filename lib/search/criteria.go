package search

import (
	"github.com/alex-zz/remoteLogParserDraft/lib/request"
	"github.com/alex-zz/remoteLogParserDraft/lib/config"
)

type Criteria struct {
	PathToLogs string
	Fields []Сomparator
	LogFile struct {
		Record struct {
			Timezone   string
			Template   string
			DateFormat string
		}
		Name struct {
			Timezone   string
			Template   string
			DateFormat string
		}
	}
}

func Build(request request.Search, config config.Project) *Criteria {
	c := &Criteria{}
	return c
}