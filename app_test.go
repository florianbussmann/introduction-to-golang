package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var app App

func TestMain(m *testing.M) {
	err := app.Initialise()
	if err != nil {
		log.Fatal(err)
	}
	m.Run()
}

func TestGetClient(t *testing.T) {
	request, _ := http.NewRequest("GET", "/client/1", nil)
	response := sendRequest(request)
	checkStatusCode(t, http.StatusOK, response.Code)
}

func checkStatusCode(t *testing.T, expectedStatusCode int, actualStatusCode int) {
	if expectedStatusCode != actualStatusCode {
		t.Errorf("Expected status: %v, Received %v", expectedStatusCode, actualStatusCode)
	}
}

func sendRequest(request *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	app.Router.ServeHTTP(recorder, request)
	return recorder
}
