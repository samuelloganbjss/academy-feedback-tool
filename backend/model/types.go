package model

type Student struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
}

type Report struct {
	ID        int    `json:"id"`
	StudentID int    `json:"student_id"`
	TutorID   int    `json:"tutor_id"`
	Content   string `json:"content"`
	Timestamp string `json:"timestamp"`
}

type Tutor struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
	Role       string `json:"role"`
}
