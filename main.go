package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/olekukonko/tablewriter"
)

var (
	description string
	amount int
	id int
)

type ExpenseTracker struct {
	ID int `json:"id"`
	Date string `json:"date"`
	Description string `json:"description"`
	Amount int `json:"amount"`
}

func loadExpenses () ([]ExpenseTracker, error) {
	jsonData, err := os.ReadFile("expenses.json")
	if err != nil {
		log.Fatalf("Error while reading file %v", err)
	}

	if len(jsonData) == 0 {
		return []ExpenseTracker{}, nil
	}

	var expenses []ExpenseTracker

	err = json.Unmarshal(jsonData, &expenses)

	if err != nil {
		log.Fatalf("Error while unmarshalling %v", err)
	}

	return expenses, err
}

func encodeExpenses (e []ExpenseTracker) []byte {
	jExpense, err := json.Marshal(e)
	if err != nil {
		log.Fatalf("Error while turning code to byte %v", err)
	}
	return jExpense
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

var listExpenseCmd = &cobra.Command{
	Use:   "list",
	Short: "List expenses",
	Long:  "List expenses to the see all expenses",
	Run:   showExpenses,
}

var summaryExpensesCmd = &cobra.Command{
	Use:   "summary",
	Short: "Summary of expenses",
	Long:  "Summary of expense totals of amount you spent",
	Run:   showSummary,
}

var deleteExpenseCmd = &cobra.Command{
	Use: "delete",
	Short: "Delete expense",
	Long: "Delete some expenses you do not need",
	Run: deleteExpense,
}

func init() {
	addExpenseCmd.Flags().StringVar(&description, "description", "", "Description of the task")
	addExpenseCmd.Flags().IntVar(&amount, "amount", 0, "Amount of the expense")
	deleteExpenseCmd.Flags().IntVar(&id, "id", 0, "Id to delete expense")
	rootCmd.AddCommand(addExpenseCmd, listExpenseCmd, deleteExpenseCmd, summaryExpensesCmd)
}

func addExpense (cmd *cobra.Command, args []string) {
	existingExpense, err := loadExpenses()
	nextID := 1

	for _, t := range existingExpense {
		if t.ID >= nextID {
			nextID = t.ID + 1
		}
	}

	now := time.Now()

	newExpense := ExpenseTracker {
		ID: nextID,
		Date: now.Format("2006-01-2"),
		Description: description,
		Amount: amount,
	}

	existingExpense = append(existingExpense, newExpense)

	byteData := encodeExpenses(existingExpense)

	err = os.WriteFile("expenses.json", byteData, 0644)

	if err != nil {
		log.Fatalf("Error while writing expenses, %v", err)
	}

	fmt.Printf("Expense added successfully (ID: %d)", newExpense.ID)
}

func showExpenses (cmd *cobra.Command, args []string) {
	existingExpenses, err := loadExpenses()

	if err != nil {
		log.Fatalf("Error while reading existing expenses %v", err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.Header("ID", "Date", "Description", "Amount")
	table.Bulk(existingExpenses)
	table.Render()
}

func showSummary (cmd * cobra.Command, args []string) {
  existingExpenses, err := loadExpenses()
	if err != nil {
		log.Fatalf("Error while extracting existing expenses, %v", err)
	}

	var total = 0

	for _, e := range existingExpenses {
		total = total + e.Amount
	}
	fmt.Println("Total expenses:", total)
}

func deleteExpense (cmd * cobra.Command, args []string) {
	existingExpenses, err := loadExpenses()
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
	byteData := encodeExpenses(existingExpenses)

	err = os.WriteFile("expenses.json", byteData, 0644)

	if err != nil {
		log.Fatalf("Error while writing updated data to the file %v", err)
	}
	fmt.Println("Expense deleted successfully")
}

func main () {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
