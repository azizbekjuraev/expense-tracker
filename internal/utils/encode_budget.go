package utils

import (
	"encoding/json"
	t "expense-tracker/internal/expense/types"
)

func EncodeBudget (e t.Budget) ([]byte, error) {
	jBudget, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}

	return jBudget, nil
}
