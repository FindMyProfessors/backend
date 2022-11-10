package database

import (
	"context"
	"github.com/FindMyProfessors/backend/graph/model"
)

func (r *Repository) CreateSchool(ctx context.Context, input *model.NewSchool) (school *model.School, err error) {
	//TODO implement me
	panic("implement me")
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
