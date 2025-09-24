package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/caiosemblano/expense-tracker/internal/models"
)

type BudgetStorage struct {
	filepath string
	mu       sync.RWMutex
}

func NewBudgetStorage() *BudgetStorage {
	dataDir := "data"
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		panic(err)
	}

	return &BudgetStorage{
		filepath: filepath.Join(dataDir, "budgets.json"),
	}
}

func (s *BudgetStorage) SaveBudget(budget *models.Budget) error {
	budgets, err := s.LoadBudgets()
	if err != nil {
		return err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	// Atualizar ou adicionar orçamento
	found := false
	for i, b := range budgets {
		if b.Month == budget.Month && b.Year == budget.Year {
			budgets[i] = budget
			found = true
			break
		}
	}

	if !found {
		budgets = append(budgets, budget)
	}

	return s.saveToFile(budgets)
}

func (s *BudgetStorage) LoadBudgets() ([]*models.Budget, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if _, err := os.Stat(s.filepath); os.IsNotExist(err) {
		return []*models.Budget{}, nil
	}

	file, err := os.ReadFile(s.filepath)
	if err != nil {
		return nil, err
	}

	var budgets []*models.Budget
	if err := json.Unmarshal(file, &budgets); err != nil {
		return nil, err
	}

	return budgets, nil
}

func (s *BudgetStorage) GetBudget(month time.Month, year int) (*models.Budget, error) {
	budgets, err := s.LoadBudgets()
	if err != nil {
		return nil, err
	}

	for _, b := range budgets {
		if b.Month == month && b.Year == year {
			return b, nil
		}
	}

	return nil, fmt.Errorf("nenhum orçamento encontrado para %s/%d", month, year)
}

func (s *BudgetStorage) saveToFile(budgets []*models.Budget) error {
	data, err := json.MarshalIndent(budgets, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.filepath, data, 0644)
}
