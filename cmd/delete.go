package cmd

import (
	s "expense-tracker/internal/storage"
	u "expense-tracker/internal/utils"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var deleteExpenseCmd = &cobra.Command{
	Use: "delete",
	Short: "Delete expense",
	Long: "Delete some expenses you do not need",
	Run: deleteExpense,
}

var deleteAllExpensesCmd = &cobra.Command{
	Use: "delete-all",
	Short: "Delete all expenses",
	Long: "Delete all expenses",
	Run: deleteAllExpenses,
}

func deleteExpense (cmd * cobra.Command, args []string) {
	existingExpenses, err := s.LoadExpenses()
	if err != nil {
		log.Fatalf("Error while reading expenses. %v", err)
	}
  filtered := existingExpenses[:0]

	for _, e := range existingExpenses {
		if e.ID != id {
			filtered = append(filtered, e)
		}
	}
	existingExpenses = filtered
	byteData := u.EncodeExpenses(existingExpenses)

	err = os.WriteFile(filename, byteData, permissionCode)

	if err != nil {
		log.Fatalf("Error while writing updated data to the file %v", err)
	}
	fmt.Println("Expense deleted successfully")
}

func deleteAllExpenses (cmd *cobra.Command, args []string) {
	existingExpenses, err := s.LoadExpenses()

	if err != nil {
		log.Fatalf("Error while reading expenses. %v", err)
	}

	if force {
		existingExpenses = existingExpenses[:0]
	}

	byteData := u.EncodeExpenses(existingExpenses)

	err = os.WriteFile(filename, byteData, permissionCode)

	if err != nil {
		log.Fatalf("Error while writing updated data to the file %v", err)
	}

	fmt.Println("All expenses deleted successfully")
}
