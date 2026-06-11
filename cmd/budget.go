package cmd

import (
	t "expense-tracker/internal/expense/types"
	s "expense-tracker/internal/storage"
	u "expense-tracker/internal/utils"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var setBudgetCmd = &cobra.Command{
	Use: "budget",
	Short: "Set your budget",
	Long: "Set your monthly spending budget",
	RunE: setBudget,
}

func setBudget (cmd *cobra.Command, args []string) error {
	if monthly < 0 {
		return fmt.Errorf("monthly budget cannot be less than 0")
	}

	existingBudget, err := s.LoadBudget() 
	if err != nil {
		return err
	}

	newBudget := t.Budget {
		Monthly: monthly,
	}

	existingBudget = newBudget

	byteData, err := u.EncodeBudget(existingBudget)
	if err != nil {
		return fmt.Errorf("writing expenses: %w", err)
	}

	err = os.WriteFile("budget.json", byteData, permissionCode)
	if err != nil {
		return fmt.Errorf("writing expenses: %w", err)
	}

	fmt.Printf("Budget set successfully!")

	return nil
}
