# Expense Tracker CLI

Uma aplicação de linha de comando (CLI) para gerenciamento de despesas pessoais, desenvolvida em Go.

## 🚀 Funcionalidades

- ✅ Gerenciamento completo de despesas (adicionar, atualizar, remover)
- 📊 Categorização de despesas
- 💰 Sistema de orçamento mensal
- 📈 Resumos e relatórios
- 📁 Exportação para CSV

## 📋 Pré-requisitos

- Go 1.25 ou superior
- Sistema operacional: Windows, Linux ou macOS

## ⚙️ Instalação

1. Clone o repositório:
```bash
git clone https://github.com/caiosemblano/expense-tracker.git
cd expense-tracker
```

2. Instale as dependências:
```bash
go mod download
```

3. Compile o projeto:
```bash
go build -o expense-tracker.exe
```

## 🎯 Como Usar

### Adicionar uma despesa
```bash
expense-tracker add --description "Almoço" --amount 25.50 --category "Food"
```

### Listar todas as despesas
```bash
expense-tracker list
```

### Listar despesas por categoria
```bash
expense-tracker list --category "Food"
```

### Atualizar uma despesa
```bash
expense-tracker update --id 1 --description "Novo almoço" --amount 30.00
```

### Remover uma despesa
```bash
expense-tracker delete --id 1
```

### Ver resumo
```bash
# Resumo geral
expense-tracker summary

# Resumo mensal
expense-tracker summary --month 9
```

### Configurar orçamento mensal
```bash
expense-tracker budget --month 9 --amount 1000.00
```

### Exportar para CSV
```bash
expense-tracker export --output "minhas_despesas.csv"
```

## 📁 Estrutura do Projeto

```
expense-tracker/
├── cmd/              # Comandos da CLI
├── internal/         # Pacotes internos
│   ├── models/      # Definições de dados
│   └── storage/     # Persistência
├── data/            # Armazenamento de dados
└── main.go          # Ponto de entrada
```

## 🛠️ Tecnologias Utilizadas

- [Go](https://golang.org/) - Linguagem de programação
- [Cobra](https://github.com/spf13/cobra) - Framework CLI
- JSON - Formato de armazenamento de dados

## 📌 Features

- [x] CRUD completo de despesas
- [x] Categorização
- [x] Sistema de orçamento
- [x] Resumos mensais
- [x] Exportação CSV
- [ ] Importação de dados
- [ ] Gráficos no terminal
- [ ] Multi-moeda

## 🤝 Contribuindo

1. Faça um Fork do projeto
2. Crie uma Branch para sua Feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a Branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📝 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## ✨ Autor

Feito por Caio Semblano
