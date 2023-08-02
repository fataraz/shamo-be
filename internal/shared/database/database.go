package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database ...
type Database struct {
	*gorm.DB
}

// New init database
func New(config ConfigDatabase) (*Database, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%v sslmode=disable search_path=%s",
		config.Username,
		config.Password,
		config.Name,
		config.Host,
		config.Port,
		config.Schema)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database")
		return nil, err
	}

	s := &Database{db}

	return s, nil
}
