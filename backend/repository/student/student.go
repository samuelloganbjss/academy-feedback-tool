package student

import (
	"github.com/samuelloganbjss/academy-feedback-tool/model"
)

type StudentRepository interface {
	GetStudents() ([]model.Student, error)
	AddReport(report model.Report) (model.Report, error)
	EditReport(id int, newContent string, tutorID int) (model.Report, error)

	Close()
}
