package config

import (
	"errors"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type env struct {
	Port        string
	DBDriver    string
	DBName      string
	RedisConfig *redis.Options
}

var ENV *env

func Load() {
	log.Println("loading app configs")

	ENV = &env{Port: "", DBDriver: "", DBName: ""}

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		var vipErr viper.ConfigFileNotFoundError
		if ok := errors.As(err, &vipErr); ok {
			log.Fatalln(fmt.Errorf("config file not found. %w", err))
		} else {
			log.Fatalln(fmt.Errorf("unexpected error loading config file. %w", err))
		}
	}

	ENV.Port = viper.GetString("server.port")
	ENV.DBDriver = viper.GetString("db.driver")
	ENV.DBName = viper.GetString("db.sqlite")
	redisConfig := &redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
		Protocol: viper.GetInt("redis.protocol"),
	}
	ENV.RedisConfig = redisConfig

	log.Println("app configs loaded")
}
