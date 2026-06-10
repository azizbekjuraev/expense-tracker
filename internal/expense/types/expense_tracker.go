package types

type ExpenseTracker struct {
	ID int `json:"id"`
	Date string `json:"date"`
	Category string `json:"category"`
	Description string `json:"description"`
	Amount float64 `json:"amount"`
}
