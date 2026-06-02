package config

import (
	"fmt"
	"net"
	"os"

	"github.com/joho/godotenv"
)

const (
	Path = "CONFIG_PATH"
)

type Config struct {
	Env string
}

func Load() (*Config, error) {
	path := fetchConfigPath()

	return LoadFromPath(path)
}

func LoadFromPath(path ...string) (*Config, error) {
	if err := godotenv.Load(path...); err != nil {
		return nil, err
	}

	cfg := &Config{
		Env: env.GetString("ENV", "local"),
	}

	return cfg, nil
}

func fetchHttpServerAddr() string {
	return net.JoinHostPort(env.GetString("HTTP_HOST", ""), env.GetString("HTTP_PORT", "8080"))
}

func fetchGRPCServerAddr() string {
	return net.JoinHostPort("", env.GetString("GRPC_SERVER_PORT", "44044"))
}

func fetchPostgresURI() string {
	user := env.GetString("POSTGRES_USER", "user")
	password := env.GetString("POSTGRES_PASSWORD", "secret")
	host := env.GetString("POSTGRES_HOST", "postgres")
	port := env.GetInt("POSTGRES_PORT", 5432)
	dbName := env.GetString("POSTGRES_DB", "db")

	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbName)
}

func fetchConfigPath() string {
	return os.Getenv(Path)
}
