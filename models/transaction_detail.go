package models

type TransactionDetail struct {
	TransactionID uint
	ProductID     uint
	Quantity      int `json:"quantity"`
}
