package db

import (
	"context"
	"fmt"
	"log"

	"twitter-uala/config"
	followModels "twitter-uala/internal/domain/follow/models"
	tweetModels "twitter-uala/internal/domain/tweet/models"
	userModels "twitter-uala/internal/domain/user/models"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var CONN *gorm.DB
var REDIS *redis.Client

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

func Redis(redisConfig *redis.Options) *redis.Client {
	REDIS := redis.NewClient(redisConfig)

	if err := REDIS.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}

	log.Println("Connected to Redis successfully")

	return REDIS
}
