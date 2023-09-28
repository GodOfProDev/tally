package main

import (
	"fmt"
	"github.com/godofprodev/tally/api/config"
	"github.com/godofprodev/tally/api/storage"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.NewDBConfig()

	mongo := storage.NewMongoStore(cfg)

	err = mongo.Connect()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(mongo.Client)
}
