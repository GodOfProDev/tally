package main

import (
	"github.com/godofprodev/tally/api/config"
	"github.com/godofprodev/tally/api/router"
	"github.com/godofprodev/tally/api/storage"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	dbConfig := config.NewDBConfig()
	serverConfig, err := config.NewServerConfig()
	if err != nil {
		log.Fatal(err)
	}

	mongo := storage.NewMongoStore(dbConfig)
	err = mongo.Connect()
	if err != nil {
		log.Fatal(err)
	}

	r := router.NewRouter(mongo)
	r.RegisterHandlers()

	err = r.Listen(serverConfig)
	if err != nil {
		log.Fatal(err)
	}
}
