package env

import (
	"github.com/joho/godotenv"
	"os"
)

type Vars struct {
	DbUri            string
	Port             string
	Host             string
	IsDBDebugEnabled bool
	JWTSecretKey     string
}

var configs = &Vars{}
var isLoaded = false

func GetEnvVars() *Vars {
	if isLoaded {
		return configs
	}
	_ = godotenv.Load()
	configs = &Vars{
		os.Getenv("DATABASE_URL"),
		os.Getenv("PORT"),
		os.Getenv("HOST"),
		os.Getenv("DB_DEBUG_ENABLED") == "1",
		os.Getenv("JWT_TOKEN_KEY"),
	}
	isLoaded = true
	return configs
}
