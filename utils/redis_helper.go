package utils

import (
	"encoding/json"
	"errors"

	"github.com/VilimRoller/Basic-CRUD/data"
	"github.com/go-redis/redis"
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
		return errors.New("SetExpenseWithKey: Json marshal failed\nError: " + err.Error() + "\n")
	}

	err = redisClient.Set(key, jsonVal, 0).Err()

	if err != nil {
		return errors.New("SetExpenseWithKey: Adding data to DB failed\nError: " + err.Error() + "\n")
	}

	return nil
}

func SetExpense(redisClient *redis.Client, expense *data.Expense) (string, error) {
	jsonVal, err := json.Marshal(expense)
	if err != nil {
		return "", errors.New("SetExpense: Json marshal failed\nError: " + err.Error() + "\n")
	}

	key := getUniqueKey()

	err = redisClient.Set(key, jsonVal, 0).Err()

	if err != nil {
		return "", errors.New("SetExpense: Adding data to DB failed\nError: " + err.Error() + "\n")
	}

	return key, nil
}

func GetExpense(redisClient *redis.Client, key string) (data.Expense, error) {
	returnValueString, err := redisClient.Get(key).Result()

	if err != nil {
		return data.EmptyExpense, errors.New("GetExpense: Retrieving key from DB failed\nError: " + err.Error() + "\n")
	}

	var returnValue data.Expense

	err = json.Unmarshal([]byte(returnValueString), &returnValue)

	if err != nil {
		return data.EmptyExpense, errors.New("GetExpense: Json unmarshal failed\nError: " + err.Error() + "\n")
	}

	return returnValue, nil
}

func getUniqueKey() string {
	uniqueId := uuid.New()
	return uniqueId.String()
}
