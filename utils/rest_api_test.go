package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"
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

	expectedResponse := "Api is running!"
	if responseRecorder.Body.String() != expectedResponse {
		tst.Errorf("Handler did not return expected body.\n Received: %v Expected: %v\n", responseRecorder.Body.String(), expectedResponse)
	}
}
