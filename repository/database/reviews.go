package database

import (
	"context"
	"github.com/FindMyProfessors/backend/graph/model"
)

// CreateReview Adds the review and its attributes to the database (associated with a professor_id) with the SQL insert command.
func (r *Repository) CreateReview(ctx context.Context, schoolID string, input *model.NewReview) (review *model.Review, err error) {
	review = &model.Review{
		Quality:    input.Quality,
		Difficulty: input.Difficulty,
		Time:       input.Time,
		Tags:       input.Tags,
		Grade:      input.Grade,
	}
	sql := `INSERT INTO reviews (quality, difficulty, time, tags, grade, professor_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	// TODO: Check if the Tags need to be casted to string
	err = r.DatabasePool.QueryRow(ctx, sql, input.Quality, input.Difficulty, input.Time, input.Tags, input.Grade, schoolID).Scan(&review.ID)
	if err != nil {
		return nil, err
	}
	return review, nil
}

func (r *Repository) GetReviewsByProfessor(ctx context.Context, id string, first int, after *string) (reviews []*model.Review, total int, err error) {
	//TODO implement me
	panic("implement me")
}
