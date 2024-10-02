package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/samuelloganbjss/academy-feedback-tool/model"
	"github.com/samuelloganbjss/academy-feedback-tool/service"
)

type StudentAPI struct {
	studentService *service.StudentService
}

func NewStudentAPI(studentService *service.StudentService) *StudentAPI {
	return &StudentAPI{
		studentService: studentService,
	}
}
func (api *StudentAPI) GetStudents(writer http.ResponseWriter, request *http.Request) {
	students, err := api.studentService.GetStudentsService()

	if err != nil {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(writer).Encode(students)
}

func (api *StudentAPI) AddStudent(writer http.ResponseWriter, request *http.Request) {
	var student model.Student
	err := json.NewDecoder(request.Body).Decode(&student)

	if err != nil {
		http.Error(writer, "Bad Request", http.StatusBadRequest)
		return
	}

	createdStudent, err := api.studentService.AddStudentService(student)
	if err != nil {
		http.Error(writer, "Error adding student", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(createdStudent)
}

func (api *StudentAPI) DeleteSingleStudent(writer http.ResponseWriter, request *http.Request) {

	id, err := api.parseId(request.PathValue("id"))

	if err != nil {
		http.Error(writer, "Bad Request ID", http.StatusBadRequest)
		return
	}

	_, err = api.studentService.DeleteStudentService(id)

	if err != nil {
		http.Error(writer, "Could not delete student", http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusOK)

}

func (api *StudentAPI) AddReport(writer http.ResponseWriter, request *http.Request) {
	var report model.Report
	err := json.NewDecoder(request.Body).Decode(&report)

	if err != nil {
		http.Error(writer, "Bad Request", http.StatusBadRequest)
		return
	}

	createdReport, err := api.studentService.AddReportService(report)
	if err != nil {
		http.Error(writer, "Error adding report", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(createdReport)
}

func (api *StudentAPI) EditReport(writer http.ResponseWriter, request *http.Request) {
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

	updatedReport, err := api.studentService.EditReportService(id, newContent.Content, tutorID)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusUnauthorized)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(updatedReport)
}

func (api *StudentAPI) GetStudentReports(writer http.ResponseWriter, request *http.Request) {
	studentID, err := strconv.Atoi(request.URL.Query().Get("student_id"))
	if err != nil {
		http.Error(writer, "Invalid student ID", http.StatusBadRequest)
		return
	}

	reports, err := api.studentService.GetStudentReportsService(studentID)
	if err != nil {
		http.Error(writer, "Error fetching student reports", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(reports)
}

func (api *StudentAPI) parseId(idStr string) (id int, err error) {

	id, err = strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Error parsing ID:", err)
		return 0, err
	}

	return id, nil

}
