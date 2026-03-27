package config

import (
	"errors"
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env string `yaml:"env" env-default:"local"`
}

func MustLoad() (*Config, error) {
	path := fetchConfigPath()

	if path == "" {
		return nil, errors.New("config file path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, errors.New("config file not found:" + path)
	}

	var config Config

	if err := cleanenv.ReadConfig(path, &config); err != nil {
		return nil, errors.New("read config error:" + err.Error())
	}

	return &config, nil
}

// CONFIG_PATH=./path/to/config.yaml <file.bin>
// <file.bin> --config=./path...
func fetchConfigPath() string {
	var res string

	// --config="/path/to/config.yaml"
	flag.StringVar(&res, "config", "", "config path")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}