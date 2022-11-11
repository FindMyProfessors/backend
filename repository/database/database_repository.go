package database

import (
	"context"
	"github.com/FindMyProfessors/backend/graph/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	DatabasePool *pgxpool.Pool
}

type Queryable interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (commandTag pgconn.CommandTag, err error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

func (r *Repository) RegisterProfessorForCourse(ctx context.Context, courseID string, professorID string, term *model.TermInput) error {
	sql := `INSERT INTO professor_courses (professor_id, course_id, year, semester) VALUES ($1, $2, $3, $3)`

	_, err := r.DatabasePool.Exec(ctx, sql, professorID, courseID, term.Year, term.Semester.String())
	if err != nil {
		return nil
	}
	return err
}
