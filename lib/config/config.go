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

func (c *Config) validate() error {
	_, err := govalidator.ValidateStruct(c)
	return err
}
