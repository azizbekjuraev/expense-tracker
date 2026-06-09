package utils

import (
	"encoding/json"
	"log"
	t "expense-tracker/internal/expense/types"
)

func EncodeExpenses (e []t.ExpenseTracker) []byte {
	jExpense, err := json.Marshal(e)
	if err != nil {
		log.Fatalf("Error while turning code to byte %v", err)
	}
	return jExpense
}
