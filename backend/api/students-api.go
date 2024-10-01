package api

import (
	"encoding/json"
	"net/http"

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
