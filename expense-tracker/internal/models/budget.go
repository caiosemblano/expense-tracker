package models

import "time"

type Budget struct {
	Month  time.Month `json:"month"`
	Year   int        `json:"year"`
	Amount float64    `json:"amount"`
}

func NewBudget(month time.Month, year int, amount float64) *Budget {
	return &Budget{
		Month:  month,
		Year:   year,
		Amount: amount,
	}
}
