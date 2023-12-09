package entity

import "time"

type Expenses struct {
	ID              string    `json:"id"`
	Date            time.Time `json:"date"`
	Amount          int       `json:"amount"`
	TransactionType string    `json:"transactionType"`
	Balance         int       `json:"balance,omitempty"`
	Description     string    `json:"description"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}
