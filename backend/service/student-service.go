package service

import (
	"errors"
	"fmt"

	"github.com/samuelloganbjss/academy-feedback-tool/model"
	"github.com/samuelloganbjss/academy-feedback-tool/repository/student"
)

type StudentService struct {
	repository student.StudentRepository
}

func NewStudentService(repo student.StudentRepository) *StudentService {
	return &StudentService{
		repository: repo,
	}
}

func (s *StudentService) GetStudentsService() ([]model.Student, error) {
	students, err := s.repository.GetStudents()

	if err != nil {
		fmt.Println("Error getting students from DB:", err)
		return nil, errors.New("there was an error getting the students from the database")
	}

	return students, nil
}

func (s *StudentService) AddReportService(report model.Report) (model.Report, error) {
	return s.repository.AddReport(report)
}

func (s *StudentService) EditReportService(id int, newContent string, tutorID int) (model.Report, error) {
	return s.repository.EditReport(id, newContent, tutorID)
}
