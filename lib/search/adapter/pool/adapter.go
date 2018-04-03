package pool

import "github.com/alex-zz/remoteLogParserDraft/lib/search"

type Adapter interface {
	Destroy()
	Find(criteria *search.Criteria) (*search.Result, error)
	IsActive() bool
}
