package search

import (
	"github.com/alex-zz/remoteLogParserDraft/lib/search/types"
)

type Criteria struct {
	PathToLogs string
	Fields []struct{
		Name string
		Value types.Comparator
	}
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

//todo Import cycle with config
/*func Build(request request.Search, config config.Config) *Criteria {
	c := &Criteria{}

	projectConfig := config.GetProjectConfig(request.Project)
	envConfig := projectConfig.GetEnvironmentConfig(request.Environment)

	c.PathToLogs = envConfig.PathToLogs

	for _, field := range request.Fields {
		fieldConfig := projectConfig.GetFiledConfig(field.Name)
		fieldData := types.Data{
			Value: field.Value,
		}
		t, _ := types.BuildType(fieldConfig.Type, fieldData)

		field := struct {
			Name string
			Value types.Comparator
		}{
			Name : field.Name,
			Value: t,
		}
		c.Fields = append(c.Fields, field)
	}

	return c
}*/