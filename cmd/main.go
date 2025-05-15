package main

import (
	"fmt"
	"twitter-uala/config"
	"twitter-uala/db"
	"twitter-uala/repositories"
	"twitter-uala/server"

	"twitter-uala/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	config.Load()
	db.Connect()

	repos := repositories.NewRepositories(db.CONN)
	services := domain.NewServices(repos)

	seed := db.NewSeeder(db.CONN)
	seed.Seed()

	server := server.NewHTTPServer(gin.Default(), services, validator.New())
	server.Run(fmt.Sprintf(":%s", config.ENV.Port))
}
