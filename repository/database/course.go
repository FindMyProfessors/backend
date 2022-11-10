package database

import (
	"context"
	"github.com/FindMyProfessors/backend/graph/model"
)

func (r *Repository) CreateCourse(ctx context.Context, schoolID string, input *model.NewCourse) (course *model.Course, err error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) GetCourseById(ctx context.Context, id string) (course *model.Course, err error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) GetCourseCodesBySchool(ctx context.Context, id string) ([]*string, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) GetCoursesByProfessor(ctx context.Context, id string, first int, after *string) (reviews []*model.Course, total int, err error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) GetCoursesBySchool(ctx context.Context, id string, first int, after *string) (courses []*model.Course, total int, err error) {
	//TODO implement me
	panic("implement me")
}
