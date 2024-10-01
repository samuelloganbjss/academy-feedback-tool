package student

import (
	"github.com/samuelloganbjss/academy-feedback-tool/model"
)

type StudentRepository interface {
	GetStudents() ([]model.Student, error)
	Close()
}
