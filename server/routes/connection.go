package routes

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client{
	err := godotenv.Load(".env")

	if err != nil{
		log.Fatal("Error handling .env file")
	}


	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)


	MongoDB := os.Getenv("MONGODB_URL")
	client, err := mongo.Connect(ctx,options.Client().ApplyURI(MongoDB))
	if err != nil{
		log.Fatal("Error handling MongoDB URL" + err.Error())
	}

	fmt.Println("Connected to MONGODB!")
	
	defer cancel()
	return client
}

var Client *mongo.Client = DBinstance() 

func OpenCollention(client *mongo.Client, collectionName string) *mongo.Collection{
	return client.Database(os.Getenv("DATABASE_NAME")).Collection(collectionName)
}