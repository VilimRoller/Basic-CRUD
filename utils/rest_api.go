package utils

import (
	"fmt"

	"net/http"

	"github.com/gorilla/mux"
)

func RegisterEndpoints() {
	fmt.Println("Registering endpoints")

	router := mux.NewRouter()
	router.HandleFunc("/", Home)

	http.Handle("/", router)

	http.ListenAndServe(":8080", router)
}

func Home(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Home")
	fmt.Fprintf(writer, "Api is running!")
	writer.WriteHeader(http.StatusOK)
}
