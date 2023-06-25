package utils

import (
	"fmt"
	"net/http"

	"github.com/go-redis/redis"
)

func getExpense(writer http.ResponseWriter, request *http.Request, redisClient *redis.Client) {
	key := getKeyFromRequest(request)

	if key == "all" {
		getAllData(writer, redisClient)
		return
	}

	getExpenseWithKey(writer, key, redisClient)
}

func getKeyFromRequest(request *http.Request) string {
	queryParams := request.URL.Query()
	return queryParams.Get("key")
}

func getAllData(writer http.ResponseWriter, redisClient *redis.Client) {
	keys, err := redisClient.Keys("*").Result()

	if err != nil {
		fmt.Fprintf(writer, "Failed to retrieve all data\n")
		fmt.Fprint(writer, err)
		fmt.Fprint(writer, "\n")
		return
	}

	for _, key := range keys {
		fmt.Fprintf(writer, "key = ")
		fmt.Fprintf(writer, key)
		fmt.Fprintf(writer, "\n\n")
		getExpenseWithKey(writer, key, redisClient)
		fmt.Fprintf(writer, "\n")
	}

}

func getExpenseWithKey(writer http.ResponseWriter, key string, redisClient *redis.Client) {
	result, err := GetExpense(redisClient, key)

	if err != nil {
		fmt.Fprintf(writer, "Key not found\n")
		fmt.Fprint(writer, err)
		fmt.Fprint(writer, "\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	stringResult, err := result.ToString()

	if err != nil {
		fmt.Fprintf(writer, "Converting result to string failed\n")
		fmt.Fprint(writer, err)
		fmt.Fprint(writer, "\n")
		writer.WriteHeader(http.StatusNotImplemented)
		return
	}

	fmt.Fprintf(writer, stringResult)

	writer.WriteHeader(http.StatusOK)
}
