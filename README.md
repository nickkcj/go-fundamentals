# Go Fundamentals - Learning Repository

This repository contains Go fundamentals with practical examples and exercises.

## 📁 Structure

```
Go - Complete Guide/
├── go.work                     # Go workspace file
├── README.md                   # This file
├── basics/                     # Basic Go concepts
│   ├── variables/             # Variable declarations & types
│   │   ├── go.mod
│   │   └── main.go
│   └── strings/               # String manipulation
│       ├── go.mod
│       └── main.go
├── exercises/                  # Practical projects
│   ├── investment-calculator/  # Investment calculator project
│   │   ├── go.mod
│   │   └── main.go
│   └── profit-calculator/      # Profit calculator project
│       ├── go.mod
│       └── main.go
└── shared/                     # Shared utilities
    └── utils/                  # Common helper functions
        ├── go.mod
        └── input.go
```

## 🚀 How to Use

### Running Individual Examples

Each folder has its own `go.mod` file, so you can run them independently:

```bash
# Run variables example
cd basics/variables
go run main.go

# Run investment calculator
cd exercises/investment-calculator  
go run main.go
```

### Using the Workspace

The `go.work` file allows sharing code between modules:

```bash
# From the root directory
go work sync  # Sync all modules
```

## 📚 Topics Covered

### Basics
- **Variables**: Different ways to declare variables in Go
- **Strings**: String manipulation and formatting

### Exercises  
- **Investment Calculator**: Learn variables, input/output, math operations
- **Profit Calculator**: Practice with calculations and formatting

### Shared Utils
- **Input Functions**: Reusable input handling functions

## 🎯 Learning Path

1. Start with `basics/variables` - Learn variable declarations
2. Move to `basics/strings` - String operations  
3. Try `exercises/investment-calculator` - Your first project
4. Build `exercises/profit-calculator` - Practice calculations
5. Explore `shared/utils` - Learn about packages and imports

## 🔧 Requirements

- Go 1.24.5 
- Basic understanding of programming concepts

## 📖 Additional Resources

- [Official Go Documentation](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://golang.org/doc/effective_go.html)
