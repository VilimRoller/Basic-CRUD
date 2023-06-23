package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type ExpenseType int

const (
	Food           ExpenseType = 1
	Cleaning       ExpenseType = 2
	Utilities      ExpenseType = 3
	Transportation ExpenseType = 4
	Entertainment  ExpenseType = 5
	Health         ExpenseType = 6
	Recreation     ExpenseType = 7
	Other          ExpenseType = 8
)

var EmptyExpense Expense = Expense{
	Name:     "",
	Date:     "",
	Type:     Other,
	Amount:   0,
	Currency: "",
}

type Expense struct {
	Name     string
	Date     string
	Type     ExpenseType
	Amount   float32
	Currency string
}

func (expense Expense) ToString() (string, error) {
	expenseType, err := GetExpenseTypeAsString(expense.Type)

	if err != nil {
		return "", errors.New("ToString: parsing type failed.")
	}

	return fmt.Sprintf("Name: %s\nDate: %s\nType: %s\nAmount: %.2f\nCurrency: %s\n", expense.Name, expense.Date, expenseType, expense.Amount, expense.Currency), nil
}

func GetExpenseFromString(inputStr string) (Expense, error) {
	elements := strings.Split(inputStr, "\n")
	returnVal := Expense{}

	for _, element := range elements {
		elementSubstring := strings.SplitN(element, ":", 2)
		if len(elementSubstring) != 2 {
			return returnVal, errors.New("GetExpenseFromString: parsing string failed.")
		}

		key := strings.TrimSpace(elementSubstring[0])
		value := strings.TrimSpace(elementSubstring[1])
		err := addKeyAndValueToExpense(key, value, &returnVal)
		if err != nil {
			return returnVal, errors.New("GetExpenseFromString: adding key and value to expense failed.")
		}
	}

	return returnVal, nil
}

func addKeyAndValueToExpense(key, value string, expense *Expense) error {
	switch key {
	case "Name":
		expense.Name = value
	case "Date":
		expense.Date = value
	case "Type":
		expenseType, err := GetStringAsExpenseType(value)
		if err != nil {
			return errors.New("addKeyAndValueToExpense: parsing type failed.")
		}
		expense.Type = expenseType
	case "Amount":
		amount, err := strconv.ParseFloat(value, 32)
		if err == nil {
			expense.Amount = float32(amount)
		}
	case "Currency":
		expense.Currency = value
	default:
		errors.New("addKeyAndValueToExpense: unknown key.")
	}

	return nil
}

func GetStringAsExpenseType(expenseType string) (ExpenseType, error) {
	switch expenseType {
	case "Food":
		return Food, nil
	case "Cleaning":
		return Cleaning, nil
	case "Utilities":
		return Utilities, nil
	case "Transportation":
		return Transportation, nil
	case "Entertainment":
		return Entertainment, nil
	case "Health":
		return Health, nil
	case "Recreation":
		return Recreation, nil
	case "Other":
		return Other, nil
	default:
		return Other, errors.New("GetStringAsExpenseType: unknown expense type")
	}
}

func GetExpenseTypeAsString(expenseType ExpenseType) (string, error) {
	switch expenseType {
	case Food:
		return "Food", nil
	case Cleaning:
		return "Cleaning", nil
	case Utilities:
		return "Utilities", nil
	case Transportation:
		return "Transportation", nil
	case Entertainment:
		return "Entertainment", nil
	case Health:
		return "Health", nil
	case Recreation:
		return "Recreation", nil
	case Other:
		return "Other", nil
	default:
		return "", errors.New("GetExpenseTypeAsString: unknown expense type")
	}
}
