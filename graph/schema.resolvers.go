package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"math"

	"github.com/FindMyProfessors/backend/analysis"
	"github.com/FindMyProfessors/backend/graph/generated"
	"github.com/FindMyProfessors/backend/graph/model"
	"github.com/FindMyProfessors/backend/pagination"
)

// School is the resolver for the school field.
func (r *courseResolver) School(ctx context.Context, obj *model.Course) (*model.School, error) {
	return r.Repository.GetSchoolByCourse(ctx, obj.ID)
}

// TaughtBy is the resolver for the taughtBy field.
func (r *courseResolver) TaughtBy(ctx context.Context, obj *model.Course, term model.TermInput, first int, after *string) (*model.ProfessorConnection, error) {
	if after != nil {
		cursor, err := pagination.DecodeCursor(after)
		if err != nil {
			return nil, err
		}
		after = &cursor
	}
	professors, total, err := r.Repository.GetProfessorsByCourse(ctx, obj.ID, first, after, &term)
	if err != nil {
		return nil, err
	}
	if total == 0 {
		return &model.ProfessorConnection{TotalCount: 0, PageInfo: &model.PageInfo{}}, nil
	}

	return &model.ProfessorConnection{TotalCount: total, PageInfo: pagination.GetPageInfo(professors[0].ID, professors[len(professors)-1].ID, len(professors), first), Professors: professors}, nil
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

// RegisterProfessorForCourse is the resolver for the registerProfessorForCourse field.
func (r *mutationResolver) RegisterProfessorForCourse(ctx context.Context, courseID string, professorID string, term model.TermInput) (bool, error) {
	err := r.Repository.RegisterProfessorForCourse(ctx, courseID, professorID, &term)

	return err != nil, err
}

// MergeProfessor is the resolver for the mergeProfessor field.
func (r *mutationResolver) MergeProfessor(ctx context.Context, schoolProfessorID string, rmpProfessorID string, input model.NewProfessor) (*model.Professor, error) {
	return r.Repository.MergeProfessor(ctx, schoolProfessorID, rmpProfessorID, &input)
}

// Linked is the resolver for the linked field.
func (r *professorResolver) Linked(ctx context.Context, obj *model.Professor) (bool, error) {
	return obj.RMPId != nil, nil
}

// Rating is the resolver for the rating field.
func (r *professorResolver) Rating(ctx context.Context, obj *model.Professor, topKpercentage *float64) (*model.Rating, error) {
	// TODO: Implement cache
	var err error

	reviews, total, err := r.Repository.GetReviewsByProfessor(ctx, obj.ID, -1, nil)
	if err != nil {
		return nil, err
	}
	m := &model.Rating{
		RatingAmount: total,
	}

	var topKTotal int

	if topKpercentage != nil {
		if *topKpercentage > 1.0 {
			return nil, errors.New("topKpercentage must be in (0, 1]")
		}

		if err != nil {
			return nil, err
		}
		topKTotal = int(float64(total) * (1 - *topKpercentage))
	}

	if total == 0 {
		return nil, nil
	}

	totalQualitySum := 0.0
	totalDifficultySum := 0.0

	topKQualitySum := 0.0
	topKDifficultySum := 0.0

	// TODO: check where the -1 needs to be
	topKStartIndex := int(math.Abs(float64(topKTotal - total - 1)))

	totalGrades := 0
	gradeIndexSum := 0

	for i, review := range reviews {
		totalQualitySum += review.Quality
		totalDifficultySum += review.Difficulty

		if topKpercentage != nil {
			if i > topKStartIndex {
				topKQualitySum += review.Quality
				topKDifficultySum += review.Difficulty
			}
		}

		index := review.Grade.GradeIndex()
		if index != -1 {
			gradeIndexSum += index
			totalGrades++
		}
	}

	if gradeIndexSum == 0 || totalGrades == 0 {
		m.AverageGrade = model.GradeOther
	} else {
		m.AverageGrade = model.AllGrade[gradeIndexSum/totalGrades]
	}

	m.TopKMostRecentDifficultyAverage = topKDifficultySum / float64(topKTotal)
	m.TopKMostRecentQualityAverage = topKQualitySum / float64(topKTotal)

	m.TotalQualityAverage = totalQualitySum / float64(total)
	m.TotalDifficultyAverage = totalDifficultySum / float64(total)

	return m, nil
}

// Analysis is the resolver for the analysis field.
func (r *professorResolver) Analysis(ctx context.Context, obj *model.Professor) (*model.ProfessorAnalysis, error) {
	reviews, _, err := r.Repository.GetReviewsByProfessor(ctx, obj.ID, -1, nil)
	if err != nil {
		return nil, err
	}

	return analysis.BeginAnalysis(reviews)
}

// School is the resolver for the school field.
func (r *professorResolver) School(ctx context.Context, obj *model.Professor) (*model.School, error) {
	if len(obj.SchoolID) > 0 {
		return r.Repository.GetSchoolById(ctx, obj.SchoolID)
	}
	return r.Repository.GetSchoolByProfessor(ctx, obj.ID)
}

// Reviews is the resolver for the reviews field.
func (r *professorResolver) Reviews(ctx context.Context, obj *model.Professor, first int, after *string) (*model.ReviewConnection, error) {
	if after != nil {
		cursor, err := pagination.DecodeCursor(after)
		if err != nil {
			return nil, err
		}
		after = &cursor
	}
	reviews, total, err := r.Repository.GetReviewsByProfessor(ctx, obj.ID, first, after)
	if err != nil {
		return nil, err
	}
	if total == 0 {
		return &model.ReviewConnection{TotalCount: 0, PageInfo: &model.PageInfo{}}, nil
	}

	return &model.ReviewConnection{TotalCount: total, PageInfo: pagination.GetPageInfo(reviews[0].ID, reviews[len(reviews)-1].ID, len(reviews), first), Reviews: reviews}, nil
}

// Teaches is the resolver for the teaches field.
func (r *professorResolver) Teaches(ctx context.Context, obj *model.Professor, term model.TermInput, first int, after *string) (*model.CourseConnection, error) {
	if after != nil {
		cursor, err := pagination.DecodeCursor(after)
		if err != nil {
			return nil, err
		}
		after = &cursor
	}
	courses, total, err := r.Repository.GetCoursesByProfessor(ctx, obj.ID, first, after, &term)
	if err != nil {
		return nil, err
	}
	if total == 0 {
		return &model.CourseConnection{TotalCount: 0, PageInfo: &model.PageInfo{}}, nil
	}

	return &model.CourseConnection{TotalCount: total, PageInfo: pagination.GetPageInfo(courses[0].ID, courses[len(courses)-1].ID, len(courses), first), Courses: courses}, nil
}

// ProfessorByRMPId is the resolver for the professorByRMPId field.
func (r *queryResolver) ProfessorByRMPId(ctx context.Context, rmpID string) (*model.Professor, error) {
	return r.Repository.GetProfessorByRMPId(ctx, rmpID)
}

// Professor is the resolver for the professor field.
func (r *queryResolver) Professor(ctx context.Context, id string) (*model.Professor, error) {
	return r.Repository.GetProfessorById(ctx, id)
}

// Course is the resolver for the course field.
func (r *queryResolver) Course(ctx context.Context, id string) (*model.Course, error) {
	return r.Repository.GetCourseById(ctx, id)
}

// School is the resolver for the school field.
func (r *queryResolver) School(ctx context.Context, id string) (*model.School, error) {
	return r.Repository.GetSchoolById(ctx, id)
}

// Schools is the resolver for the schools field.
func (r *queryResolver) Schools(ctx context.Context, first int, after *string) (*model.SchoolConnection, error) {
	if after != nil {
		cursor, err := pagination.DecodeCursor(after)
		if err != nil {
			return nil, err
		}
		after = &cursor
	}
	schools, total, err := r.Repository.GetSchools(ctx, first, after)
	if err != nil {
		return nil, err
	}
	if total == 0 {
		return &model.SchoolConnection{TotalCount: 0, PageInfo: &model.PageInfo{}}, nil
	}

	return &model.SchoolConnection{TotalCount: total, PageInfo: pagination.GetPageInfo(schools[0].ID, schools[len(schools)-1].ID, len(schools), first), Schools: schools}, nil
}

// Professors is the resolver for the professors field.
func (r *queryResolver) Professors(ctx context.Context, schoolID string, first int, after *string) (*model.ProfessorConnection, error) {
	if after != nil {
		cursor, err := pagination.DecodeCursor(after)
		if err != nil {
			return nil, err
		}
		after = &cursor
	}
	professors, total, err := r.Repository.GetProfessorsBySchool(ctx, schoolID, first, after)
	if err != nil {
		return nil, err
	}
	if total == 0 {
		return &model.ProfessorConnection{TotalCount: 0, PageInfo: &model.PageInfo{}}, nil
	}

	return &model.ProfessorConnection{TotalCount: total, PageInfo: pagination.GetPageInfo(professors[0].ID, professors[len(professors)-1].ID, len(professors), first), Professors: professors}, nil
}

// CourseCodes is the resolver for the courseCodes field.
func (r *schoolResolver) CourseCodes(ctx context.Context, obj *model.School, term model.TermInput) ([]*string, error) {
	courseCodes, err := r.Repository.GetCourseCodesBySchool(ctx, obj.ID, &term)
	if err != nil {
		return nil, err
	}
	return courseCodes, nil
}

// Courses is the resolver for the courses field.
func (r *schoolResolver) Courses(ctx context.Context, obj *model.School, term model.TermInput, filter *model.CourseFilter, first int, after *string) (*model.CourseConnection, error) {
	if after != nil {
		cursor, err := pagination.DecodeCursor(after)
		if err != nil {
			return nil, err
		}
		after = &cursor
	}
	courses, total, err := r.Repository.GetCoursesBySchool(ctx, obj.ID, first, after, &term, filter)
	if err != nil {
		return nil, err
	}
	if total == 0 {
		return &model.CourseConnection{TotalCount: 0, PageInfo: &model.PageInfo{}}, nil
	}

	return &model.CourseConnection{TotalCount: total, PageInfo: pagination.GetPageInfo(courses[0].ID, courses[len(courses)-1].ID, len(courses), first), Courses: courses}, nil
}

// Professors is the resolver for the professors field.
func (r *schoolResolver) Professors(ctx context.Context, obj *model.School, first int, after *string) (*model.ProfessorConnection, error) {
	if after != nil {
		cursor, err := pagination.DecodeCursor(after)
		if err != nil {
			return nil, err
		}
		after = &cursor
	}
	professors, total, err := r.Repository.GetProfessorsBySchool(ctx, obj.ID, first, after)
	if err != nil {
		return nil, err
	}
	if total == 0 {
		return &model.ProfessorConnection{TotalCount: 0, PageInfo: &model.PageInfo{}}, nil
	}

	return &model.ProfessorConnection{TotalCount: total, PageInfo: pagination.GetPageInfo(professors[0].ID, professors[len(professors)-1].ID, len(professors), first), Professors: professors}, nil
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
