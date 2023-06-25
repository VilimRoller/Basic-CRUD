package utils

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-redis/redis"
)

func TestAddExpenseHandler(tst *testing.T) {
	//Create new Redis client
	redisClientMock := redis.NewClient(&redis.Options{})

	//Create POST request
	payload := "Name: Testing tools\nDate: 17-02-2023\nType: Cleaning\nAmount: 123.0\nCurrency: EUR"
	postRequest, err := http.NewRequest("POST", "/Basic-Crud/expenses", strings.NewReader(payload))
	if err != nil {
		tst.Errorf(("POST request creation failed\n"))
	}

	//Response recorder is used to capture the response of http request. Handler is used to pass redis client to request.
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		addExpense(writer, request, redisClientMock)
	})

	//Send the request
	handler.ServeHTTP(responseRecorder, postRequest)

	//Check if request was successful
	status := responseRecorder.Code
	if status != http.StatusOK {
		tst.Errorf("Handler returned wrong status code.\n Received: %v Expected: %v\n", status, http.StatusOK)
		tst.Fatal(err)
	}

	// Retrieve the response body
	responseBody, err := ioutil.ReadAll(responseRecorder.Body)
	if err != nil {
		tst.Errorf(("Failed to read response body\n"))
		tst.Fatal(err)
	}

	// Verify that the response body contains the key
	if !strings.Contains(string(responseBody), "key = ") {
		tst.Errorf("Handler did not return expected body. Reponse body:%v\n", string(responseBody))
		tst.Fatal(err)
	}

	//Extract the key. Key is needed to check if added expense exists in DB
	keyPrefix := "key = "
	keyStartIndex := len(keyPrefix)
	key := string(responseBody[keyStartIndex:])

	//Check if adding was successful
	expectedResponse := testExpense
	receivedReponse, err := GetExpense(redisClientMock, key)

	if err != nil {
		tst.Errorf("Failed to retrieve expense from Redis client mock\n")
		tst.Fatal(err)
	}

	if expectedResponse != receivedReponse {
		tst.Errorf("Received response is not equal to expected response\n")
	}
}
