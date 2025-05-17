package config

import (
	"errors"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type env struct {
	Port     string `mapstructure:"PORT"`
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBName   string `mapstructure:"DB_NAME"`
}

var ENV *env

func Load() {
	log.Println("loading app configs")

	ENV = &env{Port: "", DBDriver: "", DBName: ""}

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("../")

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

	log.Println("app configs loaded")
}
