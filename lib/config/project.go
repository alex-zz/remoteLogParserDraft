package config

type Project struct {
	Name string `yaml:"name"`
	Search struct {
		Fields []Field `yaml:"fields"`
		LogFile struct {
			Record struct {
				Timezone   string `yaml:"timezone"`
				Template   string `yaml:"template"`
				DateFormat string `yaml:"dateFormat"`
			} `yaml:"record"`
			Name struct {
				Timezone   string `yaml:"timezone"`
				Template   string `yaml:"template"`
				DateFormat string `yaml:"dateFormat"`
			} `yaml:"name"`
		} `yaml:"logFile"`
		Settings struct {
			MaxSearchIntervalSeconds int `yaml:"maxSearchIntervalSeconds"`
		} `yaml:"settings"`
	} `yaml:"search"`
	Environments []Environment `yaml:"environments"`
}

type Environment struct {
	Name       string `yaml:"name"`
	PathToLogs string `yaml:"pathToLogs"`
	Settings struct {
		ConnectionPoolCapacity     int    `yaml:"connectionPoolCapacity"`
		Connection                 string `yaml:"connection"`
		ConnectionPoolInitCapacity int    `yaml:"connectionPoolInitCapacity"`
	} `yaml:"settings"`
}

type Field struct {
	IsRequired bool   `yaml:"isRequired"`
	Type       string `yaml:"type"`
	Name       string `yaml:"name"`
	Format     string `yaml:"format"`
}

func (p *Project) GetEnvironmentConfig(name string) *Environment {
	var envConfig *Environment

	for _, env := range p.Environments {
		if env.Name == name {
			envConfig = &env
			break
		}
	}

	return envConfig
}

func (p *Project) GetFiledConfig(name string) *Field {
	var fieldConfig *Field

	for _, field := range p.Search.Fields {
		if field.Name == name {
			fieldConfig = &field
			break
		}
	}

	return fieldConfig
}
