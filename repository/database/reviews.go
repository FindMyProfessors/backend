package database

import (
	"context"
	"github.com/FindMyProfessors/backend/graph/model"
)

// Adds the review and its attributes to the database (associated with a professor_id) with the SQL insert command.
func (r *Repository) CreateReview(ctx context.Context, schoolID string, input *model.NewReview) (course *model.Review, err error) {
	var newReview model.Review
	sql := `INSERT INTO reviews (quality, difficulty, time, tags, grade, professor_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING quality, difficulty, time, tags, grade, professor_id`
	err = r.DatabasePool.QueryRow(ctx, sql, input.Quality, input.Difficulty, input.Time, input.Tags, input.Grade, schoolID).Scan(&newReview.Quality, &newReview.Difficulty, &newReview.Time, &newReview.Tags, &newReview.Grade)
	if err != nil {
		return nil, err
	}
	return &newReview, nil
}

func (r *Repository) GetReviewsByProfessor(ctx context.Context, id string, first int, after *string) (reviews []*model.Review, total int, err error) {
	//TODO implement me
	panic("implement me")
}
