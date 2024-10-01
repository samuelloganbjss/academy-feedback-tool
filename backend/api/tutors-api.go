package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/samuelloganbjss/academy-feedback-tool/model"
	"github.com/samuelloganbjss/academy-feedback-tool/service"
)

type TutorAPI struct {
	tutorService *service.TutorService
}

func NewTutorAPI(tutorService *service.TutorService) *TutorAPI {
	return &TutorAPI{
		tutorService: tutorService,
	}
}
func (api *TutorAPI) GetTutors(writer http.ResponseWriter, request *http.Request) {
	students, err := api.tutorService.GetTutorsService()

	if err != nil {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(writer).Encode(students)
}

func (api *TutorAPI) AddReport(writer http.ResponseWriter, request *http.Request) {
	var report model.Report
	err := json.NewDecoder(request.Body).Decode(&report)

	if err != nil {
		http.Error(writer, "Bad Request", http.StatusBadRequest)
		return
	}

	createdReport, err := api.tutorService.AddReportService(report)
	if err != nil {
		http.Error(writer, "Error adding report", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(createdReport)
}

func (api *TutorAPI) EditReport(writer http.ResponseWriter, request *http.Request) {
	idStr := request.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(writer, "Invalid ID", http.StatusBadRequest)
		return
	}

	var newContent struct {
		Content string `json:"content"`
	}

	err = json.NewDecoder(request.Body).Decode(&newContent)
	if err != nil {
		http.Error(writer, "Bad Request", http.StatusBadRequest)
		return
	}

	tutorID := 1 // Placeholder: Replace with actual tutor ID after authentication

	updatedReport, err := api.tutorService.EditReportService(id, newContent.Content, tutorID)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusUnauthorized)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(updatedReport)
}
