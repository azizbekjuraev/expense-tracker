package cmd

import "github.com/spf13/cobra"

const filename = "expenses.json"
const permissionCode = 0644

var (
	description string
	amount float64
	id int
	force bool
	category string
	from string
	to string
	min float64
	monthly float64
)

var rootCmd = &cobra.Command {
	Use: "expense-tracker",
	Short: "Control your expenses",
	Long: "Add, Edit, Summarize your expenses",
}

func init() {
	addExpenseCmd.Flags().StringVar(&category, "category", "", "Category of the expense")
	addExpenseCmd.Flags().StringVar(&description, "description", "", "Description of the expense")
	addExpenseCmd.Flags().Float64Var(&amount, "amount", 0.0, "Amount of the expense")
	deleteExpenseCmd.Flags().IntVar(&id, "id", 0, "Id to delete expense")
	deleteAllExpensesCmd.Flags().BoolVarP(&force, "force", "f", false, "Delete all expenses")
	updateExpenseCmd.Flags().IntVar(&id, "id", 0, "Id to update expense")
	updateExpenseCmd.Flags().Float64Var(&amount, "amount", 0.0, "Update amount of particular expenses")
	updateExpenseCmd.Flags().StringVar(&description, "description", "", "Update description of particular expense")
	listExpenseCmd.Flags().StringVar(&from, "from", "", "Target date from (format: YYYY-MM-DD)")
	listExpenseCmd.Flags().StringVar(&to, "to", "", "Target date to (format: YYYY-MM-DD)")
	listExpenseCmd.Flags().StringVar(&category, "category", "", "Filter by category")
	setBudgetCmd.Flags().Float64Var(&monthly, "monthly", 0.0, "Set your bugdet")
	listExpenseCmd.Flags().Float64Var(&min, "min", 0.0, "Filter by minimum amount")
	rootCmd.AddCommand(addExpenseCmd, listExpenseCmd, deleteExpenseCmd, summaryExpensesCmd, updateExpenseCmd, deleteAllExpensesCmd, setBudgetCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
