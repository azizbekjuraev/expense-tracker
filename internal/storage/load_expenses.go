package storage

import (
	"encoding/json"
	t "expense-tracker/internal/expense/types"
	"log"
	"os"
)

func LoadExpenses() ([]t.ExpenseTracker, error) {
	jsonData, err := os.ReadFile("expenses.json")
	if err != nil {
		log.Fatalf("Error while reading file %v", err)
	}

	if len(jsonData) == 0 {
		return []t.ExpenseTracker{}, nil
	}

	var expenses []t.ExpenseTracker

	err = json.Unmarshal(jsonData, &expenses)
	if err != nil {
		log.Fatalf("Error while unmarshalling %v", err)
	}

	return expenses, err
}
