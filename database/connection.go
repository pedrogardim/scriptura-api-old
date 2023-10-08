package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Ctx = context.TODO()
var Db *mongo.Database
var Coll *mongo.Collection

func Initialize() {
	var err error
	mongoUri := os.Getenv("MONGO_URI")
	clientOptions := options.Client().ApplyURI(mongoUri)
	Client, err = mongo.Connect(Ctx, clientOptions)

	if err != nil {
		log.Fatal("Database connection failed")
	}

	err = Client.Ping(Ctx, nil)

	if err != nil {
		log.Fatal("Database ping error")
	}
	Db = Client.Database("scriptura-api")
	Coll = Db.Collection("test")
}

