package service

import (
	"fmt"
	"errors"
	"feedback-tool/model"
	"feedback-tool/repository/student"
)

type StudentService struct {
    repository student.StudentRepository
}

func NewStudentService(repo student.StudentRepository) *StudentService {
    return &StudentService{
        repository: repo,
    }
}

func (s* StudentService) GetStudentsService() ([]model.Student, error) {
	students, err := s.repository.GetStudents()

	if err != nil {
		fmt.Println("Error getting students from DB:", err)
		return nil, errors.New("there was an error getting the students from the database")
	}

	return students, nil
}