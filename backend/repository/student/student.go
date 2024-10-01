package student

import (
	"feedback-tool/model"
)

type StudentRepository interface {
    GetStudents() ([]model.Student, error)
	Close()
}
