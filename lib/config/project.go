package config

type Project struct {
	Name string `yaml:"name"`
	Settings struct {
		ConnectionPoolCapacity     int    `yaml:"connectionPoolCapacity"`
		Connection                 string `yaml:"connection"`
		ConnectionPoolInitCapacity int    `yaml:"connectionPoolInitCapacity"`
	} `yaml:"settings"`
	Search struct {
		Fields []struct {
			IsRequired bool   `yaml:"isRequired"`
			Type       string `yaml:"type"`
			Name       string `yaml:"name"`
			Format     string `yaml:"format"`
		} `yaml:"fields"`
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
	Environments []struct {
		Name       string `yaml:"name"`
		PathToLogs string `yaml:"pathToLogs"`
	} `yaml:"environments"`
}
