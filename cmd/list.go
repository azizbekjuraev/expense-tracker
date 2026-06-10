package cmd

import (
	s "expense-tracker/internal/storage"
	"os"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var listExpenseCmd = &cobra.Command{
	Use:   "list",
	Short: "List expenses",
	Long:  "List expenses to the see all expenses",
	RunE:   showExpenses,
}

func showExpenses (cmd *cobra.Command, args []string) error {
	existingExpenses, err := s.LoadExpenses()
	if err != nil {
		return err
	}

	filtered := existingExpenses[:0]
	for _, e := range existingExpenses {
		if (from == "" || e.Date >= from) && (to == "" || e.Date <= to) {
			filtered = append(filtered, e)
		}
	}
	existingExpenses = filtered
	
	table := tablewriter.NewWriter(os.Stdout)
	table.Header("ID", "Date", "Category", "Description", "Amount")
	table.Bulk(existingExpenses)
	table.Render()

	return nil
}
