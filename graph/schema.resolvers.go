package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/FindMyProfessors/backend/graph/generated"
	"github.com/FindMyProfessors/backend/graph/model"
)

// School is the resolver for the school field.
func (r *courseResolver) School(ctx context.Context, obj *model.Course) (*model.School, error) {
	return r.Repository.GetSchoolByCourse(ctx, obj.ID)
}

// TaughtBy is the resolver for the taughtBy field.
func (r *courseResolver) TaughtBy(ctx context.Context, obj *model.Course, first int, after *string) (*model.ProfessorConnection, error) {
	professors, total, err := r.Repository.GetProfessorsByCourse(ctx, obj.ID, first, after)
	if err != nil {
		return nil, err
	}
	var connection model.ProfessorConnection
	// TODO: Implement Pagination

	return &connection, nil
}

// CreateSchool is the resolver for the createSchool field.
func (r *mutationResolver) CreateSchool(ctx context.Context, input model.NewSchool) (*model.School, error) {
	return r.Repository.CreateSchool(ctx, &input)
}

// CreateProfessor is the resolver for the createProfessor field.
func (r *mutationResolver) CreateProfessor(ctx context.Context, schoolID string, input model.NewProfessor) (*model.Professor, error) {
	return r.Repository.CreateProfessor(ctx, schoolID, &input)
}

// CreateCourse is the resolver for the createCourse field.
func (r *mutationResolver) CreateCourse(ctx context.Context, schoolID string, input model.NewCourse) (*model.Course, error) {
	return r.Repository.CreateCourse(ctx, schoolID, &input)
}

// CreateReview is the resolver for the createReview field.
func (r *mutationResolver) CreateReview(ctx context.Context, professorID string, input model.NewReview) (*model.Review, error) {
	return r.Repository.CreateReview(ctx, professorID, &input)
}

// MergeProfessor is the resolver for the mergeProfessor field.
func (r *mutationResolver) MergeProfessor(ctx context.Context, schoolProfessorID string, rmpProfessorID string) (*model.Professor, error) {
	return r.Repository.MergeProfessor(ctx, schoolProfessorID, rmpProfessorID)
}

// Reviews is the resolver for the reviews field.
func (r *professorResolver) Reviews(ctx context.Context, obj *model.Professor, first int, after *string) (*model.ReviewConnection, error) {
	reviews, total, err := r.Repository.GetReviewsByProfessor(ctx, obj.ID, first, after)
	if err != nil {
		return nil, err
	}
	var connection model.ReviewConnection
	// TODO: Implement Pagination

	return &connection, nil
}

// Teaches is the resolver for the teaches field.
func (r *professorResolver) Teaches(ctx context.Context, obj *model.Professor, first int, after *string) (*model.CourseConnection, error) {
	courses, total, err := r.Repository.GetCoursesByProfessor(ctx, obj.ID, first, after)
	if err != nil {
		return nil, err
	}
	var connection model.CourseConnection
	// TODO: Implement Pagination

	return &connection, nil
}

// Professor is the resolver for the professor field.
func (r *queryResolver) Professor(ctx context.Context, id string) (*model.Professor, error) {
	return r.Repository.GetProfessorById(ctx, id)
}

// School is the resolver for the school field.
func (r *queryResolver) School(ctx context.Context, id string) (*model.School, error) {
	return r.Repository.GetSchoolById(ctx, id)
}

// Schools is the resolver for the schools field.
func (r *queryResolver) Schools(ctx context.Context, first int, after *string) (*model.SchoolConnection, error) {
	schools, total, err := r.Repository.GetSchools(ctx, first, after)
	if err != nil {
		return nil, err
	}
	var connection model.SchoolConnection
	// TODO: Implement Pagination

	return &connection, nil
}

// Professors is the resolver for the professors field.
func (r *queryResolver) Professors(ctx context.Context, schoolID string, first int, after *string) (*model.ProfessorConnection, error) {
	professors, total, err := r.Repository.GetProfessorsBySchool(ctx, schoolID, first, after)
	if err != nil {
		return nil, err
	}
	var connection model.ProfessorConnection
	// TODO: Implement Pagination

	return &connection, nil
}

// CourseCodes is the resolver for the courseCodes field.
func (r *schoolResolver) CourseCodes(ctx context.Context, obj *model.School) ([]*string, error) {
	courseCodes, err := r.Repository.GetCourseCodesBySchool(ctx, obj.ID)
	if err != nil {
		return nil, err
	}
	return courseCodes, nil
}

// Courses is the resolver for the courses field.
func (r *schoolResolver) Courses(ctx context.Context, obj *model.School, first int, after *string) (*model.CourseConnection, error) {
	courses, total, err := r.Repository.GetCoursesBySchool(ctx, obj.ID, first, after)
	if err != nil {
		return nil, err
	}
	var connection model.CourseConnection
	// TODO: Implement Pagination

	return &connection, nil
}

// Professors is the resolver for the professors field.
func (r *schoolResolver) Professors(ctx context.Context, obj *model.School, first int, after *string) (*model.ProfessorConnection, error) {
	professors, total, err := r.Repository.GetProfessorsBySchool(ctx, obj.ID, first, after)
	if err != nil {
		return nil, err
	}
	var connection model.ProfessorConnection
	// TODO: Implement Pagination

	return &connection, nil
}

// Course returns generated.CourseResolver implementation.
func (r *Resolver) Course() generated.CourseResolver { return &courseResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Professor returns generated.ProfessorResolver implementation.
func (r *Resolver) Professor() generated.ProfessorResolver { return &professorResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// School returns generated.SchoolResolver implementation.
func (r *Resolver) School() generated.SchoolResolver { return &schoolResolver{r} }

type courseResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type professorResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type schoolResolver struct{ *Resolver }