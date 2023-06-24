package utils

import (
	"testing"

	"github.com/VilimRoller/Basic-CRUD/data"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
)

var testExpense data.Expense = data.Expense{
	Name:     "Testing tools",
	Date:     "17-02-2023",
	Type:     data.Cleaning,
	Amount:   123.0,
	Currency: "EUR",
}

func TestSetAndGetExpense(tst *testing.T) {
	// Start a local Redis server for testing
	redisTestServer, err := miniredis.Run()
	if err != nil {
		tst.Fatalf("Failed to start redis test server: %v", err)
	}

	//Close the connection after test
	defer redisTestServer.Close()

	// Create a Redis client using the test server's address
	client := redis.NewClient(&redis.Options{
		Addr: redisTestServer.Addr(),
	})

	expense := testExpense

	key, err := SetExpense(client, &expense)
	if err != nil {
		tst.Errorf("Failed to set expense: %v", err)
	}

	retrievedExpense, err := GetExpense(client, key)
	if err != nil {
		tst.Errorf("Failed to get expense: %v", err)
	}

	if retrievedExpense != expense {
		tst.Errorf("Retrieved expense does not match the original expense")
	}
}

func TestSetExpenseWithKey(tst *testing.T) {
	// Start a local Redis server for testing
	redisTestServer, err := miniredis.Run()
	if err != nil {
		tst.Fatalf("Failed to start redis test server: %v", err)
	}

	//Close the connection after test
	defer redisTestServer.Close()

	// Create a Redis client using the test server's address
	client := redis.NewClient(&redis.Options{
		Addr: redisTestServer.Addr(),
	})

	expense := testExpense

	key := getUniqueKey()
	err = SetExpenseWithKey(client, key, &expense)

	if err != nil {
		tst.Errorf("Failed to set expense with key: %v", err)
	}

	retrievedExpense, err := GetExpense(client, key)
	if err != nil {
		tst.Errorf("Failed to get expense: %v", err)
	}

	if retrievedExpense != expense {
		tst.Errorf("Retrieved expense does not match the original expense")
	}
}
