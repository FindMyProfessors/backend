package database

import (
	"context"
	"github.com/FindMyProfessors/backend/graph/model"
	"github.com/jackc/pgx/v5"
	"strconv"
)

const (
	MaxBatchReviewRetrieval = 1000
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
	var intId int
	err = r.DatabasePool.QueryRow(ctx, sql, input.Quality, input.Difficulty, input.Time, input.Tags, input.Grade, schoolID).Scan(&intId)
	if err != nil {
		return nil, err
	}
	review.ID = strconv.Itoa(intId)
	return review, nil
}

func (r *Repository) GetReviewsByProfessor(ctx context.Context, id string, first int, after *string) (reviews []*model.Review, total int, err error) {
	var sql string
	var variables []any
	if first == -1 {
		first = MaxBatchReviewRetrieval
	}
	if after != nil {
		sql = `SELECT reviews.id, reviews.quality, reviews.difficulty, reviews.time, reviews.tags, reviews.grade FROM reviews WHERE professor_id = $1 AND id > $2 ORDER BY id LIMIT $3`
		variables = []any{id, *after, first}
	} else {
		sql = `SELECT reviews.id, reviews.quality, reviews.difficulty, reviews.time, reviews.tags, reviews.grade FROM reviews WHERE professor_id = $1 ORDER BY id LIMIT $2`
		variables = []any{id, first}
	}

	err = pgx.BeginTxFunc(ctx, r.DatabasePool, pgx.TxOptions{}, func(tx pgx.Tx) error {
		rows, err := tx.Query(ctx, sql, variables...)
		if err != nil {
			return err
		}

		for rows.Next() {
			var review model.Review
			var intId int
			err = rows.Scan(&review.ID, &review.Quality, &review.Difficulty, &review.Time, &review.Tags, &review.Grade)
			if err != nil {
				return err
			}
			review.ID = strconv.Itoa(intId)
			reviews = append(reviews, &review)
		}

		err = tx.QueryRow(ctx, `SELECT COUNT(*) FROM reviews WHERE professor_id = $1`, id).Scan(&total)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	return reviews, total, nil
}
