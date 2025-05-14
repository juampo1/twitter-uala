package main

import (
	"twitter-uala/config"
	"twitter-uala/internal/db"
)

func main() {
	config.Load()
	db.Connect()
}
