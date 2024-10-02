package model

import "time"

type Student struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
}

type Report struct {
	ID        int       `json:"id"`
	StudentID int       `json:"student_id"`
	TutorID   int       `json:"tutor_id"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

type Tutor struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
	Role       string `json:"role"`
}
