package database

import (
	"context"
	"github.com/FindMyProfessors/backend/graph/model"
)

func (r *Repository) CreateSchool(ctx context.Context, input *model.NewSchool) (school *model.School, err error) {
	school = &model.School{
		Name: input.Name,
	}

	sql := `INSERT INTO schools (name) VALUES ($1) RETURNING id`

	err = r.DatabasePool.QueryRow(ctx, sql, input.Name).Scan(school.ID)
	if err != nil {
		return nil, err
	}

	return school, err
}

func (r *Repository) GetSchoolById(ctx context.Context, id string) (school *model.School, err error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) GetSchoolByCourse(ctx context.Context, courseId string) (school *model.School, err error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) GetSchools(ctx context.Context, first int, after *string) (schools []*model.School, total int, err error) {
	//TODO implement me
	panic("implement me")
}
