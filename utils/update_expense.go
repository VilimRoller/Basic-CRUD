package utils

import (
	"fmt"
	"net/http"

	"github.com/go-redis/redis"
)

func UpdateExpense(writer http.ResponseWriter, request *http.Request, redisClient *redis.Client) {
	key := getKeyFromRequest(request)
	expense := getExpenseFromRequestBody(writer, request)

	err := setExpenseWithKey(writer, expense, key, redisClient)

	if err != nil {
		fmt.Fprint(writer, err)
		fmt.Fprint(writer, "\n")
		return
	}

	fmt.Fprintf(writer, "Update successful!")
	writer.WriteHeader(http.StatusOK)
}
