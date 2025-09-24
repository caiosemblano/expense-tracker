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

type Storage struct {
	filepath string
	mu       sync.RWMutex
}

func NewStorage() *Storage {
	// Cria o diretório data se não existir
	dataDir := "data"
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		panic(err)
	}

	return &Storage{
		filepath: filepath.Join(dataDir, "expenses.json"),
	}
}

func (s *Storage) SaveExpense(expense *models.Expense) error {
	expenses, err := s.LoadExpenses()
	if err != nil {
		return err
	}

	s.mu.Lock()
	defer s.mu.Unlock() // Define novo ID
	maxID := 0
	for _, e := range expenses {
		if e.ID > maxID {
			maxID = e.ID
		}
	}
	expense.ID = maxID + 1

	expenses = append(expenses, expense)
	return s.saveToFile(expenses)
}

func (s *Storage) LoadExpenses() ([]*models.Expense, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if _, err := os.Stat(s.filepath); os.IsNotExist(err) {
		return []*models.Expense{}, nil
	}

	file, err := os.ReadFile(s.filepath)
	if err != nil {
		return nil, err
	}

	var expenses []*models.Expense
	if err := json.Unmarshal(file, &expenses); err != nil {
		return nil, err
	}

	return expenses, nil
}

func (s *Storage) DeleteExpense(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	expenses, err := s.LoadExpenses()
	if err != nil {
		return err
	}

	index := -1
	for i, e := range expenses {
		if e.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("despesa com ID %d não encontrada", id)
	}

	// Remove o elemento do slice
	expenses = append(expenses[:index], expenses[index+1:]...)
	return s.saveToFile(expenses)
}

func (s *Storage) UpdateExpense(id int, description string, amount float64) error {
	expenses, err := s.LoadExpenses()
	if err != nil {
		return err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	found := false
	for i, e := range expenses {
		if e.ID == id {
			expenses[i].Description = description
			expenses[i].Amount = amount
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("despesa com ID %d não encontrada", id)
	}

	return s.saveToFile(expenses)
}

func (s *Storage) GetExpensesByMonth(year int, month time.Month) ([]*models.Expense, error) {
	expenses, err := s.LoadExpenses()
	if err != nil {
		return nil, err
	}

	var filtered []*models.Expense
	for _, e := range expenses {
		if e.Date.Year() == year && e.Date.Month() == month {
			filtered = append(filtered, e)
		}
	}

	return filtered, nil
}

func (s *Storage) saveToFile(expenses []*models.Expense) error {
	data, err := json.MarshalIndent(expenses, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.filepath, data, 0644)
}
