package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Env struct {
	LocalServerPort string `mapstructure:"LOCAL_SERVER_PORT"`
	MongoDBURL      string `mapstructure:"MONGODB_URL"`
	JWTSecret       string `mapstructure:"JWT_SECRET"`
	DBName          string `mapstructure:"DB_NAME"`
	ContextTimeout  int    `mapstructure:"CONTEXT_TIMEOUT"`
	SMTPUsername    string `mapstructure:"SMTPUsername"`
	SMTPPassword    string `mapstructure:"SMTPPassword"`
	SMTPHost        string `mapstructure:"SMTPHost"`
	SMTPPort        string `mapstructure:"SMTPPort"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
}

func NewEnv() *Env {
	// Load the .env file using godotenv
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	var env Env
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Can't find the file .env: %v", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatalf("Environment can't be loaded: %v", err)
	}

	return &env
}
