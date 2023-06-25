package data

import (
	"errors"
	"testing"
)

var testExpense Expense = Expense{
	Name:     "Testing tools",
	Date:     "24-06-2023",
	Type:     Cleaning,
	Amount:   120,
	Currency: "EUR",
}

func TestGetExpenseTypeAsString(tst *testing.T) {
	testCases := []struct {
		expenseType    ExpenseType
		expectedOutput string
		expectedError  error
	}{
		{Food, "Food", nil},
		{Cleaning, "Cleaning", nil},
		{Utilities, "Utilities", nil},
		{Transportation, "Transportation", nil},
		{Entertainment, "Entertainment", nil},
		{Health, "Health", nil},
		{Recreation, "Recreation", nil},
		{Other, "Other", nil},
		{BadType, "", errors.New("GetExpenseTypeAsString: unknown expense type\n")},
	}

	for _, testCase := range testCases {
		output, err := GetExpenseTypeAsString(testCase.expenseType)

		// Check if the output matches the expected value
		if output != testCase.expectedOutput {
			tst.Errorf("Expected: %s, Got: %s\n", testCase.expectedOutput, output)
		}

		// Check if the error matches the expected error
		if (err == nil && testCase.expectedError != nil) ||
			(err != nil && testCase.expectedError == nil) ||
			(err != nil && testCase.expectedError != nil && err.Error() != testCase.expectedError.Error()) {
			tst.Errorf("Expected error: %v, Got error: %v\n", testCase.expectedError, err)
		}
	}
}

func TestGetStringAsExpenseType(tst *testing.T) {
	testCases := []struct {
		input          string
		expectedOutput ExpenseType
		expectedError  error
	}{
		{"Food", Food, nil},
		{"Cleaning", Cleaning, nil},
		{"Utilities", Utilities, nil},
		{"Transportation", Transportation, nil},
		{"Entertainment", Entertainment, nil},
		{"Health", Health, nil},
		{"Recreation", Recreation, nil},
		{"Other", Other, nil},
		{"Unknown", BadType, errors.New("GetStringAsExpenseType: unknown expense type\n")},
	}

	for _, testCase := range testCases {
		output, err := GetStringAsExpenseType(testCase.input)

		// Check if the output matches the expected value
		if output != testCase.expectedOutput {
			tst.Errorf("Expected: %d, Got: %d\n", testCase.expectedOutput, output)
		}

		// Check if the error matches the expected error
		if (err == nil && testCase.expectedError != nil) ||
			(err != nil && testCase.expectedError == nil) ||
			(err != nil && testCase.expectedError != nil && err.Error() != testCase.expectedError.Error()) {
			tst.Errorf("Expected error: %v, Got error: %v\n", testCase.expectedError, err)
		}
	}
}

func TestAddKeyAndValueToExpenseUpdateName(tst *testing.T) {
	expense := testExpense
	err := addKeyAndValueToExpense("Name", "Testing equipment", &expense)

	if err != nil {
		tst.Errorf("Adding key and value to expense failed: %v\n", err)
	}
	if expense.Name != "Testing equipment" {
		tst.Errorf("Name was not updated correctly\n")
	}
}

func TestAddKeyAndValueToExpenseUpdateDate(tst *testing.T) {
	expense := testExpense
	err := addKeyAndValueToExpense("Date", "01-01-1970", &expense)

	if err != nil {
		tst.Errorf("Adding key and value to expense failed: %v\n", err)
	}
	if expense.Date != "01-01-1970" {
		tst.Errorf("Date was not updated correctly\n")
	}
}

func TestAddKeyAndValueToExpenseUpdateExpenseType(tst *testing.T) {
	expense := testExpense
	err := addKeyAndValueToExpense("Type", "Entertainment", &expense)

	if err != nil {
		tst.Errorf("Adding key and value to expense failed: %v\n", err)
	}
	if expense.Type != Entertainment {
		tst.Errorf("Expense type was not updated correctly\n")
	}
}

func TestAddKeyAndValueToExpenseUpdateAmount(tst *testing.T) {
	expense := testExpense
	err := addKeyAndValueToExpense("Amount", "150", &expense)

	if err != nil {
		tst.Errorf("Adding key and value to expense failed: %v", err)
	}
	if expense.Amount != float32(150) {
		tst.Errorf("Amount was not updated correctly. Amount = %.2f expected amount = 150\n", expense.Amount)
	}
}

func TestAddKeyAndValueToExpenseUpdateAmountWithEmptyInput(tst *testing.T) {
	expense := testExpense
	originalAmount := expense.Amount
	err := addKeyAndValueToExpense("Amount", "", &expense)

	if err != nil {
		tst.Errorf("Adding key and value to expense failed: %v\n", err)
	}
	if expense.Amount != originalAmount {
		tst.Errorf("Amount changed. Amount = %.2f expected amount = %.2f\n", expense.Amount, originalAmount)
	}
}

func TestAddKeyAndValueToExpenseUpdateCurrecny(tst *testing.T) {
	expense := testExpense
	err := addKeyAndValueToExpense("Currency", "USD", &expense)

	if err != nil {
		tst.Errorf("Adding key and value to expense failed: %v\n", err)
	}
	if expense.Currency != "USD" {
		tst.Errorf("Amount was not updated correctly\n")
	}
}

func TestAddKeyAndValueToExpenseUpdateCurrecnyWithBadInput(tst *testing.T) {
	expense := testExpense
	err := addKeyAndValueToExpense("Currency", "US", &expense)

	if err == nil {
		tst.Errorf("Error was expected for bad currency input\n")
	}
}

func TestAddKeyAndValueToExpenseUnknownKey(tst *testing.T) {
	expense := testExpense
	err := addKeyAndValueToExpense("Bip bop", "Bing bong", &expense)

	if err == nil {
		tst.Errorf("Error was expected for unknown key\n")
	}
}

func TestAddKeyAndValueToExpenseEmptyKey(tst *testing.T) {
	expense := testExpense
	err := addKeyAndValueToExpense("", "Something", &expense)

	if err == nil {
		tst.Errorf("Error was expected for unknown key\n")
	}
}

func TestGetExpenseFromValidString(tst *testing.T) {
	inputString := "Name: Running shoes\nDate: 01-05-2023\nType: Recreation\nAmount: 25.5\nCurrency: EUR"
	expectedExpense := Expense{
		Name:     "Running shoes",
		Date:     "01-05-2023",
		Type:     Recreation,
		Amount:   25.5,
		Currency: "EUR",
	}

	expense, err := GetExpenseFromString(inputString)
	if err != nil {
		tst.Errorf("Test case 1 failed: %v\n", err)
	}
	if expense != expectedExpense {
		tst.Errorf("Test case 1 failed: Unexpected expense value\n")
	}
}

func TestGetExpenseFromStringWithMissingKeyValuePair(tst *testing.T) {
	inputString := "Name: Running shoes\nType: Recreation\nAmount: 25.5\nCurrency: EUR"
	_, err := GetExpenseFromString(inputString)

	if err == nil {
		tst.Errorf("Expected an error for missing key-value pair\n")
	}
}

func TestGetExpenseFromStringWithInvalidValue(tst *testing.T) {
	inputString := "Name: Running shoes\nDate: 01-05-2023\nType: Wrong expense\nAmount: 25.5\nCurrency: EUR"
	_, err := GetExpenseFromString(inputString)

	if err == nil {
		tst.Errorf("Expected an error for missing key-value pair\n")
	}
}

func TestGetExpenseFromStringWithMissingValue(tst *testing.T) {
	inputString := "Name: Running shoes\nDate: 01-05-2023\nType: Recreation\nAmount:\nCurrency: EUR"
	_, err := GetExpenseFromString(inputString)

	if err == nil {
		tst.Errorf("Expected an error for missing key-value pair\n")
	}
}

func TestCheckIfAllFieldsAreFilledValidExpense(tst *testing.T) {
	expense := testExpense

	err := checkIfAllFieldsAreFilled(&expense)
	if err != nil {
		tst.Errorf("Fail: %v\n", err)
	}
}

func TestCheckIfAllFieldsAreFilledMissingNameField(tst *testing.T) {
	expense := Expense{
		Date:     "24-06-2023",
		Type:     Cleaning,
		Amount:   120,
		Currency: "EUR",
	}

	err := checkIfAllFieldsAreFilled(&expense)
	if err == nil {
		tst.Errorf("Expected an error for missing 'Name' field\n")
	}
}

func TestCheckIfAllFieldsAreFilledMissingDateField(tst *testing.T) {
	expense := Expense{
		Name:     "Testing tools",
		Type:     Cleaning,
		Amount:   120,
		Currency: "EUR",
	}

	err := checkIfAllFieldsAreFilled(&expense)
	if err == nil {
		tst.Errorf("Expected an error for missing 'Date' field\n")
	}
}

func TestCheckIfAllFieldsAreFilledMissingTypeField(tst *testing.T) {
	expense := Expense{
		Name:     "Testing tools",
		Date:     "24-06-2023",
		Amount:   120,
		Currency: "EUR",
	}

	err := checkIfAllFieldsAreFilled(&expense)
	if err == nil {
		tst.Errorf("Expected an error for missing 'Type' field\n")
	}
}

func TestCheckIfAllFieldsAreFilledMissingAmountField(tst *testing.T) {
	expense := Expense{
		Name:     "Testing tools",
		Date:     "24-06-2023",
		Type:     Cleaning,
		Currency: "EUR",
	}

	err := checkIfAllFieldsAreFilled(&expense)
	if err == nil {
		tst.Errorf("Expected an error for missing 'Amount' field\n")
	}
}

func TestCheckIfAllFieldsAreFilledMissingCurrencyField(tst *testing.T) {
	expense := Expense{
		Name:   "Testing tools",
		Date:   "24-06-2023",
		Type:   Cleaning,
		Amount: 120,
	}

	err := checkIfAllFieldsAreFilled(&expense)
	if err == nil {
		tst.Errorf("Expected an error for missing 'Currency' field\n")
	}
}

func TestCheckIfAllFieldsAreFilledMissingMultipleFields(tst *testing.T) {
	expense := Expense{
		Name:     "Testing tools",
		Type:     Cleaning,
		Currency: "EUR",
	}

	err := checkIfAllFieldsAreFilled(&expense)
	if err == nil {
		tst.Errorf("Expected an error for missing 'Date' and 'Amount' fields\n")
	}
}

func TestToStringValidString(tst *testing.T) {
	expense := testExpense
	expectedResult := "Name: Testing tools\nDate: 24-06-2023\nType: Cleaning\nAmount: 120.00\nCurrency: EUR\n"

	result, err := expense.ToString()
	if err != nil {
		tst.Errorf("Test failed: %v\n", err)
	}

	if result != expectedResult {
		tst.Errorf("Expected:\n '%s', Got:\n '%s'\n", expectedResult, result)
	}
}

func TestToStringInvalidString(tst *testing.T) {
	expense := Expense{
		Name:     "Bad expense",
		Date:     "01-01-1970",
		Type:     ExpenseType(100),
		Amount:   50,
		Currency: "EUR",
	}

	_, err := expense.ToString()
	expectedError := "ToString: parsing type failed\n"

	if err == nil {
		tst.Errorf("Expected an error\n")
	} else if err.Error() != expectedError {
		tst.Errorf("Expected:\n '%s', Got:\n '%s'\n", expectedError, err.Error())
	}
}
