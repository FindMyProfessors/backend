package repository

import (
	"context"
	"github.com/FindMyProfessors/backend/graph/model"
)

type Repository interface {
	// Create Functions
	CreateSchool(ctx context.Context, input *model.NewSchool) (school *model.School, err error)
	CreateProfessor(ctx context.Context, schoolID string, input *model.NewProfessor) (professor *model.Professor, err error)
	CreateCourse(ctx context.Context, schoolID string, input *model.NewCourse) (course *model.Course, err error)
	CreateReview(ctx context.Context, schoolID string, input *model.NewReview) (course *model.Review, err error)

	// School Type
	GetCourseCodesBySchool(ctx context.Context, id string) (courseCodes []*string, err error)
	GetCoursesBySchool(ctx context.Context, id string, first int, after *string) (courses []*model.Course, total int, err error)
	GetProfessorsBySchool(ctx context.Context, id string, first int, after *string) (professors []*model.Professor, total int, err error)
	GetSchoolById(ctx context.Context, id string) (school *model.School, err error)

	// Professor Type
	GetReviewsByProfessor(ctx context.Context, id string, first int, after *string) (reviews []*model.Review, total int, err error)
	GetCoursesByProfessor(ctx context.Context, id string, first int, after *string) (courses []*model.Course, total int, err error)
	GetProfessorById(ctx context.Context, id string) (professor *model.Professor, err error)

	// Course Type
	GetSchoolByCourse(ctx context.Context, courseId string) (school *model.School, err error)
	GetProfessorsByCourse(ctx context.Context, obj *model.Course, first int, after *string) (professors []*model.Professor, total int, err error)
	GetCourseById(ctx context.Context, id string) (course *model.Course, err error)

	// Edge Case Mutations
	MergeProfessor(ctx context.Context, schoolProfessorID string, rmpProfessorID string) (professor *model.Professor, err error)

	GetSchools(ctx context.Context, first int, after *string) (schools []*model.School, total int, err error)
}
