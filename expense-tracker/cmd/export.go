package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/caiosemblano/expense-tracker/internal/storage"
	"github.com/spf13/cobra"
)

var outputFile string

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Exporta despesas para CSV",
	RunE: func(cmd *cobra.Command, args []string) error {
		store := storage.NewStorage()
		expenses, err := store.LoadExpenses()
		if err != nil {
			return fmt.Errorf("erro ao carregar despesas: %v", err)
		}

		if len(expenses) == 0 {
			return fmt.Errorf("nenhuma despesa encontrada para exportar")
		}

		file, err := os.Create(outputFile)
		if err != nil {
			return fmt.Errorf("erro ao criar arquivo CSV: %v", err)
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		// Escrever cabeçalho
		header := []string{"ID", "Date", "Category", "Description", "Amount"}
		if err := writer.Write(header); err != nil {
			return fmt.Errorf("erro ao escrever cabeçalho: %v", err)
		}

		// Escrever despesas
		for _, e := range expenses {
			record := []string{
				strconv.Itoa(e.ID),
				e.Date.Format("2006-01-02"),
				e.Category,
				e.Description,
				fmt.Sprintf("%.2f", e.Amount),
			}

			if err := writer.Write(record); err != nil {
				return fmt.Errorf("erro ao escrever despesa: %v", err)
			}
		}

		fmt.Printf("Expenses exported successfully to %s\n", outputFile)
		return nil
	},
}

func init() {
	exportCmd.Flags().StringVarP(&outputFile, "output", "o", "expenses.csv", "Nome do arquivo CSV de saída")
	rootCmd.AddCommand(exportCmd)
}
