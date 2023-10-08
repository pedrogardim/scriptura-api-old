package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	db "github.com/pedrogardim/scriptura-api/database"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.Initialize()
	fmt.Println("Database connection successful")
}

func main() {

	db.TestQuery()
	defer db.Client.Disconnect(db.Ctx)
}
