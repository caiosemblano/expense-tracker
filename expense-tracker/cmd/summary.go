package cmd

import (
	"fmt"
	"time"

	"github.com/caiosemblano/expense-tracker/internal/storage"
	"github.com/spf13/cobra"
)

var month int

var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Mostra o resumo das despesas",
	RunE: func(cmd *cobra.Command, args []string) error {
		store := storage.NewStorage()

		if month > 0 {
			if month < 1 || month > 12 {
				return fmt.Errorf("mês inválido: deve estar entre 1 e 12")
			}

			expenses, err := store.GetExpensesByMonth(time.Now().Year(), time.Month(month))
			if err != nil {
				return fmt.Errorf("erro ao carregar despesas: %v", err)
			}

			var total float64
			for _, e := range expenses {
				total += e.Amount
			}

			// Verificar orçamento
			budgetStore := storage.NewBudgetStorage()
			budget, _ := budgetStore.GetBudget(time.Month(month), time.Now().Year())

			fmt.Printf("Total expenses for %s: $%.2f\n",
				time.Month(month).String(),
				total)

			if budget != nil {
				fmt.Printf("Monthly budget: $%.2f (%.1f%% used)\n",
					budget.Amount,
					(total/budget.Amount)*100)

				if total > budget.Amount {
					fmt.Printf("WARNING: Over budget by $%.2f\n", total-budget.Amount)
				}
			}
			return nil
		}

		// Resumo total
		expenses, err := store.LoadExpenses()
		if err != nil {
			return fmt.Errorf("erro ao carregar despesas: %v", err)
		}

		var total float64
		for _, e := range expenses {
			total += e.Amount
		}

		fmt.Printf("Total expenses: $%.2f\n", total)
		return nil
	},
}

func init() {
	summaryCmd.Flags().IntVarP(&month, "month", "m", 0, "Mês para filtrar (1-12)")
	rootCmd.AddCommand(summaryCmd)
}
