package gorm_postgesql

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type pgConfig struct {
	user     string
	password string
	host     string
	port     string
	database string
}

func NewConfig(user, password, host, port, database string) *pgConfig {
	return &pgConfig{
		user:     user,
		password: password,
		host:     host,
		port:     port,
		database: database,
	}
}

func NewClient(cfg *pgConfig) *gorm.DB {
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=Europe/Moscow",
		cfg.user, cfg.password, cfg.host, cfg.port, cfg.database,
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
