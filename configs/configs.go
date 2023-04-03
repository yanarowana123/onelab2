package configs

import (
	"os"
)

type Config struct {
	SqlDSN        string
	WebServerPort string
}

func New() *Config {
	sqlDSN := os.Getenv("sqlDSN")
	if len(sqlDSN) == 0 {
		sqlDSN = "defaultVal"
	}

	webServerPort := os.Getenv("webServerPort")
	if len(webServerPort) == 0 {
		webServerPort = "8080"
	}

	return &Config{
		SqlDSN:        sqlDSN,
		WebServerPort: webServerPort,
	}
}
