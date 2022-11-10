package database

import (
	"context"
	"github.com/FindMyProfessors/backend/graph/model"
)

// Adds the course and its attributes to the database with the SQL insert command.
func (r *Repository) CreateCourse(ctx context.Context, schoolID string, input *model.NewCourse) (course *model.Course, err error) {
	var newCourse model.Course
	sql := `INSERT INTO courses (name, code, school_id) VALUES ($1, $2, $3) RETURNING name, code, school_id`
	err = r.DatabasePool.QueryRow(ctx, sql, input.Name, input.Code, schoolID).Scan(&newCourse.Name, &newCourse.Code, &newCourse.School)
	if err != nil {
		return nil, err
	}
	return &newCourse, nil
}

// Gets the course with the given id from the database with the SQL select command.
func (r *Repository) GetCourseById(ctx context.Context, id string) (course *model.Course, err error) {
	var course1 model.Course
	sql := `SELECT name, code, school_id FROM courses WHERE id = $1`
	err = r.DatabasePool.QueryRow(ctx, sql, id).Scan(&course1.Name, &course1.Code, &course1.School	)
	if err != nil {
		return nil, err
	}
	return &course1, nil
}

// Gets a list of course ids by school string from the database with the SQL select command.
func (r *Repository) GetCourseCodesBySchool(ctx context.Context, id string) ([]*string, error) {
	var courseCodes []*string
	sql := `SELECT code FROM courses WHERE school = $1`
	rows, err := r.DatabasePool.Query(ctx, sql, id)
	if err != nil {
		return nil, err
	}
	// Iterate through the rows and append the course codes to the list.
	for rows.Next() {
		var course model.Course
		err = rows.Scan(&course.Code)
		if err != nil {
			return nil, err
		}
		courseCodes = append(courseCodes, &course.Code)
	}
	return courseCodes, nil
}

func (r *Repository) GetCoursesByProfessor(ctx context.Context, id string, first int, after *string) (reviews []*model.Course, total int, err error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) GetCoursesBySchool(ctx context.Context, id string, first int, after *string) (courses []*model.Course, total int, err error) {
	//TODO implement me
	panic("implement me")
}