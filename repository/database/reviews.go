package database

import (
	"context"
	"github.com/FindMyProfessors/backend/graph/model"
)

func (r *Repository) CreateReview(ctx context.Context, schoolID string, input *model.NewReview) (course *model.Review, err error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) GetReviewsByProfessor(ctx context.Context, id string, first int, after *string) (reviews []*model.Review, total int, err error) {
	//TODO implement me
	panic("implement me")
}
