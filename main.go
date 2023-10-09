package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	routes "github.com/pedrogardim/scriptura-api/api/rest"
	db "github.com/pedrogardim/scriptura-api/database"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.Initialize()
}

func main() {
	defer db.Client.Disconnect(db.Ctx)

	routes.Init()

	fmt.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
