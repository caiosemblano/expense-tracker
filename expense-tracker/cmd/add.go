package cmd

import (
	"fmt"

	"github.com/caiosemblano/expense-tracker/internal/models"
	"github.com/caiosemblano/expense-tracker/internal/storage"
	"github.com/spf13/cobra"
)

var description string
var amount float64
var category string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adiciona uma nova despesa",
	RunE: func(cmd *cobra.Command, args []string) error {
		if amount <= 0 {
			return fmt.Errorf("o valor deve ser maior que zero")
		}

		expense := models.NewExpense(description, amount, category)
		store := storage.NewStorage()

		if err := store.SaveExpense(expense); err != nil {
			return fmt.Errorf("erro ao salvar despesa: %v", err)
		}

		fmt.Printf("Expense added successfully (ID: %d)\n", expense.ID)
		return nil
	},
}

func init() {
	addCmd.Flags().StringVarP(&description, "description", "d", "", "Descrição da despesa")
	addCmd.Flags().Float64VarP(&amount, "amount", "a", 0, "Valor da despesa")
	addCmd.Flags().StringVarP(&category, "category", "c", "Others", "Categoria da despesa")

	addCmd.MarkFlagRequired("description")
	addCmd.MarkFlagRequired("amount")

	rootCmd.AddCommand(addCmd)
}
