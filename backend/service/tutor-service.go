package service

import (
	"errors"
	"fmt"
	"github.com/samuelloganbjss/academy-feedback-tool/model"
	"github.com/samuelloganbjss/academy-feedback-tool/repository/tutor"
)

type TutorService struct {
	repository tutor.TutorRepository
}

func NewTutorService(repo tutor.TutorRepository) *TutorService {
	return &TutorService{
		repository: repo,
	}
}

func (s *TutorService) GetTutorsService() ([]model.Tutor, error) {
	tutors, err := s.repository.GetTutors()

	if err != nil {
		fmt.Println("Error getting students from DB:", err)
		return nil, errors.New("there was an error getting the students from the database")
	}

	return tutors, nil
}

func (s *TutorService) DeleteTutorService(id int) (int,error) {
	return s.repository.DeleteSingleTutor(id)
}

func (s *TutorService) AddTutorService(tutor model.Tutor) (model.Tutor, error) {
	return s.repository.AddTutor(tutor)
}

func (s *TutorService) AddReportService(report model.Report) (model.Report, error) {
	return s.repository.AddReport(report)
}

func (s *TutorService) EditReportService(id int, newContent string, tutorID int) (model.Report, error) {
	return s.repository.EditReport(id, newContent, tutorID)
}

func (s *TutorService) GetStudentReportsService(studentID int) ([]model.Report, error) {
	reports, err := s.repository.GetReportsByStudentID(studentID)
	if err != nil {
		return nil, errors.New("error retrieving reports for the student")
	}

	fmt.Printf("Fetched reports for student ID %d: %+v\n", studentID, reports)

	return reports, nil
}