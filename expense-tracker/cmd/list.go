package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/caiosemblano/expense-tracker/internal/models"
	"github.com/caiosemblano/expense-tracker/internal/storage"
	"github.com/spf13/cobra"
)

var categoryFilter string

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lista todas as despesas",
	RunE: func(cmd *cobra.Command, args []string) error {
		store := storage.NewStorage()
		expenses, err := store.LoadExpenses()
		if err != nil {
			return fmt.Errorf("erro ao carregar despesas: %v", err)
		}

		if len(expenses) == 0 {
			fmt.Println("Nenhuma despesa encontrada")
			return nil
		}

		// Filtrar por categoria se especificado
		if categoryFilter != "" {
			var filtered []*models.Expense
			for _, e := range expenses {
				if e.Category == categoryFilter {
					filtered = append(filtered, e)
				}
			}
			expenses = filtered
		}

		if len(expenses) == 0 {
			fmt.Printf("Nenhuma despesa encontrada na categoria: %s\n", categoryFilter)
			return nil
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "ID\tData\tCategoria\tDescrição\tValor")
		fmt.Fprintln(w, "--\t----\t---------\t---------\t-----")

		for _, e := range expenses {
			fmt.Fprintf(w, "%d\t%s\t%s\t%s\t$%.2f\n",
				e.ID,
				e.Date.Format("2006-01-02"),
				e.Category,
				e.Description,
				e.Amount)
		}

		return w.Flush()
	},
}

func init() {
	listCmd.Flags().StringVarP(&categoryFilter, "category", "c", "", "Filtrar por categoria")
	rootCmd.AddCommand(listCmd)
}
