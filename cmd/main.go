package main

import (
	"twitter-uala/config"
	"twitter-uala/db"
	"twitter-uala/repositories"

	"twitter-uala/internal/domain"
)

func main() {
	config.Load()
	db.Connect()

	repos := repositories.NewRepositories(db.CONN)
	services := domain.NewServices(repos)

}
