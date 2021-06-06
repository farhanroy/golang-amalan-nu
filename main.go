package main

import (
	"os"
	"log"

	"github.com/joho/godotenv"
	"github.com/farhanroy/api-amalan-nahdliyin/src/routes"
)

func main() {

	err := godotenv.Load()
	e := routes.Init()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e.Logger.Fatal(e.Start(":"+ os.Getenv("PORT")))
}
