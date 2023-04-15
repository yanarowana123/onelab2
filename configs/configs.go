package configs

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	MySqlDSN              string
	PgSqlDSN              string
	WebServerPort         string
	JWTAccessTokenSecret  string
	JWTRefreshTokenSecret string
	JWTAccessTokenTTL     time.Duration
	JWTRefreshTokenTTL    time.Duration
	GracefulTimeout       time.Duration
}

func New() (*Config, error) {
	mySqlDSN := os.Getenv("MYSQL_DSN")
	if len(mySqlDSN) == 0 {
		mySqlDSN = "defaultVal"
	}

	pqSqlHost := os.Getenv("PGSQL_HOST")
	if len(pqSqlHost) == 0 {
		return nil, errors.New("please specify PGSQL_HOST variable in env")
	}

	pqSqlDB := os.Getenv("PGSQL_DB")
	if len(pqSqlDB) == 0 {
		return nil, errors.New("please specify PGSQL_DB variable in env")
	}

	pqSqlUser := os.Getenv("PGSQL_USER")
	if len(pqSqlUser) == 0 {
		return nil, errors.New("please specify PGSQL_USER variable in env")
	}

	pqSqlPassword := os.Getenv("PGSQL_PASSWORD")
	if len(pqSqlPassword) == 0 {
		return nil, errors.New("please specify PGSQL_PASSWORD variable in env")
	}

	pqSqlPort := os.Getenv("PGSQL_PORT")
	if len(pqSqlPort) == 0 {
		return nil, errors.New("please specify PGSQL_PORT variable in env")
	}

	pgSqlSSLMode := os.Getenv("PGSQL_SSL_MODE")
	if len(pgSqlSSLMode) == 0 {
		pgSqlSSLMode = "disable"
	}

	pgSqlDSN := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		pqSqlUser, pqSqlPassword, pqSqlHost, pqSqlPort, pqSqlDB, pgSqlSSLMode)
	webServerPort := os.Getenv("WEB_SERVER_PORT")
	if len(webServerPort) == 0 {
		return nil, errors.New("please specify WEB_SERVER_PORT variable in env")
	}

	jwtAccessTokenSecret := os.Getenv("JWT_ACCESS_TOKEN_SECRET")
	if len(jwtAccessTokenSecret) == 0 {
		return nil, errors.New("please specify JWT_ACCESS_TOKEN_SECRET variable in env")
	}

	jwtRefreshTokenSecret := os.Getenv("JWT_REFRESH_TOKEN_SECRET")
	if len(jwtRefreshTokenSecret) == 0 {
		return nil, errors.New("please specify JWT_REFRESH_TOKEN_SECRET variable in env")
	}

	jwtAccessTokenTTLString := os.Getenv("JWT_ACCESS_TOKEN_TTL")
	jwtAccessTokenTTL, _ := strconv.Atoi(jwtAccessTokenTTLString)
	if jwtAccessTokenTTL == 0 {
		jwtAccessTokenTTL = 100
	}

	jwtRefreshTokenTTLString := os.Getenv("JWT_REFRESH_TOKEN_TTL")
	jwtRefreshTokenTTL, _ := strconv.Atoi(jwtRefreshTokenTTLString)
	if jwtRefreshTokenTTL == 0 {
		jwtRefreshTokenTTL = 1000
	}

	GracefulTimeOutString := os.Getenv("GRACEFUL_TIMEOUT")
	GracefulTimeOut, _ := strconv.Atoi(GracefulTimeOutString)
	if GracefulTimeOut == 0 {
		GracefulTimeOut = 15
	}

	return &Config{
		MySqlDSN:              mySqlDSN,
		PgSqlDSN:              pgSqlDSN,
		WebServerPort:         webServerPort,
		JWTAccessTokenSecret:  jwtAccessTokenSecret,
		JWTRefreshTokenSecret: jwtRefreshTokenSecret,
		JWTAccessTokenTTL:     time.Duration(jwtAccessTokenTTL) * time.Second,
		JWTRefreshTokenTTL:    time.Duration(jwtRefreshTokenTTL) * time.Second,
		GracefulTimeout:       time.Duration(GracefulTimeOut) * time.Second,
	}, nil
}
