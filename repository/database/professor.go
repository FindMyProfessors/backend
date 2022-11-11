package database

import (
	"context"
	"errors"
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
		sql = `SELECT id, first_name, last_name, rmp_id FROM professors WHERE school_id = $1 AND id > $2 ORDER BY id LIMIT $3`
		variables = []any{id, *after, first}
	} else {
		sql = `SELECT id, first_name, last_name, rmp_id FROM professors WHERE school_id = $1 ORDER BY id LIMIT $2`
		variables = []any{id, first}
	}

	err = pgx.BeginTxFunc(ctx, r.DatabasePool, pgx.TxOptions{}, func(tx pgx.Tx) error {
		rows, err := r.DatabasePool.Query(ctx, sql, variables...)
		if err != nil {
			return err
		}

		for rows.Next() {
			professor := model.Professor{SchoolID: id}
			err = rows.Scan(&professor.ID, &professor.FirstName, &professor.LastName, professor.RMPId)
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
	return GetProfessorByIdWithQueryable(ctx, r.DatabasePool, id)
}

func GetProfessorByIdWithQueryable(ctx context.Context, queryable Queryable, id string) (professor *model.Professor, err error) {
	professor = &model.Professor{
		ID: id,
	}

	sql := `SELECT first_name, last_name, rmp_id FROM professors WHERE id = $1`

	err = queryable.QueryRow(ctx, sql, id).Scan(&professor.FirstName, &professor.LastName, professor.RMPId)
	if err != nil {
		return nil, err
	}

	return professor, err
}

func (r *Repository) GetProfessorsByCourse(ctx context.Context, courseId string, first int, after *string) (professors []*model.Professor, total int, err error) {
	var sql string
	var variables []any
	if after != nil {
		sql = `SELECT professors.id, professors.first_name, professors.last_name, professors.school_id FROM professors INNER JOIN professor_courses pc on professors.id = pc.professor_id WHERE course_id = $1 AND id > $2 ORDER BY id LIMIT $3`
		variables = []any{courseId, *after, first}
	} else {
		sql = `SELECT professors.id, professors.first_name, professors.last_name, professors.school_id FROM professors INNER JOIN professor_courses pc on professors.id = pc.professor_id WHERE course_id = $1 ORDER BY id LIMIT $2`
		variables = []any{courseId, first}
	}

	err = pgx.BeginTxFunc(ctx, r.DatabasePool, pgx.TxOptions{}, func(tx pgx.Tx) error {
		rows, err := r.DatabasePool.Query(ctx, sql, variables...)
		if err != nil {
			return err
		}

		for rows.Next() {
			var professor model.Professor
			err = rows.Scan(&professor.ID, &professor.FirstName, &professor.LastName, &professor.SchoolID)
			if err != nil {
				return err
			}
			professors = append(professors, &professor)
		}

		err = tx.QueryRow(ctx, `SELECT COUNT(*) FROM professor_courses WHERE course_id = $1`, courseId).Scan(&total)
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

func (r *Repository) MergeProfessor(ctx context.Context, schoolProfessorID string, rmpProfessorID string, input *model.NewProfessor) (professor *model.Professor, err error) {
	// rmp -> school
	err = pgx.BeginTxFunc(ctx, r.DatabasePool, pgx.TxOptions{}, func(tx pgx.Tx) error {
		// Safeguards
		schoolProfessor, err := GetProfessorByIdWithQueryable(ctx, tx, schoolProfessorID)
		if schoolProfessor.RMPId != nil && len(*schoolProfessor.RMPId) > 0 {
			return errors.New("the school professor provided is already linked to an rmp professor")
		}

		// Merge
		// update rmp reviews to new professor
		updateReviewsSql := `UPDATE reviews SET professor_id = $1 WHERE professor_id = $2 `

		_, err = tx.Exec(ctx, updateReviewsSql, schoolProfessorID, rmpProfessorID)
		if err != nil {
			return err
		}

		// delete rmp professor
		deleteRmpProfessor := `DELETE FROM professors WHERE id = $1`
		_, err = tx.Exec(ctx, deleteRmpProfessor, rmpProfessorID)
		if err != nil {
			return err
		}

		if input != nil {
			// update name of new professor
			updateProfessorNameSql := `UPDATE professors set (first_name, last_name) = ($1, $2) WHERE id = $3`
			_, err = tx.Exec(ctx, updateProfessorNameSql, schoolProfessorID)
			if err != nil {
				return err
			}
		}

		return nil
	})
	return professor, err
}
