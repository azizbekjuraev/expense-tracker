package cmd

import (
	t "expense-tracker/internal/expense/types"
	s "expense-tracker/internal/storage"
	u "expense-tracker/internal/utils"
	"time"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var addExpenseCmd = &cobra.Command{
	Use:   "add",
	Short: "Add expense",
	Long:  "Add expense to the file",
	Run:   addExpense,
}

func addExpense (cmd *cobra.Command, args []string) {
	existingExpense, err := s.LoadExpenses()
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
		Description: description,
		Amount: amount,
	}

	existingExpense = append(existingExpense, newExpense)

	byteData := u.EncodeExpenses(existingExpense)

	err = os.WriteFile(filename, byteData, permissionCode)

	if err != nil {
		log.Fatalf("Error while writing expenses, %v", err)
	}

	fmt.Printf("Expense added successfully (ID: %d)", newExpense.ID)
}

