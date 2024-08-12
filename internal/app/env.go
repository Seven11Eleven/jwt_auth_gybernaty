package app

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	DBName                 string `mapstructure:"DATABASE_NAME"`
	DBPass                 string `mapstructure:"DATABASE_PASSWORD"`
	DBUser                 string `mapstructure:"DATABASE_USER"`
	DBHost                 string `mapstructure:"DATABASE_HOST"`
	DBPort                 string `mapstructure:"DATABASE_PORT"`
	ServerPort             string `mapstructure:"PORT"`
	JWTSecret              string `mapstructure:"JWT_SECRET"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_EXPIRY_HOUR"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf(".env not found: %v ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatalf("env couldn't be loaded: %v", err)
	}
	return &env
}
