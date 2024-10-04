package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	middleware "github.com/samuelloganbjss/academy-feedback-tool/admin"
	"github.com/samuelloganbjss/academy-feedback-tool/api"
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
		"tutor_id":   1,
		"student_id": 1,
		"content":    "Test report content",
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
		"student_id": 1,
		"tutor_id": 1,
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
		"student_id": 1,
		"tutor_id": 1,
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

func TestGetStudentReports_ValidRequest(t *testing.T) {
	setup()

	report := map[string]interface{}{
		"tutor_id":   1,
		"student_id": 1,
		"content":    "Sample report content",
	}

	reportJSON, err := json.Marshal(report)
	if err != nil {
		t.Fatalf("Could not marshal report JSON: %v", err)
	}

	reqAdd, err := http.NewRequest("POST", "/api/students/reports", bytes.NewBuffer(reportJSON))
	if err != nil {
		t.Fatal(err)
	}
	reqAdd.Header.Set("Content-Type", "application/json")
	rrAdd := httptest.NewRecorder()
	http.HandlerFunc(tutorAPI.AddReport).ServeHTTP(rrAdd, reqAdd)

	req, err := http.NewRequest("GET", "/admin/students/reports/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Role", "admin")
	rr := httptest.NewRecorder()

	handler := middleware.AdminMiddleware(getTutorRoleFromRequest)(http.HandlerFunc(tutorAPI.GetStudentReports))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var reports []map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &reports)
	if err != nil {
		t.Fatalf("Could not unmarshal response JSON: %v", err)
	}

	if len(reports) != 1 {
		t.Errorf("Expected one report, got %v", len(reports))
	}

	if reports[0]["content"] != "Sample report content" {
		t.Errorf("handler returned wrong report content: got %v want %v", reports[0]["content"], "Sample report content")
	}
}

func TestGetStudentReports_NoReports(t *testing.T) {
	setup()

	req, err := http.NewRequest("GET", "/admin/students/reports/2", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Role", "admin")

	rr := httptest.NewRecorder()

	handler := middleware.AdminMiddleware(getTutorRoleFromRequest)(http.HandlerFunc(tutorAPI.GetStudentReports))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedBody := "[]"
	if rr.Body.String() != expectedBody {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expectedBody)
	}
}
