package main

import (
	"log"
	"pm/config"
	"pm/internal/app"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("config error:", err)
	}
	cfg.Log()
	if err := app.Run(cfg); err != nil {
		log.Fatal("server error:", err)
	}
}
