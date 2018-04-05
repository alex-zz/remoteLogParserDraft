package config

import (
	"github.com/asaskevich/govalidator"
	"gopkg.in/yaml.v2"
	"io/ioutil"
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
