package tutor

import (
	"github.com/samuelloganbjss/academy-feedback-tool/model"
)

type TutorRepository interface {
	GetTutors() ([]model.Tutor, error)
	AddReport(report model.Report) (model.Report, error)
	EditReport(id int, newContent string, tutorID int) (model.Report, error)
	Close()
}
