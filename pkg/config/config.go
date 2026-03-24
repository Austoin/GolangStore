package config

import "os"

type App struct {
	Env  string
	Name string
}

type MySQL struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type Redis struct {
	Host     string
	Port     string
	Password string
}

type Config struct {
	App   App
	MySQL MySQL
	Redis Redis
}

func Load() Config {
	return Config{
		App: App{
			Env:  getEnv("APP_ENV", "local"),
			Name: getEnv("APP_NAME", "GolangStore"),
		},
		MySQL: MySQL{
			Host:     getEnv("MYSQL_HOST", "127.0.0.1"),
			Port:     getEnv("MYSQL_PORT", "3306"),
			User:     getEnv("MYSQL_USER", "root"),
			Password: getEnv("MYSQL_PASSWORD", "root"),
			Database: getEnv("MYSQL_DATABASE", "golang_store"),
		},
		Redis: Redis{
			Host:     getEnv("REDIS_HOST", "127.0.0.1"),
			Port:     getEnv("REDIS_PORT", "6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
		},
	}
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}
