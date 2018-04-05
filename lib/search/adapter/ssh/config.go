package ssh

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	KeyPath  string
}