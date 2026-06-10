package cmd

import (
	t "expense-tracker/internal/expense/types"
	s "expense-tracker/internal/storage"
	u "expense-tracker/internal/utils"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var addExpenseCmd = &cobra.Command{
	Use:   "add",
	Short: "Add expense",
	Long:  "Add expense to the file",
	RunE:   addExpense,
}

func addExpense (cmd *cobra.Command, args []string) error {
	if strings.TrimSpace(description) == "" {
		return fmt.Errorf("description cannot be empty")
	}
	if strings.TrimSpace(category) == "" {
		return fmt.Errorf("category cannot be empty")
	}
	if amount < 0 {
		return fmt.Errorf("amount must be positive, got %.2f", amount)
	}

	existingExpense, err := s.LoadExpenses()
	if err != nil {
		return err
	}

	nextID := 1

	for _, t := range existingExpense {
		if t.ID >= nextID {
			nextID = t.ID + 1
		}
	}

	now := time.Now()

	newExpense := t.ExpenseTracker {
		ID: nextID,
		Date: now.Format("2006-01-02"),
		Category: category,
		Description: description,
		Amount: amount,
	}

	existingExpense = append(existingExpense, newExpense)

	byteData, err := u.EncodeExpenses(existingExpense)
	if err != nil {
		return fmt.Errorf("encoding expenses: %w", err)
	}

	err = os.WriteFile(filename, byteData, permissionCode)
	if err != nil {
    return fmt.Errorf("writing expenses: %w", err)
	}

	fmt.Printf("Expense added successfully (ID: %d)", newExpense.ID)
	return nil
}

