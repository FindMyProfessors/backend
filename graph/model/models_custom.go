package model

type Course struct {
	ID       string               `json:"id"`
	Name     string               `json:"name"`
	Code     string               `json:"code"`
	SchoolID string               `json:"school_id"`
	School   *School              `json:"school"`
	TaughtBy *ProfessorConnection `json:"taughtBy"`
}
