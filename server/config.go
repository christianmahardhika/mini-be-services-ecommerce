package server

import (
	"log"
	"os"

	"github.com/christianmahardhika/mini-be-services-ecommerce/config"
	"github.com/joho/godotenv"
)

func InitiateConfig() config.AppConfig {
	// InitiateConfig is a function that initiates the configuration.
	if os.Getenv("APP_ENV") == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	var appConfig config.AppConfig
	// Load Database configuration
	appConfig.DBConfig.Host = os.Getenv("DB_HOST")
	appConfig.DBConfig.Port = os.Getenv("DB_PORT")
	appConfig.DBConfig.User = os.Getenv("DB_USER")
	appConfig.DBConfig.Pass = os.Getenv("DB_PASS")
	appConfig.DBConfig.Name = os.Getenv("DB_NAME")
	appConfig.DBConfig.SSLMode = os.Getenv("DB_SSL_MODE")

	// load cors configuration
	appConfig.AllowedOrigins = []string{os.Getenv("ALLOWED_ORIGINS")}
	appConfig.AllowedHeaders = []string{os.Getenv("ALLOWED_HEADERS")}

	return appConfig
}
