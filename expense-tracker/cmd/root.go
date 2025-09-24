package cmd

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "expense-tracker",
	Short: "Expense Tracker is a CLI application to manage your expenses.",
	Long:  `Expense Tracker is a CLI application to manage your expenses. You can add, list, and delete expenses.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Olá! Você chamou a CLI sem subcomandos.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}