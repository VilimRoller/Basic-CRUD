package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/VilimRoller/Basic-CRUD/data"
	"github.com/go-redis/redis"
)

func PostExpense(writer http.ResponseWriter, request *http.Request, redisClient *redis.Client) {
	expense := getExpenseFromRequestBody(writer, request)

	key := setExpense(writer, expense, redisClient)

	fmt.Fprint(writer, "key = ", key)
	writer.WriteHeader(http.StatusOK)
}

func setExpense(writer http.ResponseWriter, expense data.Expense, redisClient *redis.Client) string {
	key, err := SetExpense(redisClient, &expense)

	if err != nil {
		writer.WriteHeader(http.StatusNotImplemented)
		fmt.Fprint(writer, "Failed to add value to DB\n")
		fmt.Fprint(writer, err)
		fmt.Fprint(writer, "\n")
		return ""
	}

	return key
}

func setExpenseWithKey(writer http.ResponseWriter, expense data.Expense, key string, redisClient *redis.Client) error {
	err := SetExpenseWithKey(redisClient, key, &expense)

	if err != nil {
		writer.WriteHeader(http.StatusNotImplemented)
		fmt.Fprint(writer, "Failed to update value in DB\n")
		fmt.Fprint(writer, err)
		fmt.Fprint(writer, "\n")
		return err
	}

	return nil
}

func getExpenseFromRequestBody(writer http.ResponseWriter, request *http.Request) data.Expense {
	requestBody, err := ioutil.ReadAll(request.Body)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "Failed to read request body\n")
		fmt.Fprint(writer, err)
		fmt.Fprint(writer, "\n")
		return data.Expense{}
	}

	expenseString := string(requestBody)

	expense, err := data.GetExpenseFromString(expenseString)

	if err != nil {
		writer.WriteHeader(http.StatusNotImplemented)
		fmt.Fprint(writer, "Failed to parse string\n")
		fmt.Fprint(writer, err)
		fmt.Fprint(writer, "\n")
		return data.Expense{}
	}

	return expense
}
