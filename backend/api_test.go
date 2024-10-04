package main

import (
	"bytes"
	"encoding/json"
	middleware "github.com/samuelloganbjss/academy-feedback-tool/admin"
	"github.com/samuelloganbjss/academy-feedback-tool/api"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/samuelloganbjss/academy-feedback-tool/repository/tutor"
	"github.com/samuelloganbjss/academy-feedback-tool/service"
)

var tutorAPI *api.TutorAPI

func setup() {

	tutorRepo := tutor.NewInMemoryTutorRepository()

	tutorService := service.NewTutorService(tutorRepo)
	tutorAPI = api.NewTutorAPI(tutorService)
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

	handler := http.HandlerFunc(tutorAPI.AddReport)
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

func TestAdminAccess_AddReport(t *testing.T) {
	setup()

	req, err := http.NewRequest("POST", "/api/students/reports", bytes.NewBuffer([]byte(`{
		"studentID": 1,
		"tutorID": 1,
		"content": "New report from admin"
	}`)))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Role", "admin")

	rr := httptest.NewRecorder()

	handler := middleware.AdminMiddleware(getTutorRoleFromRequest)(http.HandlerFunc(tutorAPI.AddReport))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestNonAdminAccess_AddReport(t *testing.T) {

	setup()

	req, err := http.NewRequest("POST", "/api/students/reports", bytes.NewBuffer([]byte(`{
		"studentID": 1,
		"tutorID": 1,
		"content": "New report"
	}`)))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Role", "tutor")

	rr := httptest.NewRecorder()

	handler := middleware.AdminMiddleware(getTutorRoleFromRequest)(http.HandlerFunc(tutorAPI.AddReport))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusForbidden {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusForbidden)
	}
}
