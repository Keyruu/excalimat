package config

import "os"

var (
	// database
	DbHost     = Config("DB_HOST")
	DbName     = Config("DB_NAME")
	DbPassword = Config("DB_PASSWORD")
	DbUser     = Config("DB_USER")
	DbPort     = Config("DB_PORT")
	// jwt
	JwtKeyUrl  = Config("JWT_KEY_URL")
	AdminGroup = Config("ADMIN_GROUP")
	UserGroup  = Config("USER_GROUP")
)

// Config func to get env value
func Config(key string) string {
	return os.Getenv(key)
}
