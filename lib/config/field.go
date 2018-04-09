package config

type Field struct {
	Name       string `yaml:"name" valid:"alphanum"`
	Type       string `yaml:"type" valid:"typeValidator"`
	IsRequired bool   `yaml:"isRequired"`
}
