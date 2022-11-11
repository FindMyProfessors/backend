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

// GradeIndex Transforms a grade into the index in the AllGrade array defined in models_gen.go
func (e Grade) GradeIndex() int {
	switch e {
	case GradeAPlus:
		// not real just used to distinguish betwen A+ and A
		return 0
	case GradeA:
		return 1
	case GradeAMinus:
		return 2

	case GradeBPlus:
		return 3
	case GradeB:
		return 4
	case GradeBMinus:
		return 5

	case GradeCPlus:
		return 6
	case GradeC:
		return 7
	case GradeCMinus:
		return 8

	case GradeDPlus:
		return 9
	case GradeD:
		return 10
	case GradeDMinus:
		return 11

	case GradeFPlus:
		return 12
	case GradeF:
		return 13
	case GradeFMinus:
		return 14

	default:
		return -1
	}
}
