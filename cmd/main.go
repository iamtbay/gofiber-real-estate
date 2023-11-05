package main

import (
	"log"

	"github.com/iamtbay/real-estate-api/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	//Start web server !
	database.Init()
	//

	StartServer()
	//
}
