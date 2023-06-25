package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/VilimRoller/Basic-CRUD/data"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

func RegisterEndpoints(redisClient *redis.Client) {
	fmt.Println("Application is running!")

	//Define closures
	getExpenseHandler := func(writer http.ResponseWriter, request *http.Request) {
		getExpense(writer, request, redisClient)
	}
	addExpenseHandler := func(writer http.ResponseWriter, request *http.Request) {
		addExpense(writer, request, redisClient)
	}
	updateExpenseHandler := func(writer http.ResponseWriter, request *http.Request) {
		updateExpense(writer, request, redisClient)
	}
	deleteExpenseHandler := func(writer http.ResponseWriter, request *http.Request) {
		deleteExpense(writer, request, redisClient)
	}

	//Define router and http methods
	router := mux.NewRouter()
	router.HandleFunc("/Basic-Crud", home)
	router.HandleFunc("/Basic-Crud/expenses", getExpenseHandler).Methods("GET")
	router.HandleFunc("/Basic-Crud/expenses", addExpenseHandler).Methods("POST")
	router.HandleFunc("/Basic-Crud/expenses", updateExpenseHandler).Methods("PUT")
	router.HandleFunc("/Basic-Crud/expenses", deleteExpenseHandler).Methods("DELETE")

	http.Handle("/Basic-Crud", router)

	http.ListenAndServe(":8080", router)
}

func home(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Api is running!\n")
	writer.WriteHeader(http.StatusOK)
}

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
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	stringResult, err := result.ToString()

	if err != nil {
		fmt.Fprintf(writer, "Converting result to string failed\n")
		writer.WriteHeader(http.StatusNotImplemented)
		return
	}

	fmt.Fprintf(writer, stringResult)

	writer.WriteHeader(http.StatusOK)
}

func addExpense(writer http.ResponseWriter, request *http.Request, redisClient *redis.Client) {
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
		return ""
	}

	return key
}

func setExpenseWithKey(writer http.ResponseWriter, expense data.Expense, key string, redisClient *redis.Client) error {
	err := SetExpenseWithKey(redisClient, key, &expense)

	if err != nil {
		writer.WriteHeader(http.StatusNotImplemented)
		fmt.Fprint(writer, "Failed to update value in DB\n")
		return err
	}

	return nil
}

func getExpenseFromRequestBody(writer http.ResponseWriter, request *http.Request) data.Expense {
	requestBody, err := ioutil.ReadAll(request.Body)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "Failed to read request body\n")
		return data.Expense{}
	}

	expenseString := string(requestBody)

	expense, err := data.GetExpenseFromString(expenseString)

	if err != nil {
		writer.WriteHeader(http.StatusNotImplemented)
		fmt.Fprint(writer, "Failed to parse string\n")
		return data.Expense{}
	}

	return expense
}

func updateExpense(writer http.ResponseWriter, request *http.Request, redisClient *redis.Client) {
	key := getKeyFromRequest(request)
	expense := getExpenseFromRequestBody(writer, request)

	err := setExpenseWithKey(writer, expense, key, redisClient)

	if err != nil {
		return
	}

	fmt.Fprintf(writer, "Update successful!\n")
	writer.WriteHeader(http.StatusOK)
}

func deleteExpense(writer http.ResponseWriter, request *http.Request, redisClient *redis.Client) {
	key := getKeyFromRequest(request)

	if key == "all" {
		deleteAllExpenses(writer, redisClient)
		return
	}

	err := redisClient.Del(key).Err()

	if err != nil {
		fmt.Fprintf(writer, "Delete failed\n")
		return
	}

	fmt.Fprintf(writer, "Delete successful!\n")
	writer.WriteHeader(http.StatusOK)
}

func deleteAllExpenses(writer http.ResponseWriter, redisClient *redis.Client) {
	err := redisClient.FlushAll().Err()

	if err != nil {
		fmt.Fprintf(writer, "FlushAll failed\n")
	}
}
