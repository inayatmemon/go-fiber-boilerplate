package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// ServerConfig - Server configuration
type ServerConfig struct {
	AppName   string
	Port      int
	ServerUri string
	LogLevel  string
}

// Config structure
type Config struct {
	Server ServerConfig
}

// AppConfig - Appconfig object
var AppConfig = &Config{
	Server: ServerConfig{
		AppName:   "Fiber-FrontEnd",
		Port:      3009,
		ServerUri: "http://127.0.0.1/",
		LogLevel:  "DEBUG",
	},
}

// LoadEnv - function load Enviroment variable from .env file
func LoadEnv() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	config := &Config{
		Server: ServerConfig{
			AppName:   getEnv("APP_NAME", ""),
			Port:      getEnvAsInt("API_PORT", 3009),
			ServerUri: getEnv("SERVER_URI", ""),
			LogLevel:  getEnv("LOG_LEVEL", "DEBUG"),
		},
	}
	AppConfig = config
	return config
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// Simple helper function to read an environment variable into integer or return a default value
func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

// Helper to read an environment variable into a bool or return default value
func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}
