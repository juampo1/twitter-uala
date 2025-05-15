package db

import (
	"fmt"
	"log"

	"twitter-uala/config"
	followModels "twitter-uala/internal/domain/follow/models"
	tweetModels "twitter-uala/internal/domain/tweet/models"
	userModels "twitter-uala/internal/domain/user/models"

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
	err = CONN.AutoMigrate(&userModels.User{}, &tweetModels.Tweet{}, &followModels.Follow{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	CONN.Exec("PRAGMA foreign_keys = ON;")

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
