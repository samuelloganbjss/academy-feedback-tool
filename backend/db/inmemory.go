package db

import (
	"github.com/samuelloganbjss/academy-feedback-tool/model"
)

var students []model.Student

type InMemoryRepository struct{}

func NewInMemoryRepository() *InMemoryRepository {
	InitDB() // Initialize the in-memory database with sample data
	return &InMemoryRepository{}
}

func InitDB() {
	students = []model.Student{
		{ID: 1, Name: "Alice", Department: "Engineering"},
		{ID: 2, Name: "Bob", Department: "Sparck"},
		{ID: 3, Name: "Bluey", Department: "Consulting"},
	}
}

func (repo *InMemoryRepository) GetStudents() ([]model.Student, error) {
	return students, nil
}

func (repo *InMemoryRepository) Close() {

}
