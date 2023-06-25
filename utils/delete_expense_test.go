package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-redis/redis"
)

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
	expectedResponse := "Delete successful!"
	if responseRecorder.Body.String() != expectedResponse {
		tst.Errorf("Handler did not return expected body.\n Received: %v Expected: %v\n", responseRecorder.Body.String(), expectedResponse)
	}

	// Verify deletion
	_, err = GetExpense(redisClientMock, key)
	if err == nil {
		tst.Errorf("Expense with key %s still exists after deletion\n", key)
	}
}
