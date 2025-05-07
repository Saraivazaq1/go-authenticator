package database

import (
	"goauthenticator/environment"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConectarDB() {
	var err error

	dsn := environment.GetDSN()

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		log.Fatal("Não foi possível se conectar ao banco de dados")
	}
}