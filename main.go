package main

import (
	"fmt"

	expenses "github.com/VilimRoller/Basic-CRUD/data"

	"github.com/VilimRoller/Basic-CRUD/utils"
)

func main() {
	expense1 := expenses.Expense{
		Name:     "Bread",
		Date:     "01-06-2023",
		Type:     expenses.Food,
		Amount:   1.4,
		Currency: "EUR",
	}

	redisClient := utils.GetDefaultRedisClient()

	utils.SetExpense(redisClient, "12345", expense1)

	expense2 := utils.GetExpense(redisClient, "12345")

	if expense1 == expense2 {
		fmt.Println("It works!")
	}
}
