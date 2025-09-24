package cmd

import (
	"fmt"

	"github.com/caiosemblano/expense-tracker/internal/storage"
	"github.com/spf13/cobra"
)

var expenseID int

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove uma despesa pelo ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		if expenseID <= 0 {
			return fmt.Errorf("ID inválido: deve ser maior que zero")
		}

		store := storage.NewStorage()
		if err := store.DeleteExpense(expenseID); err != nil {
			return fmt.Errorf("erro ao deletar despesa: %v", err)
		}

		fmt.Println("Expense deleted successfully")
		return nil
	},
}

func init() {
	deleteCmd.Flags().IntVarP(&expenseID, "id", "i", 0, "ID da despesa a ser removida")
	deleteCmd.MarkFlagRequired("id")
	rootCmd.AddCommand(deleteCmd)
}
