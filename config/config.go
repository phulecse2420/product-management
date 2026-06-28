package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost, DBPort, DBUser, DBPassword, DBName, Port string
}

func Load() (*Config, error) {
	_ = godotenv.Load()
	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres_user"),
		DBPassword: getEnv("DB_PASSWORD", "postgres_password"),
		DBName:     getEnv("DB_NAME", "product_management"),
		Port:       getEnv("PORT", "8080"),
	}, nil
}

func (c *Config) Log() {
	log.Println("Configuration:")
	log.Printf("  DB Host:     %s", c.DBHost)
	log.Printf("  DB Port:     %s", c.DBPort)
	log.Printf("  DB User:     %s", c.DBUser)
	log.Printf("  DB Password: %s", "***")
	log.Printf("  DB Name:     %s", c.DBName)
	log.Printf("  App Port:    %s", c.Port)
}

func (c *Config) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName,
	)
}

func getEnv(key, fallback string) string {
	v := os.Getenv(key)
	if v != "" {
		return v
	}
	return fallback
}
