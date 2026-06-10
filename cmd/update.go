package cmd

import (
	s "expense-tracker/internal/storage"
	u "expense-tracker/internal/utils"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var updateExpenseCmd = &cobra.Command{
	Use: "update",
	Short: "Update expense",
	Long: "Update the expense you want",
	Run: updateExpense,
}

func updateExpense (cmd *cobra.Command, args []string) {
	existingExpenses, err := s.LoadExpenses()
	if err != nil {
		log.Fatalf("Error while reading expenses. %v", err)
	}

	now := time.Now()

	for i, e := range existingExpenses {
		if e.ID == id {
			if amount > 0 {
				existingExpenses[i].Amount = amount
			}
			if description != "" {
				existingExpenses[i].Description = description
			}
			if category != "" {
				existingExpenses[i].Category = category
			}
		  existingExpenses[i].Date = now.Format("2006-01-02")
		}
	}

	byteData, err := u.EncodeExpenses(existingExpenses)
	if err != nil {
		log.Fatalf("Error while encoding expenses: %v", err)
	}

	err = os.WriteFile(filename, byteData, permissionCode)
	if err != nil {
		log.Fatalf("Error while writing updated data to the file %v", err)
	}

	fmt.Println("Expense updated successfully")
}
