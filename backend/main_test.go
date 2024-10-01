package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"io"
)

func TestRootHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder
	rr := httptest.NewRecorder()
	// Create a handler
	handler := http.HandlerFunc(rootHandler)
	//ACT
	// Serve the request
	handler.ServeHTTP(rr, req)

	    //ASSERT
    // Check the status code
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    // Check the response body
    expected := "Hello, welcome to the Feedback tool for the academy!"
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
    }
}

func TestRootHandlerWithServer(t *testing.T) {
	//ARRANGE
	// Create a new server with the handler
	server := httptest.NewServer(http.HandlerFunc(rootHandler))
	defer server.Close()

	//ACT
	// Send a GET request to the server
	resp, err := http.Get(server.URL + "/")
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}
	defer resp.Body.Close()

	// Check the status code
	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := "Hello, welcome to the Feedback tool for the academy!"
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	if string(bodyBytes) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", string(bodyBytes), expected)
	}
}
