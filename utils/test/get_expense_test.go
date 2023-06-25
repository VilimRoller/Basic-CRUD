package utils_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/VilimRoller/Basic-CRUD/utils"
	"github.com/go-redis/redis"
)

func TestGetExpenseHandler(tst *testing.T) {
	//Create new Redis client and add test data
	redisClientMock := redis.NewClient(&redis.Options{})
	expense := testExpense

	key, err := utils.SetExpense(redisClientMock, &expense)
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
		utils.GetExpense(writer, request, redisClientMock)
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
