package models

type Invoice struct {
	ID       int64
	DateFrom string
	DateTo   string
	Amount   float64
}
