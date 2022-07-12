package databases

import (
	f "fmt"
)

// Congif ...
type Config struct {
	drivername string
	user       string
	password   string
	ip         string
	port       string
}

// Создание новой конфигурации
func NewConfig(drivername, user, password, ip, port string) *Config {
	return &Config{
		drivername: drivername,
		user:       user,
		password:   password,
		ip:         ip,
		port:       port,
	}
}

// для получения sql драйвера
func (c *Config) GetDriver() string {
	return c.drivername
}

// дял получения DataSource
func (c *Config) GetDataSrc(s string) string {
	return f.Sprintf("%s:%s@tcp(%s:%s)/%s", c.user, c.password, c.ip, c.port, s)
}
