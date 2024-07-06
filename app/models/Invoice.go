package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Invoice struct {
	ID       int64              `bson:"ID"`
	DateFrom primitive.DateTime `bson:"DateFrom"`
	DateTo   primitive.DateTime `bson:"DateTo"`
	Amount   float64            `bson:"Amount"`
}
