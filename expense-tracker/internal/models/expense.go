package models

import (
	"time"
)

// Expense representa uma despesa individual
type Expense struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Date        time.Time `json:"date"`
	Category    string    `json:"category"`
}

// NewExpense cria uma nova despesa
func NewExpense(description string, amount float64, category string) *Expense {
	return &Expense{
		Description: description,
		Amount:      amount,
		Date:        time.Now(),
		Category:    category,
	}
}
