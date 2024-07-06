package services

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	models "water-supply-manager/models"
)

func Connect() error {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	} else {
		log.Println("Environment variables loaded")
		// fmt.Println("MONGODB_URI:", os.Getenv("MONGODB_URI"))
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. " +
			"See: " +
			"www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	} else {
		log.Println("MONGODB_URI:", uri)
	}
	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	return nil
}

func GetInvoices() ([]models.Invoice, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
		fmt.Println("env file not found")
	} else {
		log.Println("Environment variables loaded")
		fmt.Println("MONGODB_URI:", os.Getenv("MONGODB_URI"))
	}
	uri := os.Getenv("MONGODB_URI")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println("Error connecting to MongoDB")
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("local").Collection("Invoices")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err == mongo.ErrNoDocuments {
		fmt.Println("No documents found")
		return nil, nil
	}
	if err != nil {
		panic(err)
	}
	defer cursor.Close(context.TODO())

	var results []models.Invoice
	for cursor.Next(context.TODO()) {
		var elem models.Invoice
		err := cursor.Decode(&elem)
		if err != nil {
			return nil, err
		}
		results = append(results, elem)
	}

	return results, nil
}
