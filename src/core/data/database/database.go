package database

import (
	"errors"
	"log"
	"os"
	"time"

	"gwi/assignment/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	ErrSQLConnection = errors.New("sql connection timeout")
)

var dbConnection *gorm.DB
var connectionString string

func Initialize(connectionStr string) {
	connectionString = connectionStr
}

func GetGormConnection() (*gorm.DB, error) {
	if dbConnection == nil {
		connection, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				logger.Config{
					SlowThreshold:             time.Second,
					LogLevel:                  logger.Info,
					IgnoreRecordNotFoundError: true,
					Colorful:                  true,
				},
			),
		})

		if err != nil {
			return nil, ErrSQLConnection
		}

		dbConnection = connection

		sql.Initialize(dbConnection)
	}

	return dbConnection, nil
}
