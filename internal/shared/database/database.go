package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database ...
type Database struct {
	*gorm.DB
}

// New init database
func New() *Database {
	dsn := "host=localhost user=postgres password=Alesha2021 dbname=shamobe port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf(err.Error())
	}
	return &Database{db}
}
