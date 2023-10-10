package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
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
	r := gin.Default()
	routes.Init(r)
	fmt.Println("Listening on port 8080")
	r.Run()
}
