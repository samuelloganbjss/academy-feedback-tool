package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/samuelloganbjss/academy-feedback-tool/api"
	"github.com/samuelloganbjss/academy-feedback-tool/db"
	"github.com/samuelloganbjss/academy-feedback-tool/service"
)

var studentAPI *api.StudentAPI

func setup() {

	dbRepo := db.NewInMemoryRepository()

	studentService := service.NewStudentService(dbRepo)
	studentAPI = api.NewStudentAPI(studentService)
}

func TestAddReport(t *testing.T) {
	setup()

	report := map[string]interface{}{
		"tutorID":   1,
		"studentID": 1,
		"content":   "Test report content",
	}
	reportJSON, err := json.Marshal(report)
	if err != nil {
		t.Fatalf("Could not marshal report JSON: %v", err)
	}

	req, err := http.NewRequest("POST", "/api/students/reports", bytes.NewBuffer(reportJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(studentAPI.AddReport)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var addedReport map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &addedReport)
	if err != nil {
		t.Fatalf("Could not unmarshal response JSON: %v", err)
	}

	if addedReport["content"] != "Test report content" {
		t.Errorf("handler returned wrong report content: got %v want %v", addedReport["content"], "Test report content")
	}

	if addedReport["id"] == nil {
		t.Errorf("Expected report ID to be set, but got nil")
	}

}

func TestEditReport(t *testing.T) {
	setup()

	initialReport := map[string]interface{}{
		"tutorID":   1,
		"studentID": 1,
		"content":   "Initial report content",
	}
	initialReportJSON, err := json.Marshal(initialReport)
	if err != nil {
		t.Fatalf("Could not marshal initial report JSON: %v", err)
	}

	reqAdd, err := http.NewRequest("POST", "/api/students/reports", bytes.NewBuffer(initialReportJSON))
	if err != nil {
		t.Fatal(err)
	}
	reqAdd.Header.Set("Content-Type", "application/json")

	rrAdd := httptest.NewRecorder()
	handlerAdd := http.HandlerFunc(studentAPI.AddReport)
	handlerAdd.ServeHTTP(rrAdd, reqAdd)

	if status := rrAdd.Code; status != http.StatusOK {
		t.Errorf("Add report handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var addedReport map[string]interface{}
	err = json.Unmarshal(rrAdd.Body.Bytes(), &addedReport)
	if err != nil {
		t.Fatalf("Could not unmarshal added report JSON: %v", err)
	}
	reportID := int(addedReport["id"].(float64))

	editedContent := map[string]interface{}{
		"content": "Updated report content",
	}
	editedContentJSON, err := json.Marshal(editedContent)
	if err != nil {
		t.Fatalf("Could not marshal edited report JSON: %v", err)
	}

	editURL := "/api/students/reports/edit?id=" + strconv.Itoa(reportID)
	reqEdit, err := http.NewRequest("PUT", editURL, bytes.NewBuffer(editedContentJSON))
	if err != nil {
		t.Fatal(err)
	}
	reqEdit.Header.Set("Content-Type", "application/json")

	rrEdit := httptest.NewRecorder()
	handlerEdit := http.HandlerFunc(studentAPI.EditReport)
	handlerEdit.ServeHTTP(rrEdit, reqEdit)

	if status := rrEdit.Code; status != http.StatusOK {
		t.Errorf("Edit report handler returned wrong status code: got %v want %v", status, http.StatusOK)
		t.Logf("Response body: %s", rrEdit.Body.String())
	}

	var updatedReport map[string]interface{}
	err = json.Unmarshal(rrEdit.Body.Bytes(), &updatedReport)
	if err != nil {
		t.Fatalf("Could not unmarshal edited report JSON: %v", err)
	}

	if updatedReport["content"] != "Updated report content" {
		t.Errorf("handler returned wrong report content: got %v want %v", updatedReport["content"], "Updated report content")
	}
}
