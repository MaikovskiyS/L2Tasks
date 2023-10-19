package config

type Config struct {
	port string
}

func New() *Config {
	return &Config{
		port: ":8080",
	}
}
func (c *Config) Port() string {
	return c.port
}
