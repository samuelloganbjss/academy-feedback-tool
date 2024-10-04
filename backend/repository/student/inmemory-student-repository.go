package student

import (
	"errors"
	"github.com/samuelloganbjss/academy-feedback-tool/model"
)

var students []model.Student

var studentIDCounter = 4

type InMemoryStudentRepository struct{}

func NewInMemoryStudentRepository() *InMemoryStudentRepository {
	InitDB() // Initialize the in-memory database with sample data
	return &InMemoryStudentRepository{}
}

func InitDB() {
	students = []model.Student{
		{ID: 1, Name: "Alice", Department: "Engineering"},
		{ID: 2, Name: "Bob", Department: "Sparck"},
		{ID: 3, Name: "Bluey", Department: "Consulting"},
	}

}

func (repo *InMemoryStudentRepository) GetStudents() ([]model.Student, error) {
	return students, nil
}

func (repo *InMemoryStudentRepository) AddStudent(student model.Student) (model.Student, error) {
	student.ID = studentIDCounter
	studentIDCounter++
	students = append(students, student)
	return student, nil
}

func (repo *InMemoryStudentRepository) DeleteSingleStudent(id int) (int, error) {
	for i, student := range students {
		if student.ID == id {
			students = append(students[:i], students[i+1:]...)
			return student.ID, nil
		}
	}
	return 0, errors.New("student not found")
}

func (repo *InMemoryStudentRepository) Close() {

}
