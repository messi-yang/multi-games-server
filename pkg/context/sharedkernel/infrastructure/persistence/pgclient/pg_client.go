package pgclient

import (
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbSingleton *gorm.DB = nil
var dbCreatorLock = &sync.Mutex{}

func NewPgClient() (gormDb *gorm.DB) {
	dbCreatorLock.Lock()
	defer dbCreatorLock.Unlock()

	if dbSingleton == nil {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB"),
		)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("Failed to open postgres connection")
		}

		dbSingleton = db
		return dbSingleton
	}

	return dbSingleton
}
