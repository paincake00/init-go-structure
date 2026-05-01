package config

import (
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

func fetchConfigPath() string {
	return os.Getenv(Path)
}
