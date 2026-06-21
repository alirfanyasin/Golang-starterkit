package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName      string
	AppEnv       string
	AppDebug     bool
	AppPort      string
	AppURL       string
	AppTimezone  string

	DBConnection     string
	DBHost           string
	DBPort           string
	DBDatabase       string
	DBUsername       string
	DBPassword       string
	DBSSLMode        string
	DBMaxIdleConns   int
	DBMaxOpenConns   int
	DBConnMaxLifetime time.Duration

	JWTSecret          string
	JWTAccessTokenTTL  time.Duration
	JWTRefreshTokenTTL time.Duration

	LogLevel  string
	LogFormat string
}

var AppConfig *Config

// LoadConfig loads the environment variables into the AppConfig struct
func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found or failed to read, using system environment variables")
	}

	AppConfig = &Config{
		AppName:      getEnv("APP_NAME", "Golang Starterkit"),
		AppEnv:       getEnv("APP_ENV", "local"),
		AppDebug:     getEnvAsBool("APP_DEBUG", true),
		AppPort:      getEnv("APP_PORT", "8080"),
		AppURL:       getEnv("APP_URL", "http://localhost:8080"),
		AppTimezone:  getEnv("APP_TIMEZONE", "UTC"),

		DBConnection:     getEnv("DB_CONNECTION", "sqlite"),
		DBHost:           getEnv("DB_HOST", "127.0.0.1"),
		DBPort:           getEnv("DB_PORT", "5432"),
		DBDatabase:       getEnv("DB_DATABASE", "golang_starterkit.db"),
		DBUsername:       getEnv("DB_USERNAME", "postgres"),
		DBPassword:       getEnv("DB_PASSWORD", ""),
		DBSSLMode:        getEnv("DB_SSL_MODE", "disable"),
		DBMaxIdleConns:   getEnvAsInt("DB_MAX_IDLE_CONNS", 10),
		DBMaxOpenConns:   getEnvAsInt("DB_MAX_OPEN_CONNS", 100),
		DBConnMaxLifetime: getEnvAsDuration("DB_CONN_MAX_LIFETIME", time.Hour),

		JWTSecret:          getEnv("JWT_SECRET", "secret"),
		JWTAccessTokenTTL:  getEnvAsDuration("JWT_ACCESS_TOKEN_TTL", 15*time.Minute),
		JWTRefreshTokenTTL: getEnvAsDuration("JWT_REFRESH_TOKEN_TTL", 7*24*time.Hour),

		LogLevel:  getEnv("LOG_LEVEL", "debug"),
		LogFormat: getEnv("LOG_FORMAT", "text"),
	}

	return AppConfig
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	valStr := getEnv(key, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valStr := getEnv(key, "")
	if val, err := strconv.Atoi(valStr); err == nil {
		return val
	}
	return defaultValue
}

func getEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	valStr := getEnv(key, "")
	if val, err := time.ParseDuration(valStr); err == nil {
		return val
	}
	return defaultValue
}
