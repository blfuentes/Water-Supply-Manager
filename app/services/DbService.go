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

func Init() (*mongo.Client, error) {
	// load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
		return nil, err
	} else {
		log.Println("Environment variables loaded")
	}

	// get the MongoDB URI from the environment
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. " +
			"See: " +
			"www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
		return nil, fmt.Errorf("Set your 'MONGODB_URI' environment variable")
	} else {
		log.Println("MONGODB_URI:", uri)
	}

	// connect to MongoDB for verification
	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	return client, nil
}

func GetInvoices(client *mongo.Client) ([]models.InvoiceDto, error) {
	log.Println("Service::Getting invoices")

	collection := client.Database("local").Collection("Invoices")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err == mongo.ErrNoDocuments {
		log.Println("No documents found")
		return nil, nil
	}
	if err != nil {
		panic(err)
	}
	defer cursor.Close(context.TODO())

	var results []models.InvoiceDto
	for cursor.Next(context.TODO()) {
		var elem models.Invoice
		err := cursor.Decode(&elem)
		if err != nil {
			return nil, err
		}
		results = append(results, elem.ToDto())
	}

	return results, nil
}

func GetInvoice(client *mongo.Client, id int64) (models.InvoiceDto, error) {
	log.Println("Service::Getting invoice")

	collection := client.Database("local").Collection("Invoices")
	filter := bson.D{{"ID", id}}
	var resultdb models.Invoice
	err := collection.FindOne(context.TODO(), filter).Decode(&resultdb)
	if err == mongo.ErrNoDocuments {
		log.Println("No documents found")
		return models.InvoiceDto{}, nil
	}
	if err != nil {
		panic(err)
	}

	return resultdb.ToDto(), nil
}

func PostInvoice(client *mongo.Client, invoice models.InvoiceDto) error {
	log.Println("Service::Posting invoice")

	collection := client.Database("local").Collection("Invoices")
	_, err := collection.InsertOne(context.TODO(), invoice.ToModel())
	if err != nil {
		panic(err)
	}

	return nil
}

func UpdateInvoice(client *mongo.Client, id int64, invoice models.InvoiceDto) error {
	log.Println("Service::Updating invoice")

	collection := client.Database("local").Collection("Invoices")
	filter := bson.D{{"ID", id}}
	update := bson.D{{"$set", invoice.ToModel()}}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}

	return nil

}

func DeleteInvoice(client *mongo.Client, id int64) error {
	log.Println("Service::Deleting invoice")

	collection := client.Database("local").Collection("Invoices")
	filter := bson.D{{"ID", id}}
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	return nil
}
