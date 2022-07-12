package databases

import (
	"testing"
)

// Тестирование метода создания новой конфигурации mysql
func TestNewConfig(t *testing.T) {

	// тестовые кейсы
	tests := []struct {
		want       Config
		drivername string
		user       string
		password   string
		ip         string
		port       string
	}{
		{
			want: Config{
				drivername: "mysql",
				user:       "root",
				password:   "myPassWord",
				ip:         "127.0.0.1",
				port:       "3307",
			},
			drivername: "mysql",
			user:       "root",
			password:   "myPassWord",
			ip:         "127.0.0.1",
			port:       "3307",
		},
		{
			want: Config{
				drivername: "mysql",
				user:       "root",
				password:   "mypassWord",
				ip:         "127.0.0.1",
				port:       "3309",
			},
			drivername: "mysql",
			user:       "root",
			password:   "mypassWord",
			ip:         "127.0.0.1",
			port:       "3309",
		},
	}

	for _, tt := range tests {
		resp := NewConfig(tt.drivername, tt.user, tt.password, tt.ip, tt.port)
		if *resp != tt.want {
			t.Errorf("NewConfig(%v,%v,%v,%v,%v)=%v, wanted %v", tt.drivername, tt.user, tt.password, tt.ip, tt.port, resp, tt.want)
		}
	}
}

// Тестирование метода получения адреса базы данных
func TestGetDataSrc(t *testing.T) {

	tests := []struct {
		pick string
		want string
		test Config
	}{
		{
			test: Config{
				drivername: "mysql",
				user:       "root",
				password:   "myPassWord",
				ip:         "127.0.0.1",
				port:       "3307",
			},
			pick: "base",
			want: "root:myPassWord@tcp(127.0.0.1:3307)/base",
		},
		{
			test: Config{
				drivername: "mysql",
				user:       "root",
				password:   "mypassWord",
				ip:         "127.0.0.1",
				port:       "3309",
			},
			pick: "baze",
			want: "root:mypassWord@tcp(127.0.0.1:3309)/baze",
		},
	}

	for _, tt := range tests {
		resp := tt.test.GetDataSrc(tt.pick)
		if resp != tt.want {
			t.Errorf("NewConfig(%v)=%v, wanted %v", tt.pick, resp, tt.want)
		}
	}
}

// Тестирование метода получения движка базы данных
func TestGetDriver(t *testing.T) {

	tests := []struct {
		want string
		test Config
	}{
		{
			test: Config{
				drivername: "mysql",
				user:       "root",
				password:   "myPassWord",
				ip:         "127.0.0.1",
				port:       "3307",
			},
			want: "mysql",
		},
		{
			test: Config{
				drivername: "sqllite",
				user:       "root",
				password:   "mypassWord",
				ip:         "127.0.0.1",
				port:       "3309",
			},
			want: "sqllite",
		},
	}

	for _, tt := range tests {
		resp := tt.test.GetDriver()
		if resp != tt.want {
			t.Errorf("NewConfig()=%v, wanted %v", resp, tt.want)
		}
	}
}
