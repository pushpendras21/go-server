package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type envConfig struct {
	AppPort string
}

func (e *envConfig) LoadConfig() {
	err := godotenv.Load()

	if err != nil {
		log.Panic("Error ENV file not loading.")
	}
	val, ok := os.LookupEnv("APP_PORT")
	if !ok {
		log.Panic("APP_PORT is not loaded in ENV")

	}
	e.AppPort = val

}

var Config envConfig

func init() {
	Config.LoadConfig()
}
