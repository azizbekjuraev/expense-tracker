package types

type ExpenseTracker struct {
	ID int `json:"id"`
	Date string `json:"date"`
	Description string `json:"description"`
	Amount int `json:"amount"`
}
