package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/samuelloganbjss/academy-feedback-tool/model"
	"github.com/samuelloganbjss/academy-feedback-tool/service"
	"fmt"
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

func (api *TutorAPI) AddTutor(writer http.ResponseWriter, request *http.Request) {
	var tutor model.Tutor
	err := json.NewDecoder(request.Body).Decode(&tutor)

	if err != nil {
		http.Error(writer, "Bad Request", http.StatusBadRequest)
		return
	}

	createdTutor, err := api.tutorService.AddTutorService(tutor)
	if err != nil {
		http.Error(writer, "Error adding tutor", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(createdTutor)
}

func (api *TutorAPI) DeleteSingleTutor(writer http.ResponseWriter, request *http.Request) {

	id, err := api.parseId(request.PathValue("id"))

	if err != nil {
        http.Error(writer, "Bad Request ID", http.StatusBadRequest)
        return
    }

    _, err = api.tutorService.DeleteTutorService(id)
    
    if err != nil {
        http.Error(writer, "Could not delete tutor", http.StatusBadRequest)
        return
    }

    writer.WriteHeader(http.StatusOK)

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
	studentID, err := api.parseId(request.PathValue("id"))

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

	updatedReport, err := api.tutorService.EditReportService(studentID, newContent.Content, tutorID)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusUnauthorized)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(updatedReport)
}


func (api *TutorAPI) GetStudentReports(writer http.ResponseWriter, request *http.Request) {
	studentID, err := api.parseId(request.PathValue("id"))

	if err != nil {
		http.Error(writer, "Invalid student_id", http.StatusBadRequest)
		return
	}

	reports, err := api.tutorService.GetStudentReportsService(studentID)
	if err != nil || reports == nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("[]"))
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(reports)
}

func (api *TutorAPI) parseId(idStr string) (id int, err error){
    
    id, err = strconv.Atoi(idStr)
    if err != nil {
        fmt.Println("Error parsing ID:", err)
        return 0, err
    }

    return id, nil

}