package utils_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/VilimRoller/Basic-CRUD/data"
	"github.com/VilimRoller/Basic-CRUD/utils"
	"github.com/go-redis/redis"
)

func TestUpdateExpenseHandler(tst *testing.T) {
	//Create new Redis client
	redisClientMock := redis.NewClient(&redis.Options{})

	//Add test data to DB
	expense := testExpense
	key, err := utils.SetExpense(redisClientMock, &expense)

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
		utils.UpdateExpense(writer, request, redisClientMock)
	})

	//Send the request
	handler.ServeHTTP(responseRecorder, putRequest)

	//Check if request was successful
	status := responseRecorder.Code
	if status != http.StatusOK {
		tst.Errorf("Handler returned wrong status code.\n Received: %v Expected: %v\n", status, http.StatusOK)
	}

	expectedResponse := "Update successful!"
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

	updatedExpense, err := utils.RetrieveExpense(redisClientMock, key)
	if err != nil {
		tst.Errorf("Failed to retrieve expense from Redis client mock\n")
		tst.Fatal(err)
	}

	if updatedExpense != expectedUpdatedExpense {
		tst.Errorf("Update failed. Expense value is not equal to expected expense value\n")
	}

}
