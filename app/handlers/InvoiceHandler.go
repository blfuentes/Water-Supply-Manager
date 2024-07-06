package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	models "water-supply-manager/models"
)

var invoices = []models.Invoice{
	{ID: 1, DateFrom: "2019-01-01", DateTo: "2019-01-31", Amount: 1000.00},
	{ID: 2, DateFrom: "2019-02-01", DateTo: "2019-02-28", Amount: 2000.00},
}

// GetInvoices returns the list of all invoices
func GetInvoices(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, invoices)
}

// GetInvoice returns a single invoice
func GetInvoice(c *gin.Context) {
	if id, err := strconv.ParseInt(c.Param("id"), 10, 0); err == nil {
		for _, invoice := range invoices {
			if invoice.ID == id {
				c.IndentedJSON(http.StatusOK, invoice)
				return
			}
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "invoice not found"})

}

// post a new invoice
func PostInvoice(c *gin.Context) {
	var newInvoice models.Invoice
	if err := c.BindJSON(&newInvoice); err != nil {
		return
	}
	invoices = append(invoices, newInvoice)
	c.IndentedJSON(http.StatusCreated, newInvoice)
}
