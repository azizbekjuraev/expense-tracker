package utils

import (
	"encoding/json"
	t "expense-tracker/internal/expense/types"
)

func EncodeExpenses (e []t.ExpenseTracker) ([]byte, error) {
	jExpense, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	return jExpense, nil
}
