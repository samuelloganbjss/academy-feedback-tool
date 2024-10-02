package db

import (
	"errors"
	"time"

	"github.com/samuelloganbjss/academy-feedback-tool/model"
)

var students []model.Student

var reports []model.Report

var tutors []model.Tutor

var reportIDCounter = 1
var studentIDCounter = 4
var tutorIDCounter = 4

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

	tutors = []model.Tutor{
		{ID: 1, Name: "Peter", Department: "Engineering"},
		{ID: 2, Name: "Clark", Department: "Sparck"},
		{ID: 3, Name: "Lois", Department: "Consulting"},
	}
}

func (repo *InMemoryRepository) GetStudents() ([]model.Student, error) {
	return students, nil
}

func (repo *InMemoryRepository) GetTutors() ([]model.Tutor, error) {
	return tutors, nil
}

func (repo *InMemoryRepository) AddStudent(student model.Student) (model.Student, error) {
	student.ID = studentIDCounter
	studentIDCounter++
	students = append(students, student)
	return student, nil
}

func (repo *InMemoryRepository) AddTutor(tutor model.Tutor) (model.Tutor, error) {
	tutor.ID = tutorIDCounter
	tutorIDCounter++
	tutors = append(tutors, tutor)
	return tutor, nil
}

func (repo *InMemoryRepository) DeleteSingleTutor(id int) (int, error) {
	for i, tutor := range tutors {
		if tutor.ID == id {

			tutors = append(tutors[:i], tutors[i+1:]...)
			return tutor.ID, nil 
		}
	}
	return 0, errors.New("tutor not found")
}

func (repo *InMemoryRepository) DeleteSingleStudent(id int) (int, error) {
	for i, student := range students {
		if student.ID == id {
			students = append(students[:i], students[i+1:]...)
			return student.ID, nil
		}
	}
	return 0, errors.New("student not found")
}

func (repo *InMemoryRepository) AddReport(report model.Report) (model.Report, error) {
	report.ID = reportIDCounter
	reportIDCounter++
	report.Timestamp = time.Now().Format(time.RFC3339)
	reports = append(reports, report)
	return report, nil
}

func (repo *InMemoryRepository) EditReport(id int, newContent string, tutorID int) (model.Report, error) {
	for i, rpt := range reports {
		if rpt.ID == id && rpt.TutorID == tutorID {
			reports[i].Content = newContent
			return reports[i], nil
		}
	}
	return model.Report{}, errors.New("report not found or tutor unauthorized")
}

func (repo *InMemoryRepository) GetReportsByStudent(studentID int) ([]model.Report, error) {
	var studentReports []model.Report
	for _, rpt := range reports {
		if rpt.StudentID == studentID {
			studentReports = append(studentReports, rpt)
		}
	}
	return studentReports, nil
}

func (repo *InMemoryRepository) Close() {

}
