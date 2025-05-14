package db

import (
	"fmt"
	"log"

	"twitter-uala/config"
	"twitter-uala/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var CONN *gorm.DB

func Connect() *gorm.DB {
	if CONN != nil {
		return CONN
	}

	log.Println("Database connection started")
	//run migration

	var err error
	if CONN, err = gorm.Open(sqlite.Open(getDataBaseName()), &gorm.Config{}); err != nil {
		panic(fmt.Errorf("failed to open the database connection. %w", err))
	}

	// Migrar las estructuras a la base de datos
	err = CONN.AutoMigrate(&models.User{}, &models.Tweet{}, &models.Follow{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	log.Println("Database migrated successfully")

	return CONN
}

func getDataBaseName() string {
	dbName := config.ENV.DBName
	if dbName == "" {
		log.Fatalln("database name is not defined to open a connection")
	}

	return dbName
}
