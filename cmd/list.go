package cmd

import (
	s "expense-tracker/internal/storage"
	"log"
	"os"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var listExpenseCmd = &cobra.Command{
	Use:   "list",
	Short: "List expenses",
	Long:  "List expenses to the see all expenses",
	Run:   showExpenses,
}

func showExpenses (cmd *cobra.Command, args []string) {
	existingExpenses, err := s.LoadExpenses()

	if err != nil {
		log.Fatalf("Error while reading existing expenses %v", err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.Header("ID", "Date", "Category", "Description", "Amount")
	table.Bulk(existingExpenses)
	table.Render()
}
