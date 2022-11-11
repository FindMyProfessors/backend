package database

import (
	"context"
	"github.com/FindMyProfessors/backend/graph/model"
	"strconv"
)

func (r *Repository) CreateSchool(ctx context.Context, input *model.NewSchool) (school *model.School, err error) {
	school = &model.School{
		Name: input.Name,
	}

	sql := `INSERT INTO schools (name) VALUES ($1) RETURNING id`
	var intId int

	err = r.DatabasePool.QueryRow(ctx, sql, input.Name).Scan(&intId)
	if err != nil {
		return nil, err
	}

	school.ID = strconv.Itoa(intId)

	return school, err
}

func (r *Repository) GetSchoolById(ctx context.Context, id string) (school *model.School, err error) {
	school = &model.School{
		ID: id,
	}

	sql := `SELECT name FROM schools WHERE id = $1`

	err = r.DatabasePool.QueryRow(ctx, sql, id).Scan(school.Name)
	if err != nil {
		return nil, err
	}

	return school, err
}

func (r *Repository) GetSchoolByCourse(ctx context.Context, courseId string) (school *model.School, err error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) GetSchools(ctx context.Context, first int, after *string) (schools []*model.School, total int, err error) {
	//TODO implement me
	panic("implement me")
}
