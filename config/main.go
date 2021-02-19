package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"log"
)

var App = AppConfig{}

func LoadConfig(){

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env.Parse(&App)
}
