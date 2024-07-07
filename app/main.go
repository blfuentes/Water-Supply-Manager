package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	mongo "go.mongodb.org/mongo-driver/mongo"

	handlers "water-supply-manager/handlers"
	dbService "water-supply-manager/services"
)

func GetInvoicesHandler(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		handlers.GetInvoices(client, c)
	}
}

func GetInvoiceHandler(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		handlers.GetInvoice(client, c)
	}
}

func PostInvoiceHandler(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		handlers.PostInvoice(client, c)
	}
}

func main() {
	// This is the main function
	router := gin.Default()

	// Initialize the database
	if dbClient, err := dbService.Init(); err != nil {
		log.Fatal("Error initializing database")
	} else {
		defer func() {
			if err := dbClient.Disconnect(context.TODO()); err != nil {
				panic(err)
			}
		}()

		router.GET("/invoices", GetInvoicesHandler(dbClient))
		router.GET("/invoices/:id", GetInvoiceHandler(dbClient))
		router.POST("/invoices", PostInvoiceHandler(dbClient))

		router.Run("0.0.0.0:8080")
	}

}
