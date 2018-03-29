package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"github.com/asaskevich/govalidator"
)

type Config struct {
	Connections []Connection `valid:"required"`
	Projects    []Project
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

	err = c.Validate()
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Config) Validate() error {
	_, err := govalidator.ValidateStruct(c)
	return err
}