package database

import (
	"context"
	"github.com/FindMyProfessors/backend/graph/model"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) CreateProfessor(ctx context.Context, schoolID string, input *model.NewProfessor) (professor *model.Professor, err error) {
	professor = &model.Professor{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		SchoolID:  schoolID,
	}

	sql := `INSERT INTO professors (school_id, first_name, last_name, rmp_id) VALUES ($1, $2, $3, $4) RETURNING id`

	err = r.DatabasePool.QueryRow(ctx, sql, schoolID, input.FirstName, input.LastName, input.RmpID).Scan(professor.ID)
	if err != nil {
		return nil, err
	}

	return professor, err
}

func (r *Repository) GetProfessorsBySchool(ctx context.Context, id string, first int, after *string) (professors []*model.Professor, total int, err error) {
	var sql string
	var variables []any
	if after != nil {
		sql = `SELECT id, first_name, last_name FROM professors WHERE school_id = $1 AND id > $2 ORDER BY id LIMIT $3`
		variables = []any{id, *after, first}
	} else {
		sql = `SELECT id, first_name, last_name FROM professors WHERE school_id = $1 ORDER BY id LIMIT $2`
		variables = []any{id, first}
	}

	err = pgx.BeginTxFunc(ctx, r.DatabasePool, pgx.TxOptions{}, func(tx pgx.Tx) error {
		rows, err := r.DatabasePool.Query(ctx, sql, variables...)
		if err != nil {
			return err
		}

		for rows.Next() {
			professor := model.Professor{SchoolID: id}
			err = rows.Scan(&professor.ID, &professor.FirstName, &professor.LastName)
			if err != nil {
				return err
			}
			professors = append(professors, &professor)
		}

		err = tx.QueryRow(ctx, `SELECT COUNT(*) FROM professors WHERE school_id = $1`, id).Scan(&total)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	return professors, total, err
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
