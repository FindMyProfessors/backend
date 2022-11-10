package database

import (
	"context"
	"github.com/FindMyProfessors/backend/graph/model"
)

func (r *Repository) CreateProfessor(ctx context.Context, schoolID string, input *model.NewProfessor) (professor *model.Professor, err error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) GetProfessorsBySchool(ctx context.Context, id string, first int, after *string) (courses []*model.Professor, total int, err error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) GetProfessorById(ctx context.Context, id string) (professor *model.Professor, err error) {
	professor = &model.Professor{
		ID: id,
	}

	sql := `SELECT first_name, last_name FROM professors WHERE id = $1`

	err = r.DatabasePool.QueryRow(ctx, sql, id).Scan(professor.FirstName, professor.LastName)
	if err != nil {
		return nil, err
	}

	return professor, err
}

func (r *Repository) GetProfessorsByCourse(ctx context.Context, courseId string, first int, after *string) (professors []*model.Professor, total int, err error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) MergeProfessor(ctx context.Context, schoolProfessorID string, rmpProfessorID string) (professor *model.Professor, err error) {
	//TODO implement me
	panic("implement me")
}
