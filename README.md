# Expense Tracker CLI

A command-line interface (CLI) application for personal expense management, built with Go.

## 🚀 Features

- ✅ Complete expense management (add, update, delete)
- 📊 Expense categorization
- 💰 Monthly budget system
- 📈 Summaries and reports
- 📁 CSV export

## 📋 Prerequisites

- Go 1.25 or higher
- Operating System: Windows, Linux, or macOS

## ⚙️ Installation

1. Clone the repository:
```bash
git clone https://github.com/caiosemblano/expense-tracker.git
cd expense-tracker
```

2. Install dependencies:
```bash
go mod download
```

3. Build the project:
```bash
go build -o expense-tracker.exe
```

## 🎯 How to Use

### Add an expense
```bash
expense-tracker add --description "Lunch" --amount 25.50 --category "Food"
```

### List all expenses
```bash
expense-tracker list
```

### List expenses by category
```bash
expense-tracker list --category "Food"
```

### Update an expense
```bash
expense-tracker update --id 1 --description "New lunch" --amount 30.00
```

### Delete an expense
```bash
expense-tracker delete --id 1
```

### View summary
```bash
# General summary
expense-tracker summary

# Monthly summary
expense-tracker summary --month 9
```

### Set monthly budget
```bash
expense-tracker budget --month 9 --amount 1000.00
```

### Export to CSV
```bash
expense-tracker export --output "my_expenses.csv"
```

## 📁 Project Structure

```
expense-tracker/
├── cmd/              # CLI Commands
├── internal/         # Internal packages
│   ├── models/      # Data models
│   └── storage/     # Data persistence
├── data/            # Data storage
└── main.go          # Entry point
```

## 🛠️ Technologies Used

- [Go](https://golang.org/) - Programming language
- [Cobra](https://github.com/spf13/cobra) - CLI Framework
- JSON - Data storage format

## 📌 Features

- [x] Complete expense CRUD
- [x] Categorization
- [x] Budget system
- [x] Monthly summaries
- [x] CSV export
- [ ] Data import
- [ ] Terminal charts
- [ ] Multi-currency

## 📝 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## ✨ Author

Made by Caio Semblano
