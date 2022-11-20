package postgres

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var postgresClient *gorm.DB

func NewPostgresClient() (*gorm.DB, error) {
	if postgresClient == nil {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB"),
		)
		postgresClient, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return &gorm.DB{}, err
		}

		return postgresClient, nil
	}

	return postgresClient, nil
}
