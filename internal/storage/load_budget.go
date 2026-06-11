package storage


import (
	"encoding/json"
	t "expense-tracker/internal/expense/types"
	"log"
	"os"
)

func LoadBudget() (t.Budget, error) {
	jsonData, err := os.ReadFile("budget.json")
	if err != nil {
		log.Fatalf("Error while reading file %v", err)
	}

	if len(jsonData) == 0 {
		return t.Budget{}, nil
	}

	var budget t.Budget

	err = json.Unmarshal(jsonData, &budget)
	if err != nil {
		log.Fatalf("Error while unmarshalling %v", err)
	}

	return budget, err 
}
