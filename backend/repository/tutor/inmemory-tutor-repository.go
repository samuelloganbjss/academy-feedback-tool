package tutor

import (
	"errors"
	"sort"
	"time"
	"github.com/samuelloganbjss/academy-feedback-tool/model"
)

var reports []model.Report

var tutors []model.Tutor

var reportIDCounter = 1
var tutorIDCounter = 4

type InMemoryTutorRepository struct{}

func NewInMemoryTutorRepository() *InMemoryTutorRepository {
	InitDB() // Initialize the in-memory database with sample data
	return &InMemoryTutorRepository{}
}

func InitDB() {
	tutors = []model.Tutor{
		{ID: 1, Name: "Peter", Department: "Engineering"},
		{ID: 2, Name: "Clark", Department: "Sparck"},
		{ID: 3, Name: "Lois", Department: "Consulting"},
	}

}

func (repo *InMemoryTutorRepository) GetTutors() ([]model.Tutor, error) {
	return tutors, nil
}

func (repo *InMemoryTutorRepository) AddTutor(tutor model.Tutor) (model.Tutor, error) {
	tutor.ID = tutorIDCounter
	tutorIDCounter++
	tutors = append(tutors, tutor)
	return tutor, nil
}

func (repo *InMemoryTutorRepository) DeleteSingleTutor(id int) (int, error) {
	for i, tutor := range tutors {
		if tutor.ID == id {

			tutors = append(tutors[:i], tutors[i+1:]...)
			return tutor.ID, nil
		}
	}
	return 0, errors.New("tutor not found")
}

func (repo *InMemoryTutorRepository) AddReport(report model.Report) (model.Report, error) {
	report.ID = reportIDCounter
	reportIDCounter++
	report.Timestamp = time.Now()
	reports = append(reports, report)
	return report, nil
}

func (repo *InMemoryTutorRepository) EditReport(id int, newContent string, tutorID int) (model.Report, error) {
	for i, rpt := range reports {
		if rpt.ID == id && rpt.TutorID == tutorID {
			reports[i].Content = newContent
			return reports[i], nil
		}
	}
	return model.Report{}, errors.New("report not found or tutor unauthorized")
}

func (repo *InMemoryTutorRepository) GetReportsByStudentID(studentID int) ([]model.Report, error) {
	var studentReports []model.Report
	for _, report := range reports {
		if report.StudentID == studentID {
			studentReports = append(studentReports, report)
		}
	}

	// Sort reports by Timestamp
	sort.Slice(studentReports, func(i, j int) bool {
		return studentReports[i].Timestamp.Before(studentReports[j].Timestamp)
	})

	return studentReports, nil
}

func (repo *InMemoryTutorRepository) Close() {

}
