package utils

import (
	"fmt"

	expenses "github.com/VilimRoller/Basic-CRUD/data"

	"github.com/go-redis/redis"

	"encoding/json"
)

func GetDefaultRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return client
}

func SetExpense(redisClient *redis.Client, key string, expense expenses.Expense) {
	jsonVal, err := json.Marshal(expense)
	if err != nil {
		fmt.Println(err)
	}

	redisErr := redisClient.Set(key, jsonVal, 0)

	if redisErr != nil {
		fmt.Println(redisErr)
	}
}

func GetExpense(redisClient *redis.Client, key string) expenses.Expense {

	returnValueString, err := redisClient.Get(key).Result()

	if err != nil {
		fmt.Println(err)
		return expenses.EmptyExpense
	}

	var returnValue expenses.Expense

	err = json.Unmarshal([]byte(returnValueString), &returnValue)

	if err != nil {
		fmt.Println(err)
		return expenses.EmptyExpense
	}

	return returnValue
}
