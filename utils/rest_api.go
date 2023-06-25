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
		GetExpense(writer, request, redisClient)
	}
	addExpenseHandler := func(writer http.ResponseWriter, request *http.Request) {
		PostExpense(writer, request, redisClient)
	}
	updateExpenseHandler := func(writer http.ResponseWriter, request *http.Request) {
		UpdateExpense(writer, request, redisClient)
	}
	deleteExpenseHandler := func(writer http.ResponseWriter, request *http.Request) {
		DeleteExpense(writer, request, redisClient)
	}

	//Define router and http methods
	router := mux.NewRouter()
	router.HandleFunc("/Basic-Crud", Home)
	router.HandleFunc("/Basic-Crud/expenses", getExpenseHandler).Methods("GET")
	router.HandleFunc("/Basic-Crud/expenses", addExpenseHandler).Methods("POST")
	router.HandleFunc("/Basic-Crud/expenses", updateExpenseHandler).Methods("PUT")
	router.HandleFunc("/Basic-Crud/expenses", deleteExpenseHandler).Methods("DELETE")

	http.Handle("/Basic-Crud", router)

	http.ListenAndServe(":8080", router)
}

func Home(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Api is running!")
	writer.WriteHeader(http.StatusOK)
}
