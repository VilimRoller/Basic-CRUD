package utils

import (
	"fmt"
	"net/http"

	"github.com/go-redis/redis"
)

func deleteExpense(writer http.ResponseWriter, request *http.Request, redisClient *redis.Client) {
	key := getKeyFromRequest(request)

	if key == "all" {
		deleteAllExpenses(writer, redisClient)
		return
	}

	err := redisClient.Del(key).Err()

	if err != nil {
		fmt.Fprintf(writer, "Delete failed\n")
		fmt.Fprint(writer, err)
		fmt.Fprint(writer, "\n")
		return
	}

	fmt.Fprintf(writer, "Delete successful!")
	writer.WriteHeader(http.StatusOK)
}

func deleteAllExpenses(writer http.ResponseWriter, redisClient *redis.Client) {
	err := redisClient.FlushAll().Err()

	if err != nil {
		fmt.Fprintf(writer, "FlushAll failed\n")
		fmt.Fprint(writer, err)
		fmt.Fprint(writer, "\n")
	}
}
