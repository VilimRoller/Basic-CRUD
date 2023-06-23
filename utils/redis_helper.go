package utils

import (
	"fmt"

	"github.com/VilimRoller/Basic-CRUD/data"

	"github.com/go-redis/redis"

	"encoding/json"

	"errors"

	"github.com/google/uuid"
)

func GetDefaultRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return client
}

func SetExpenseWithKey(redisClient *redis.Client, key string, expense *data.Expense) error {
	jsonVal, err := json.Marshal(expense)
	if err != nil {
		return errors.New("SetExpenseWithKey: Json marshal failed")
	}

	err = redisClient.Set(key, jsonVal, 0).Err()

	if err != nil {
		return errors.New("SetExpenseWithKey: Adding data to DB failed")
	}

	return nil
}

func SetExpense(redisClient *redis.Client, expense *data.Expense) (string, error) {
	jsonVal, err := json.Marshal(expense)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("SetExpense: Json marshal failed")
	}

	key := getUniqueKey()

	err = redisClient.Set(key, jsonVal, 0).Err()

	if err != nil {
		fmt.Println(err)
		return "", errors.New("SetExpense: Adding data to DB failed")
	}

	return key, nil
}

func GetExpense(redisClient *redis.Client, key string) (data.Expense, error) {
	returnValueString, err := redisClient.Get(key).Result()

	if err != nil {
		return data.EmptyExpense, errors.New("GetExpense: Retrieving key from DB failed")
	}

	var returnValue data.Expense

	err = json.Unmarshal([]byte(returnValueString), &returnValue)

	if err != nil {
		fmt.Println(err)
		return data.EmptyExpense, errors.New("GetExpense: Json unmarshal failed")
	}

	return returnValue, nil
}

func getUniqueKey() string {
	uniqueId := uuid.New()
	return uniqueId.String()
}
