package main

import (
	"twitter-uala/config"
	"twitter-uala/internal/db"
	"twitter-uala/internal/domain"
	"twitter-uala/internal/repositories"
)

func main() {
	config.Load()
	db.Connect()

	repos := repositories.NewRepositories(db.CONN)
	services := domain.NewServices(repos)

}
