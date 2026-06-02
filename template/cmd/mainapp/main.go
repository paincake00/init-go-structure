package main

import (
	"log"
)

// Example run: CONFIG_PATH=".env" go run cmd/mainapp/main.go

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config", sl.Err(err))
	}

	app.Run(cfg)
}
