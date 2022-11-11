package database

import (
	"context"
	"github.com/FindMyProfessors/backend/graph/model"
	"github.com/jackc/pgx/v5"
)

// Adds the course and its attributes to the database with the SQL insert command.
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

// Gets the course with the given id from the database with the SQL select command.
func (r *Repository) GetCourseById(ctx context.Context, id string) (course *model.Course, err error) {
	course = &model.Course{
		ID: id,
	}

	sql := `SELECT name, code, school_id FROM courses  WHERE id = $1`

	err = r.DatabasePool.QueryRow(ctx, sql, course.ID).Scan(course.Name, course.Code, course.SchoolID)
	if err != nil {
		return nil, err
	}

	return course, err
}

func (r *Repository) GetCourseCodesBySchool(ctx context.Context, id string) (courseCodes []*string, err error) {
	courseCodes = []*string{}

	sql := `SELECT code FROM courses WHERE school_id = $1 ORDER BY code DESC`

	rows, err := r.DatabasePool.Query(ctx, sql, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var code string
		err = rows.Scan(&code)
		if err != nil {
			return nil, err
		}
		courseCodes = append(courseCodes, &code)
	}
	if err != nil {
		return nil, err
	}

	return courseCodes, err
}

func (r *Repository) GetCoursesByProfessor(ctx context.Context, id string, first int, after *string) (courses []*model.Course, total int, err error) {
	var sql string
	var variables []any
	if after != nil {
		sql = `SELECT courses.id, courses.name, courses.code, courses.school_id FROM courses INNER JOIN professor_courses pc on courses.id = pc.course_id WHERE professor_id = $1 AND id > $2 ORDER BY courses.id LIMIT $3`
		variables = []any{id, first}
	} else {
		sql = `SELECT courses.id, courses.name, courses.code, courses.school_id FROM courses INNER JOIN professor_courses pc on courses.id = pc.course_id WHERE professor_id = $1 ORDER BY courses.id LIMIT $2`
		variables = []any{id, first}
	}

	err = pgx.BeginTxFunc(ctx, r.DatabasePool, pgx.TxOptions{}, func(tx pgx.Tx) error {
		rows, err := tx.Query(ctx, sql, variables...)
		if err != nil {
			return err
		}

		for rows.Next() {
			var course model.Course
			err = rows.Scan(&course.ID, &course.Name, &course.Code, &course.SchoolID)
			if err != nil {
				return err
			}
			courses = append(courses, &course)
		}

		err = tx.QueryRow(ctx, `SELECT COUNT(*) FROM professor_courses WHERE professor_id = $1`, id).Scan(&total)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	return courses, total, nil
}

func (r *Repository) GetCoursesBySchool(ctx context.Context, id string, first int, after *string) (courses []*model.Course, total int, err error) {
	var sql string
	var variables []any

	if after != nil {
		sql = `SELECT id, name, code FROM courses WHERE school_id = $1 AND id > $2 ORDER BY id LIMIT $3`
		variables = []any{id, *after, first}
	} else {
		sql = `SELECT id, name, code FROM courses WHERE school_id = $1 ORDER BY id LIMIT $2`
		variables = []any{id, first}
	}

	err = pgx.BeginTxFunc(ctx, r.DatabasePool, pgx.TxOptions{}, func(tx pgx.Tx) error {
		rows, err := tx.Query(ctx, sql, variables...)
		if err != nil {
			return err
		}

		for rows.Next() {
			var course model.Course
			err = rows.Scan(&course.ID, &course.Name, &course.Code, &course.SchoolID)
			if err != nil {
				return err
			}
			courses = append(courses, &course)
		}

		err = tx.QueryRow(ctx, `SELECT COUNT(*) FROM courses WHERE school_id = $1`, id).Scan(&total)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, 0, err
	}

	return courses, total, nil
}
