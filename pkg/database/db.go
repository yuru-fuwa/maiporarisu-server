package database

import (
	"fmt"
	"sqlite/pkg/config"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(cfg *config.Config) (*gorm.DB, error) {
	user := cfg.DBUser
	password := cfg.DBPassword
	host := cfg.DBHost
	port := cfg.DBPort
	name := cfg.DBName
	//繋がらなかったらsslmodeをrequireにしよう！
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, name)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Print("connect to database")

	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	db.AutoMigrate(&Task{})

	return db, nil
}
