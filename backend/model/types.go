package model

type Student struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
}