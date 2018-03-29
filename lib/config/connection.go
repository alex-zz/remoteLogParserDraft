package config

type Connection struct {
	Name    string `yaml:"name" valid:"required"`
	Adapter string `yaml:"adapter"`
	Settings struct {
		Host     string `yaml:"host" valid:"required"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		KeyPath  string `yaml:"keyPath"`
	} `yaml:"settings"`
}
