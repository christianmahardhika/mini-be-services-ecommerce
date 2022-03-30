package server

import (
	"fmt"
	"log"

	"github.com/christianmahardhika/mini-be-services-ecommerce/config"
	"github.com/christianmahardhika/mini-be-services-ecommerce/domain/products"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

// Initialize database connection

func InitDb(config config.DBConfig) {
	var err error

	// open connection to database
	dbString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s", config.Host, config.User, config.Pass, config.Name, config.SSLMode)
	dbConn, err = gorm.Open(postgres.Open(dbString))

	if err != nil {
		log.Panicf("Error while connecting to DB, err: %s", err)
	}

	migrateDb()

}

func GetDbConnection(config config.DBConfig) *gorm.DB {
	if dbConn == nil {
		InitDb(config)
	}
	return dbConn
}

func migrateDb() {
	err := dbConn.AutoMigrate(&products.Products{})
	if err != nil {
		log.Panicf("Error while migrating DB, err: %s", err)
	}
}
