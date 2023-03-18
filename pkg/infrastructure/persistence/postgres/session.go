package postgres

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var gormDbSingleton *gorm.DB

func NewSession() (gormDb *gorm.DB, err error) {
	if gormDbSingleton == nil {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB"),
		)
		gormDbSingleton, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return gormDb, err
		}

		return gormDbSingleton, nil
	}

	return gormDbSingleton, nil
}
