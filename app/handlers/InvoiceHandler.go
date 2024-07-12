package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	models "water-supply-manager/models"
	services "water-supply-manager/services"
)

// GetInvoices returns the list of all invoices
func GetInvoices(client *mongo.Client, c *gin.Context) {
	if dbinvoices, err := services.GetInvoices(client); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	} else {
		c.IndentedJSON(http.StatusOK, dbinvoices)
	}
}

// GetInvoice returns a single invoice
func GetInvoice(client *mongo.Client, c *gin.Context) {
	if id, err := strconv.ParseInt(c.Param("id"), 10, 0); err == nil {
		if invoice, err := services.GetInvoice(client, id); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		} else {
			c.IndentedJSON(http.StatusOK, invoice)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "invoice not found"})

}

// post a new invoice
func PostInvoice(client *mongo.Client, c *gin.Context) {
	var newInvoice models.InvoiceDto
	if err := c.BindJSON(&newInvoice); err != nil {
		return
	}

	if err := services.PostInvoice(client, newInvoice); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newInvoice)
}

// update an invoice
func UpdateInvoice(client *mongo.Client, c *gin.Context) {
	if id, err := strconv.ParseInt(c.Param("id"), 10, 0); err == nil {

		var updatedInvoice models.InvoiceDto
		if err := c.BindJSON(&updatedInvoice); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
			return
		}

		if updatedInvoice.ID != id {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID in request does not match ID in body"})
			return
		}

		if err := services.UpdateInvoice(client, id, updatedInvoice); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.IndentedJSON(http.StatusOK, updatedInvoice)
	}
}

// delete an invoice
func DeleteInvoice(client *mongo.Client, c *gin.Context) {
	if id, err := strconv.ParseInt(c.Param("id"), 10, 0); err == nil {
		if err := services.DeleteInvoice(client, id); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{"message": "invoice deleted"})
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "invoice not found"})
}
