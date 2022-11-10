package database

import (
	"context"
	"github.com/FindMyProfessors/backend/graph/model"
)

func (r *Repository) CreateCourse(ctx context.Context, schoolID string, input *model.NewCourse) (course *model.Course, err error) {
	course = &model.Course{
		Name:     input.Name,
		Code:     input.Code,
		SchoolID: schoolID,
	}

	sql := `INSERT INTO courses (name, code, school_id) VALUES ($1, $2, $3) RETURNING id`

	err = r.DatabasePool.QueryRow(ctx, sql, input.Name, input.Code, schoolID).Scan(course.ID)
	if err != nil {
		return nil, err
	}

	return course, err
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
