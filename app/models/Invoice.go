package models

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InvoiceDto struct {
	ID       int64
	DateFrom string
	DateTo   string
	Amount   float64
}

type Invoice struct {
	ID       int64              `bson:"ID"`
	DateFrom primitive.DateTime `bson:"DateFrom"`
	DateTo   primitive.DateTime `bson:"DateTo"`
	Amount   float64            `bson:"Amount"`
}

// ToModel converts the Invoice DTO to a model
func (i InvoiceDto) ToModel() Invoice {
	datefrom, _ := time.Parse(time.RFC3339, i.DateFrom)
	dateto, _ := time.Parse(time.RFC3339, i.DateTo)
	fmt.Println(datefrom)
	fmt.Println(dateto)
	return Invoice{
		ID:       i.ID,
		DateFrom: primitive.NewDateTimeFromTime(datefrom),
		DateTo:   primitive.NewDateTimeFromTime(dateto),
		Amount:   i.Amount,
	}
}

// ToDto converts the Invoice model to a DTO
func (i Invoice) ToDto() InvoiceDto {
	return InvoiceDto{
		ID:       i.ID,
		DateFrom: i.DateFrom.Time().String(),
		DateTo:   i.DateTo.Time().String(),
		Amount:   i.Amount,
	}
}
