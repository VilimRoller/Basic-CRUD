package utils

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/VilimRoller/Basic-CRUD/data"
	"github.com/go-redis/redis"
)

func TestHome(tst *testing.T) {
	//Create GET request
	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		tst.Errorf("GET request creation failed\n")
		tst.Fatal(err)
	}

	//Serve the request
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(home)
	handler.ServeHTTP(responseRecorder, request)

	//Validate response
	status := responseRecorder.Code
	if status != http.StatusOK {
		tst.Errorf("Handler returned wrong status code.\n Received: %v Expected: %v\n", status, http.StatusOK)
	}

	expectedResponse := "Api is running!\n"
	if responseRecorder.Body.String() != expectedResponse {
		tst.Errorf("Handler did not return expected body.\n Received: %v Expected: %v\n", responseRecorder.Body.String(), expectedResponse)
	}
}

func TestGetExpenseHandler(tst *testing.T) {
	//Create new Redis client and add test data
	redisClientMock := redis.NewClient(&redis.Options{})
	expense := testExpense

	key, err := SetExpense(redisClientMock, &expense)
	if err != nil {
		tst.Errorf("Set expense failed\n")
		tst.Fatal(err)
	}

	// Create a GET request
	queryParameter := "/Basic-Crud/expenses?key=" + key
	getRequest, err := http.NewRequest("GET", queryParameter, nil)
	if err != nil {
		tst.Errorf("GET request creation failed\n")
		tst.Fatal(err)
	}

	//Response recorder is used to capture the response of http request. Handler is used to pass redis client to request.
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		getExpense(writer, request, redisClientMock)
	})

	//Send the request
	handler.ServeHTTP(responseRecorder, getRequest)

	////Check if request was successful
	status := responseRecorder.Code
	if status != http.StatusOK {
		tst.Errorf("Handler returned wrong status code.\n Received: %v Expected: %v\n", status, http.StatusOK)
	}

	expectedResponse := "Name: Testing tools\nDate: 17-02-2023\nType: Cleaning\nAmount: 123.00\nCurrency: EUR\n"
	if responseRecorder.Body.String() != expectedResponse {
		tst.Errorf("Handler did not return expected body.\n Received: %v Expected: %v\n", responseRecorder.Body.String(), expectedResponse)
	}
}

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

func TestUpdateExpenseHandler(tst *testing.T) {
	//Create new Redis client
	redisClientMock := redis.NewClient(&redis.Options{})

	//Add test data to DB
	expense := testExpense
	key, err := SetExpense(redisClientMock, &expense)

	if err != nil {
		tst.Errorf("Set expense failed\n")
		tst.Fatal(err)
	}

	//Create PUT request
	payloadUpdate := "Name: Testing cools\nDate: 10-02-2023\nType: Cleaning\nAmount: 120.0\nCurrency: USD"
	queryParameter := "/Basic-Crud/expenses?key=" + key
	putRequest, err := http.NewRequest("PUT", queryParameter, strings.NewReader(payloadUpdate))

	if err != nil {
		tst.Errorf("PUT request creation failed\n")
		tst.Fatal(err)
	}

	//Response recorder is used to capture the response of http request. Handler is used to pass redis client to request.
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		updateExpense(writer, request, redisClientMock)
	})

	//Send the request
	handler.ServeHTTP(responseRecorder, putRequest)

	//Check if request was successful
	status := responseRecorder.Code
	if status != http.StatusOK {
		tst.Errorf("Handler returned wrong status code.\n Received: %v Expected: %v\n", status, http.StatusOK)
	}

	expectedResponse := "Update successful!\n"
	if responseRecorder.Body.String() != expectedResponse {
		tst.Errorf("Handler did not return expected body.\n Received: %v Expected: %v\n", responseRecorder.Body.String(), expectedResponse)
	}

	//Check if data was updated successfully
	expectedUpdatedExpense := data.Expense{
		Name:     "Testing cools",
		Date:     "10-02-2023",
		Type:     data.Cleaning,
		Amount:   120.0,
		Currency: "USD",
	}

	updatedExpense, err := GetExpense(redisClientMock, key)
	if err != nil {
		tst.Errorf("Failed to retrieve expense from Redis client mock\n")
		tst.Fatal(err)
	}

	if updatedExpense != expectedUpdatedExpense {
		tst.Errorf("Update failed. Expense value is not equal to expected expense value\n")
	}

}

func TestDeleteExpenseHandler(tst *testing.T) {
	//Create new Redis client
	redisClientMock := redis.NewClient(&redis.Options{})

	//Set data that will be deleted
	expense := testExpense
	key, err := SetExpense(redisClientMock, &expense)
	if err != nil {
		tst.Errorf("Set expense failed\n")
		tst.Fatal(err)
	}

	//Create DELETE request
	queryParameter := "/Basic-Crud/expenses?key=" + key
	deleteRequest, err := http.NewRequest("DELETE", queryParameter, nil)
	if err != nil {
		tst.Errorf("DELETE request creation failed\n")
		tst.Fatal(err)
	}

	//Response recorder is used to capture the response of http request. Handler is used to pass redis client to request.
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		deleteExpense(writer, request, redisClientMock)
	})

	//Send the request
	handler.ServeHTTP(responseRecorder, deleteRequest)

	//Check if request was successful
	status := responseRecorder.Code
	if status != http.StatusOK {
		tst.Errorf("Handler returned wrong status code.\n Received: %v Expected: %v\n", status, http.StatusOK)
	}

	//Check if deletion was successful
	expectedResponse := "Delete successful!\n"
	if responseRecorder.Body.String() != expectedResponse {
		tst.Errorf("Handler did not return expected body.\n Received: %v Expected: %v\n", responseRecorder.Body.String(), expectedResponse)
	}

	// Verify deletion
	_, err = GetExpense(redisClientMock, key)
	if err == nil {
		tst.Errorf("Expense with key %s still exists after deletion\n", key)
	}
}
