package config

import "os"

var (
	// database
	DB_HOST     = Config("DB_HOST")
	DB_NAME     = Config("DB_NAME")
	DB_PASSWORD = Config("DB_PASSWORD")
	DB_USER     = Config("DB_USER")
	DB_PORT     = Config("DB_PORT")
	// jwt
	JWT_KEY_URL = Config("JWT_KEY_URL")
	ADMIN_GROUP = Config("ADMIN_GROUP")
	USER_GROUP  = Config("USER_GROUP")
)

// Config func to get env value
func Config(key string) string {
	return os.Getenv(key)
}
