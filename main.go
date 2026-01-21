package main

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var (
	description string
	amount int
)

type ExpenseTracker struct {
	ID int `json:"id"`
	Date time.Time `json:"date"`
	Description string `json:"description"`
	Amount int `json:"amount"`
}

var rootCmd = &cobra.Command {
	Use: "expense-tracker",
	Short: "Control your expenses",
	Long: "Add, Edit, Summarize your expenses",
}

var addExpenseCmd = &cobra.Command{
	Use:   "add",
	Short: "Add expense",
	Long:  "Add expense to the file",
	Run:   addExpense,
}

func init() {
	addExpenseCmd.Flags().StringVar(&description, "description", "", "Description of the task")
	addExpenseCmd.Flags().IntVar(&amount, "amount", 0, "Amount of the expense")
	rootCmd.AddCommand(addExpenseCmd)
}

func addExpense (cmd *cobra.Command, args []string) {
	fmt.Println(description, amount)
}

func main () {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
