package cmd

import (
	"fmt"

	"github.com/caiosemblano/expense-tracker/internal/storage"
	"github.com/spf13/cobra"
)

var updateID int
var newDescription string
var newAmount float64

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Atualiza uma despesa existente",
	RunE: func(cmd *cobra.Command, args []string) error {
		if updateID <= 0 {
			return fmt.Errorf("ID inválido: deve ser maior que zero")
		}

		if newAmount <= 0 {
			return fmt.Errorf("valor deve ser maior que zero")
		}

		store := storage.NewStorage()
		if err := store.UpdateExpense(updateID, newDescription, newAmount); err != nil {
			return fmt.Errorf("erro ao atualizar despesa: %v", err)
		}

		fmt.Printf("Despesa com ID %d atualizada com sucesso!\n", updateID)
		return nil
	},
}

func init() {
	updateCmd.Flags().IntVarP(&updateID, "id", "i", 0, "ID da despesa a ser atualizada")
	updateCmd.Flags().StringVarP(&newDescription, "description", "d", "", "Nova descrição")
	updateCmd.Flags().Float64VarP(&newAmount, "amount", "a", 0, "Novo valor")

	updateCmd.MarkFlagRequired("id")
	updateCmd.MarkFlagRequired("description")
	updateCmd.MarkFlagRequired("amount")

	rootCmd.AddCommand(updateCmd)
}
