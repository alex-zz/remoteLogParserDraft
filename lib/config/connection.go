package config

type Connection struct {
	Name     string   `yaml:"name" valid:"required"`
	Adapter  string   `yaml:"adapter" valid:"adapterValidator"`
	Settings Settings `yaml:"settings" valid:"required"`
}

type Settings struct {
	Host     string `yaml:"host" valid:"required"`
	Port     string `yaml:"port" valid:"required, port"`
	User     string `yaml:"user" valid:"required"`
	Password string `yaml:"password" valid:"sshAuthValidator"`
	KeyPath  string `yaml:"keyPath"`
}
