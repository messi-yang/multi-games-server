package pgmodel

import (
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB = nil
var dbCreatorLock = &sync.Mutex{}

func NewClient() (gormDb *gorm.DB, err error) {
	dbCreatorLock.Lock()
	defer dbCreatorLock.Unlock()

	if db == nil {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB"),
		)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return gormDb, err
		}

		return db, nil
	}

	return db, nil
}
