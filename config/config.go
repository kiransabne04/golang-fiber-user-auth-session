package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)


type EConfig struct {
	DBurl      string
	JWTsecret  string
	ServerPort string
	TokenTTL   time.Duration
	RefreshTTL time.Duration
}

func LoadConfig() *EConfig {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("../config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading config file: %v", err)
	}

	return &EConfig{
		DBurl:      viper.GetString("db_url"),
		JWTsecret:  viper.GetString("jwt_secret"),
		ServerPort: viper.GetString("server_port"),
		TokenTTL:   time.Second * time.Duration(viper.GetInt("token_ttl")),
		RefreshTTL: time.Second * time.Duration(viper.GetInt("refresh_ttl")),
	}
}