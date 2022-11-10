package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/FindMyProfessors/backend/graph/generated"
	"github.com/FindMyProfessors/backend/graph/model"
)

// CreateSchool is the resolver for the createSchool field.
func (r *mutationResolver) CreateSchool(ctx context.Context, input model.NewSchool) (*model.School, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreateProfessor is the resolver for the createProfessor field.
func (r *mutationResolver) CreateProfessor(ctx context.Context, schoolID string, input model.NewProfessor) (*model.Professor, error) {
	panic(fmt.Errorf("not implemented"))
}

// Reviews is the resolver for the reviews field.
func (r *professorResolver) Reviews(ctx context.Context, obj *model.Professor, after *string, first *int) (*model.ReviewConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Professor is the resolver for the professor field.
func (r *queryResolver) Professor(ctx context.Context, id string) (*model.Professor, error) {
	panic(fmt.Errorf("not implemented"))
}

// School is the resolver for the school field.
func (r *queryResolver) School(ctx context.Context, id string) (*model.School, error) {
	panic(fmt.Errorf("not implemented"))
}

// Schools is the resolver for the schools field.
func (r *queryResolver) Schools(ctx context.Context, after *string, first *int) (*model.SchoolConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Professors is the resolver for the professors field.
func (r *queryResolver) Professors(ctx context.Context, schoolID string, after *string, first *int) (*model.ProfessorConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Professors is the resolver for the professors field.
func (r *schoolResolver) Professors(ctx context.Context, obj *model.School, courseCode string, after *string, first *int) (*model.ProfessorConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Professor returns generated.ProfessorResolver implementation.
func (r *Resolver) Professor() generated.ProfessorResolver { return &professorResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// School returns generated.SchoolResolver implementation.
func (r *Resolver) School() generated.SchoolResolver { return &schoolResolver{r} }

type mutationResolver struct{ *Resolver }
type professorResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type schoolResolver struct{ *Resolver }
