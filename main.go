package main

import (
	"fmt"

	"github.com/VilimRoller/Basic-CRUD/utils"

	"github.com/VilimRoller/Basic-CRUD/data"
)

func main() {
	expense1 := data.Expense{
		Name:     "Milk",
		Date:     "01-07-2023",
		Type:     data.Food,
		Amount:   2,
		Currency: "EUR",
	}

	redisClient := utils.GetDefaultRedisClient()

	key1, _ := utils.SetExpense(redisClient, &expense1)

	fmt.Println(key1)

	expense2, _ := utils.GetExpense(redisClient, key1)

	if expense1 == expense2 {
		fmt.Println("It works!")
	}

	utils.RegisterEndpoints(redisClient)

}
