package postgres

import (
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbClient *gorm.DB = nil
var dbClientCreatorLock = &sync.Mutex{}

func NewDbClient() (gormDb *gorm.DB, err error) {
	dbClientCreatorLock.Lock()
	defer dbClientCreatorLock.Unlock()

	if dbClient == nil {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB"),
		)
		dbClient, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return gormDb, err
		}

		return dbClient, nil
	}

	return dbClient, nil
}
