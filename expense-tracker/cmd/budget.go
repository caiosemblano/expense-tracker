package cmd

import (
	"fmt"
	"time"

	"github.com/caiosemblano/expense-tracker/internal/models"
	"github.com/caiosemblano/expense-tracker/internal/storage"
	"github.com/spf13/cobra"
)

var budgetAmount float64
var budgetMonth int

var budgetCmd = &cobra.Command{
	Use:   "budget",
	Short: "Gerencia orçamentos mensais",
	RunE: func(cmd *cobra.Command, args []string) error {
		if budgetAmount <= 0 {
			return fmt.Errorf("o valor do orçamento deve ser maior que zero")
		}

		if budgetMonth < 1 || budgetMonth > 12 {
			return fmt.Errorf("mês inválido: deve estar entre 1 e 12")
		}

		// Criar orçamento
		budget := models.NewBudget(time.Month(budgetMonth), time.Now().Year(), budgetAmount)
		store := storage.NewBudgetStorage()

		if err := store.SaveBudget(budget); err != nil {
			return fmt.Errorf("erro ao salvar orçamento: %v", err)
		}

		fmt.Printf("Budget set for %s: $%.2f\n", time.Month(budgetMonth), budgetAmount)

		// Verificar gastos do mês
		expenseStore := storage.NewStorage()
		expenses, err := expenseStore.GetExpensesByMonth(time.Now().Year(), time.Month(budgetMonth))
		if err != nil {
			return fmt.Errorf("erro ao carregar despesas: %v", err)
		}

		var total float64
		for _, e := range expenses {
			total += e.Amount
		}

		if total > budgetAmount {
			fmt.Printf("WARNING: Current expenses ($%.2f) exceed budget by $%.2f\n",
				total, total-budgetAmount)
		} else {
			fmt.Printf("Current expenses: $%.2f (%.1f%% of budget)\n",
				total, (total/budgetAmount)*100)
		}

		return nil
	},
}

func init() {
	budgetCmd.Flags().Float64VarP(&budgetAmount, "amount", "a", 0, "Valor do orçamento mensal")
	budgetCmd.Flags().IntVarP(&budgetMonth, "month", "m", int(time.Now().Month()), "Mês (1-12)")

	budgetCmd.MarkFlagRequired("amount")

	rootCmd.AddCommand(budgetCmd)
}
