package utils

import (
	"fmt"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

func RegisterEndpoints(redisClient *redis.Client) {
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
	fmt.Fprintf(writer, "Api is running!")
	writer.WriteHeader(http.StatusOK)
}
