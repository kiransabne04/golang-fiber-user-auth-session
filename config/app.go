package config

import (
	"errors"
	"fiber-user-auth-session/pkg"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AppConfig struct {
	EnvConfig *EConfig
	Database  *pgxpool.Pool
}

// Initialize app config initializes & load env configs later setting up database connections
func InitializeAppConfig(_ string) (*AppConfig, error) {
	//load config
	config := LoadConfig()
	if config.DBurl == "" {
		log.Fatalf("failed loading configuration")
		return nil, errors.New("Failed loading configuration")
	}

	//connect to database
	db, err := pkg.ConnectToDB(config.DBurl)
	if err != nil {
		log.Fatalf("failed to connect to database : %v", err)
		return nil, err
	}

	//services := internal.NewAppServices()

	return &AppConfig{
		EnvConfig: config,
		Database:  db,
	}, nil
}
