package config

import (
	"github.com/asaskevich/govalidator"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"github.com/alex-zz/remoteLogParserDraft/lib/search/adapter"
	"github.com/alex-zz/remoteLogParserDraft/lib/search/types"
)

type Config struct {
	Connections []Connection `valid:"required"`
	Projects    []Project `valid:"required"`
}

func Load() (*Config, error) {
	c := &Config{}

	source, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(source, c)
	if err != nil {
		return nil, err
	}

	c.initValidators()
	err = c.validate()
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Config) GetConnectionConfig(name string) *Connection {
	var connConfig *Connection

	for _, connection := range c.Connections {
		if connection.Name == name {
			connConfig = &connection
			break
		}
	}

	return connConfig
}

func (c *Config) GetProjectConfig(name string) *Project {
	var projectConfig *Project

	for _, project := range c.Projects {
		if project.Name == name {
			projectConfig = &project
			break
		}
	}

	return projectConfig
}

func (c *Config) validate() error {
	_, err := govalidator.ValidateStruct(c)
	return err
}

func (c *Config) initValidators() {
	adapterValidator := govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {
		res := false
		adapters := adapter.GetAvailableAdapters()
		for _, v := range adapters {
			if i == v {
				res = true
				break
			}
		}

		return res
	})

	sshAuthValidator := govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {
		settings := context.(Settings)
		return len(settings.Password) > 0 || len(settings.KeyPath) > 0
	})

	connectionValidator := govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {
		connConfig := c.GetConnectionConfig(i.(string))
		return connConfig != nil
	})

	poolCapacityValidator := govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {
		settings := context.(EnvironmentSettings)
		capacity := settings.ConnectionPoolCapacity
		initCap := settings.ConnectionPoolInitCapacity

		if capacity <= 0 || initCap <= 0 {
			return false
		}

		if initCap > capacity {
			return false
		}

		return true
	})

	typeValidator := govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {
		res := false
		availableTypes := types.GetAvailableTypes()
		for _, v := range availableTypes {
			if i == v {
				res = true
				break
			}
		}

		return res
	})

	govalidator.CustomTypeTagMap.Set("adapterValidator", adapterValidator)
	govalidator.CustomTypeTagMap.Set("sshAuthValidator", sshAuthValidator)
	govalidator.CustomTypeTagMap.Set("connectionValidator", connectionValidator)
	govalidator.CustomTypeTagMap.Set("poolCapacityValidator", poolCapacityValidator)
	govalidator.CustomTypeTagMap.Set("typeValidator", typeValidator)
}