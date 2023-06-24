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
		{BadType, "", errors.New("GetExpenseTypeAsString: unknown expense type")},
	}

	for _, testCase := range testCases {
		output, err := GetExpenseTypeAsString(testCase.expenseType)

		// Check if the output matches the expected value
		if output != testCase.expectedOutput {
			tst.Errorf("Expected: %s, Got: %s", testCase.expectedOutput, output)
		}

		// Check if the error matches the expected error
		if (err == nil && testCase.expectedError != nil) ||
			(err != nil && testCase.expectedError == nil) ||
			(err != nil && testCase.expectedError != nil && err.Error() != testCase.expectedError.Error()) {
			tst.Errorf("Expected error: %v, Got error: %v", testCase.expectedError, err)
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
		{"Unknown", BadType, errors.New("GetStringAsExpenseType: unknown expense type")},
	}

	for _, testCase := range testCases {
		output, err := GetStringAsExpenseType(testCase.input)

		// Check if the output matches the expected value
		if output != testCase.expectedOutput {
			tst.Errorf("Expected: %d, Got: %d", testCase.expectedOutput, output)
		}

		// Check if the error matches the expected error
		if (err == nil && testCase.expectedError != nil) ||
			(err != nil && testCase.expectedError == nil) ||
			(err != nil && testCase.expectedError != nil && err.Error() != testCase.expectedError.Error()) {
			tst.Errorf("Expected error: %v, Got error: %v", testCase.expectedError, err)
		}
	}
}

func TestAddKeyAndValueToExpenseUpdateName(tst *testing.T) {
	expense := testExpense
	err := addKeyAndValueToExpense("Name", "Testing equipment", &expense)

	if err != nil {
		tst.Errorf("Adding key and value to expense failed: %v", err)
	}
	if expense.Name != "Testing equipment" {
		tst.Errorf("Name was not updated correctly")
	}
}

func TestAddKeyAndValueToExpenseUpdateDate(tst *testing.T) {
	expense := testExpense
	err := addKeyAndValueToExpense("Date", "01-01-1970", &expense)

	if err != nil {
		tst.Errorf("Adding key and value to expense failed: %v", err)
	}
	if expense.Date != "01-01-1970" {
		tst.Errorf("Date was not updated correctly")
	}
}

func TestAddKeyAndValueToExpenseUpdateExpenseType(tst *testing.T) {
	expense := testExpense
	err := addKeyAndValueToExpense("Type", "Entertainment", &expense)

	if err != nil {
		tst.Errorf("Adding key and value to expense failed: %v", err)
	}
	if expense.Type != Entertainment {
		tst.Errorf("Expense type was not updated correctly")
	}
}

func TestAddKeyAndValueToExpenseUpdateAmount(tst *testing.T) {
	expense := testExpense
	err := addKeyAndValueToExpense("Amount", "150", &expense)

	if err != nil {
		tst.Errorf("Adding key and value to expense failed: %v", err)
	}
	if expense.Amount != float32(150) {
		tst.Errorf("Amount was not updated correctly. Amount = %.2f expected amount = 150", expense.Amount)
	}
}

func TestAddKeyAndValueToExpenseUpdateAmountWithEmptyInput(tst *testing.T) {
	expense := testExpense
	originalAmount := expense.Amount
	err := addKeyAndValueToExpense("Amount", "", &expense)

	if err != nil {
		tst.Errorf("Adding key and value to expense failed: %v", err)
	}
	if expense.Amount != originalAmount {
		tst.Errorf("Amount changed. Amount = %.2f expected amount = %.2f", expense.Amount, originalAmount)
	}
}

func TestAddKeyAndValueToExpenseUpdateCurrecny(tst *testing.T) {
	expense := testExpense
	err := addKeyAndValueToExpense("Currency", "USD", &expense)

	if err != nil {
		tst.Errorf("Adding key and value to expense failed: %v", err)
	}
	if expense.Currency != "USD" {
		tst.Errorf("Amount was not updated correctly")
	}
}

func TestAddKeyAndValueToExpenseUpdateCurrecnyWithBadInput(tst *testing.T) {
	expense := testExpense
	err := addKeyAndValueToExpense("Currency", "US", &expense)

	if err == nil {
		tst.Errorf("Error was expected for bad currency input")
	}
}

func TestAddKeyAndValueToExpenseUnknownKey(tst *testing.T) {
	expense := testExpense
	err := addKeyAndValueToExpense("Bip bop", "Bing bong", &expense)

	if err == nil {
		tst.Errorf("Error was expected for unknown key")
	}
}

func TestAddKeyAndValueToExpenseEmptyKey(tst *testing.T) {
	expense := testExpense
	err := addKeyAndValueToExpense("", "Something", &expense)

	if err == nil {
		tst.Errorf("Error was expected for unknown key")
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
		tst.Errorf("Test case 1 failed: %v", err)
	}
	if expense != expectedExpense {
		tst.Errorf("Test case 1 failed: Unexpected expense value")
	}
}

func TestGetExpenseFromStringWithMissingKeyValuePair(tst *testing.T) {
	inputString := "Name: Running shoes\nType: Recreation\nAmount: 25.5\nCurrency: EUR"
	_, err := GetExpenseFromString(inputString)

	if err == nil {
		tst.Errorf("Expected an error for missing key-value pair")
	}
}

func TestGetExpenseFromStringWithInvalidValue(tst *testing.T) {
	inputString := "Name: Running shoes\nDate: 01-05-2023\nType: Wrong expense\nAmount: 25.5\nCurrency: EUR"
	_, err := GetExpenseFromString(inputString)

	if err == nil {
		tst.Errorf("Expected an error for missing key-value pair")
	}
}

func TestGetExpenseFromStringWithMissingValue(tst *testing.T) {
	inputString := "Name: Running shoes\nDate: 01-05-2023\nType: Recreation\nAmount:\nCurrency: EUR"
	_, err := GetExpenseFromString(inputString)

	if err == nil {
		tst.Errorf("Expected an error for missing key-value pair")
	}
}

func TestCheckIfAllFieldsAreFilledValidExpense(tst *testing.T) {
	expense := testExpense

	err := checkIfAllFieldsAreFilled(&expense)
	if err != nil {
		tst.Errorf("Fail: %v", err)
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
		tst.Errorf("Expected an error for missing 'Name' field")
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
		tst.Errorf("Expected an error for missing 'Date' field")
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
		tst.Errorf("Expected an error for missing 'Type' field")
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
		tst.Errorf("Expected an error for missing 'Amount' field")
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
		tst.Errorf("Expected an error for missing 'Currency' field")
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
		tst.Errorf("Expected an error for missing 'Date' and 'Amount' fields")
	}
}

func TestToStringValidString(tst *testing.T) {
	expense := testExpense
	expectedResult := "Name: Testing tools\nDate: 24-06-2023\nType: Cleaning\nAmount: 120.00\nCurrency: EUR\n"

	result, err := expense.ToString()
	if err != nil {
		tst.Errorf("Test failed: %v", err)
	}

	if result != expectedResult {
		tst.Errorf("Expected:\n '%s', Got:\n '%s'", expectedResult, result)
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
	expectedError := "ToString: parsing type failed."

	if err == nil {
		tst.Errorf("Expected an error")
	} else if err.Error() != expectedError {
		tst.Errorf("Expected:\n '%s', Got:\n '%s'", expectedError, err.Error())
	}
}
