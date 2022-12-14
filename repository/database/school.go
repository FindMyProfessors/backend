package database

import (
	"context"
	"github.com/FindMyProfessors/backend/graph/model"
	"github.com/jackc/pgx/v5"
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

	err = r.DatabasePool.QueryRow(ctx, sql, id).Scan(&school.Name)
	if err != nil {
		return nil, err
	}

	return school, err
}

func (r *Repository) GetSchoolByCourse(ctx context.Context, courseId string) (school *model.School, err error) {
	course := &model.Course{
		ID: courseId,
	}
	err = pgx.BeginTxFunc(ctx, r.DatabasePool, pgx.TxOptions{}, func(tx pgx.Tx) error {
		sql := `SELECT school_id FROM courses WHERE id = $1`
		err = tx.QueryRow(ctx, sql, courseId).Scan(&course.SchoolID)
		if err != nil {
			return err
		}

		school = &model.School{
			ID: course.SchoolID,
		}
		sql = `SELECT name FROM schools WHERE id = $1`
		err = tx.QueryRow(ctx, sql, course.SchoolID).Scan(&school.Name)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return school, err
}

func (r *Repository) GetSchools(ctx context.Context, first int, after *string) (schools []*model.School, total int, err error) {
	var sql string
	var variables []any
	if after != nil && len(*after) > 0 {
		sql = `SELECT id, name FROM schools WHERE id > $1 ORDER BY id LIMIT $2`
		variables = []any{*after, first}
	} else {
		sql = `SELECT id, name FROM schools ORDER BY id LIMIT $1`
		variables = []any{first}
	}

	err = pgx.BeginTxFunc(ctx, r.DatabasePool, pgx.TxOptions{}, func(tx pgx.Tx) error {
		rows, err := tx.Query(ctx, sql, variables...)
		if err != nil {
			return err
		}

		for rows.Next() {
			var school model.School
			var intId int
			if err = rows.Scan(&intId, &school.Name); err != nil {
				return err
			}
			school.ID = strconv.Itoa(intId)
			schools = append(schools, &school)
		}

		if err = tx.QueryRow(ctx, `SELECT COUNT(*) FROM schools`).Scan(&total); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, 0, err
	}
	return schools, total, nil
}

func (r *Repository) GetSchoolByProfessor(ctx context.Context, id string) (school *model.School, err error) {
	sql := `SELECT schools.id, schools.name FROM schools JOIN professors p on schools.id = p.school_id WHERE p.id = $1`

	school = &model.School{}

	err = r.DatabasePool.QueryRow(ctx, sql, id).Scan(&school.ID, &school.Name)
	if err != nil {
		return nil, err
	}
	return school, nil
}
