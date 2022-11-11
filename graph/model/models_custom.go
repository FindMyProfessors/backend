package model

type Course struct {
	ID       string               `json:"id"`
	Name     string               `json:"name"`
	Code     string               `json:"code"`
	SchoolID string               `json:"schoolId"`
	School   *School              `json:"school"`
	TaughtBy *ProfessorConnection `json:"taughtBy"`
}

type Professor struct {
	ID        string             `json:"id"`
	FirstName string             `json:"firstName"`
	LastName  string             `json:"lastName"`
	Rating    *Rating            `json:"rating"`
	Analysis  *ProfessorAnalysis `json:"analysis"`
	SchoolID  string             `json:"schoolId"`
	RMPId     *string            `json:"rmpId"`
	School    *School            `json:"school"`
	Reviews   *ReviewConnection  `json:"reviews"`
	Teaches   *CourseConnection  `json:"teaches"`
}
