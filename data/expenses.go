package data

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

func GetAsString(expenseType ExpenseType) string {
	switch expenseType {
	case Food:
		return "Food"
	case Cleaning:
		return "Cleaning"
	case Utilities:
		return "Utilities"
	case Transportation:
		return "Transportation"
	case Entertainment:
		return "Entertainment"
	case Health:
		return "Health"
	case Recreation:
		return "Recreation"
	default:
		return "Other"
	}
}
