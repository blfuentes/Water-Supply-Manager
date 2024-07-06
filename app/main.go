package main

import (
	"github.com/gin-gonic/gin"

	handlers "water-supply-manager/handlers"
)

func main() {
	// This is the main function
	router := gin.Default()
	router.GET("/invoices", handlers.GetInvoices)
	router.GET("/invoices/:id", handlers.GetInvoice)
	router.POST("/invoices", handlers.PostInvoice)

	router.Run("0.0.0.0:8080")
}
