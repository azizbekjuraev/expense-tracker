package cmd

import (
	s "expense-tracker/internal/storage"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var summaryExpensesCmd = &cobra.Command{
	Use:   "summary",
	Short: "Summary of expenses",
	Long:  "Summary of expense totals of amount you spent",
	Run:   showSummary,
}

func showSummary (cmd * cobra.Command, args []string) {
  existingExpenses, err := s.LoadExpenses()
	if err != nil {
		log.Fatalf("Error while extracting existing expenses, %v", err)
	}

	var total = 0.0

	for _, e := range existingExpenses {
		total = total + e.Amount
	}
	fmt.Println("Total expenses:", total)
}
