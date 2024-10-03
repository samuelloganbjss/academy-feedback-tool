package student

import (
	"github.com/samuelloganbjss/academy-feedback-tool/model"
)

type StudentRepository interface {
	GetStudents() ([]model.Student, error)
	AddStudent(student model.Student) (model.Student, error)
	DeleteSingleStudent(id int) (int, error)
	Close()
}
