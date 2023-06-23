package utils

import (
	"fmt"

	"net/http"

	"github.com/gorilla/mux"
)

func RegisterEndpoints() {
	fmt.Println("Registering endpoints")

	router := mux.NewRouter()
	router.HandleFunc("/Basic-Crud", Home)
	router.HandleFunc("/Basic-Crud/expenses", GetExpenses).Methods("GET")
	router.HandleFunc("/Basic-Crud/expenses", AddExpense).Methods("POST")
	router.HandleFunc("/Basic-Crud/expenses", UpdateExpense).Methods("PUT")
	router.HandleFunc("/Basic-Crud/expenses", DeleteExpense).Methods("DELETE")

	http.Handle("/Basic-Crud", router)

	http.ListenAndServe(":8080", router)
}

func Home(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Api is running!")
	writer.WriteHeader(http.StatusOK)
}

func GetExpenses(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "GetExpense")
	writer.WriteHeader(http.StatusOK)
}

func AddExpense(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "AddExpense")
	writer.WriteHeader(http.StatusOK)
}

func UpdateExpense(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "UpdateExpense")
	writer.WriteHeader(http.StatusOK)
}

func DeleteExpense(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "DeleteExpense")
	writer.WriteHeader(http.StatusOK)
}
