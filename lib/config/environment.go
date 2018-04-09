package config

type Environment struct {
	Name       string              `yaml:"name" valid:"required"`
	PathToLogs string              `yaml:"pathToLogs" valid:"required"`
	Settings   EnvironmentSettings `yaml:"settings" valid:"required"`
}

type EnvironmentSettings struct {
	Connection                 string `yaml:"connection" valid:"connectionValidator"`
	ConnectionPoolCapacity     int    `yaml:"connectionPoolCapacity" valid:"required, poolCapacityValidator"`
	ConnectionPoolInitCapacity int    `yaml:"connectionPoolInitCapacity" valid:"required"`
}
