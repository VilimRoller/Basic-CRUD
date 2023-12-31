package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type ExpenseType int

const (
	BadType        ExpenseType = 0
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
	Type:     BadType,
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
		return "", errors.New("ToString: parsing type failed\n")
	}

	return fmt.Sprintf("Name: %s\nDate: %s\nType: %s\nAmount: %.2f\nCurrency: %s\n", expense.Name, expense.Date, expenseType, expense.Amount, expense.Currency), nil
}

func GetExpenseFromString(inputStr string) (Expense, error) {
	elements := strings.Split(inputStr, "\n")
	returnVal := Expense{}

	for _, element := range elements {
		elementSubstring := strings.SplitN(element, ":", 2)
		if len(elementSubstring) != 2 {
			return returnVal, errors.New("GetExpenseFromString: parsing string failed.\n")
		}

		key := strings.TrimSpace(elementSubstring[0])
		value := strings.TrimSpace(elementSubstring[1])
		err := AddKeyAndValueToExpense(key, value, &returnVal)
		if err != nil {
			return returnVal, errors.New("GetExpenseFromString: adding key and value to expense failed.\n")
		}
	}

	err := CheckIfAllFieldsAreFilled(&returnVal)

	if err != nil {
		return returnVal, err
	}

	return returnVal, nil
}

func CheckIfAllFieldsAreFilled(expense *Expense) error {
	if expense.Name == "" {
		return errors.New("checkIfAllFieldsAreFilled: 'Name' field is required.\n")
	}
	if expense.Date == "" {
		return errors.New("checkIfAllFieldsAreFilled: 'Date' field is required.\n")
	}
	if expense.Type == BadType {
		return errors.New("checkIfAllFieldsAreFilled: 'Type' field is required.\n")
	}
	if expense.Amount == 0.0 {
		return errors.New("checkIfAllFieldsAreFilled: 'Amount' field is required.\n")
	}
	if expense.Currency == "" {
		return errors.New("checkIfAllFieldsAreFilled: 'Currency' field is required.\n")
	}

	return nil
}

func AddKeyAndValueToExpense(key, value string, expense *Expense) error {
	switch key {
	case "Name":
		expense.Name = value
	case "Date":
		expense.Date = value
	case "Type":
		expenseType, err := GetStringAsExpenseType(value)
		if err != nil {
			return errors.New("addKeyAndValueToExpense: parsing type failed.\n")
		}
		expense.Type = expenseType
	case "Amount":
		amount, err := strconv.ParseFloat(value, 32)
		if err == nil {
			expense.Amount = float32(amount)
		}
	case "Currency":
		if len(value) != 3 {
			return errors.New("addKeyAndValueToExpense: currency string should follow ISO 4217 standard (3 letters only).\n")
		}
		expense.Currency = value
	default:
		return errors.New("addKeyAndValueToExpense: unknown key.\n")
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
		return BadType, errors.New("GetStringAsExpenseType: unknown expense type\n")
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
		return "", errors.New("GetExpenseTypeAsString: unknown expense type\n")
	}
}
