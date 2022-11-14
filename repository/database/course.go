package database

import (
	"context"
	"github.com/FindMyProfessors/backend/graph/model"
	"github.com/jackc/pgx/v5"
	"strconv"
)

// Adds the course and its attributes to the database with the SQL insert command.
func (r *Repository) CreateCourse(ctx context.Context, schoolID string, input *model.NewCourse) (course *model.Course, err error) {
	course = &model.Course{
		Name:     input.Name,
		Code:     input.Code,
		SchoolID: schoolID,
	}

	sql := `INSERT INTO courses (name, code, school_id) VALUES ($1, $2, $3) RETURNING id`

	var intId int
	err = r.DatabasePool.QueryRow(ctx, sql, input.Name, input.Code, schoolID).Scan(&intId)
	if err != nil {
		return nil, err
	}

	course.ID = strconv.Itoa(intId)

	return course, err
}

// Gets the course with the given id from the database with the SQL select command.
func (r *Repository) GetCourseById(ctx context.Context, id string) (course *model.Course, err error) {
	course = &model.Course{
		ID: id,
	}

	sql := `SELECT name, code, school_id FROM courses  WHERE id = $1`

	err = r.DatabasePool.QueryRow(ctx, sql, course.ID).Scan(&course.Name, &course.Code, &course.SchoolID)
	if err != nil {
		return nil, err
	}

	return course, err
}

func (r *Repository) GetCourseCodesBySchool(ctx context.Context, id string, input *model.TermInput) (courseCodes []*string, err error) {
	courseCodes = []*string{}

	sql := `SELECT courses.code FROM courses INNER JOIN professor_courses pc on courses.id = pc.course_id WHERE courses.school_id = $1 AND year = $2 AND semester = $3 ORDER BY id`

	rows, err := r.DatabasePool.Query(ctx, sql, id, input.Year, input.Semester.String())
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

func (r *Repository) GetCoursesByProfessor(ctx context.Context, id string, first int, after *string, input *model.TermInput) (courses []*model.Course, total int, err error) {
	var sql string
	var variables []any
	if after != nil {
		sql = `SELECT courses.id, courses.name, courses.code, courses.school_id FROM courses INNER JOIN professor_courses pc on courses.id = pc.course_id WHERE professor_id = $1 AND year = $2 AND semester = $3 AND id > $4 ORDER BY id LIMIT $5`
		variables = []any{id, input.Year, input.Semester, *after, first}
	} else {
		sql = `SELECT courses.id, courses.name, courses.code, courses.school_id FROM courses INNER JOIN professor_courses pc on courses.id = pc.course_id WHERE professor_id = $1 AND year = $2 AND semester = $3 ORDER BY id LIMIT $4`
		variables = []any{id, input.Year, input.Semester, first}
	}

	err = pgx.BeginTxFunc(ctx, r.DatabasePool, pgx.TxOptions{}, func(tx pgx.Tx) error {
		rows, err := tx.Query(ctx, sql, variables...)
		if err != nil {
			return err
		}

		for rows.Next() {
			var course model.Course
			var intId int
			err = rows.Scan(&intId, &course.Name, &course.Code, &course.SchoolID)
			if err != nil {
				return err
			}
			course.ID = strconv.Itoa(intId)
			courses = append(courses, &course)
		}

		err = tx.QueryRow(ctx, `SELECT COUNT(*) FROM professor_courses WHERE professor_id = $1 AND year = $2 AND semester = $3`, id, input.Year, input.Semester.String()).Scan(&total)
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

func (r *Repository) GetCoursesBySchool(ctx context.Context, id string, first int, after *string, input *model.TermInput) (courses []*model.Course, total int, err error) {
	var sql string
	var variables []any

	if after != nil {
		sql = `SELECT courses.id, courses.name, courses.code FROM courses INNER JOIN professor_courses pc on courses.id = pc.course_id AND pc.semester = $1 AND pc.year = $2 WHERE courses.school_id = $3 AND id > $4 ORDER BY id LIMIT $5`
		variables = []any{input.Semester, input.Year, id, *after, first}
	} else {
		sql = `SELECT courses.id, courses.name, courses.code FROM courses INNER JOIN professor_courses pc on courses.id = pc.course_id AND pc.semester = $1 AND pc.year = $2 WHERE courses.school_id = $3 ORDER BY id LIMIT $4`
		variables = []any{input.Semester, input.Year, id, first}
	}

	err = pgx.BeginTxFunc(ctx, r.DatabasePool, pgx.TxOptions{}, func(tx pgx.Tx) error {
		rows, err := tx.Query(ctx, sql, variables...)
		if err != nil {
			return err
		}

		for rows.Next() {
			var course model.Course
			var intId int
			err = rows.Scan(&intId, &course.Name, &course.Code, &course.SchoolID)
			if err != nil {
				return err
			}
			course.ID = strconv.Itoa(intId)
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
