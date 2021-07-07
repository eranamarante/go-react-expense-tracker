package database

import (
	"context"
	"fmt"
	"log"

	"github.com/eranamarante/go-react-expense-tracker/server/helper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Collection {
	config := helper.GetConfiguration()

	// Set client options
	clientOptions := options.Client().ApplyURI(config.ConnectionString)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("go-expense-tracker").Collection("expenses")

	return collection
}
