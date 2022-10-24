package controller

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/cavelms/internal/model"
)

// CreateCourse is the resolver for the createCourse field.
func (r *mutationResolver) CreateCourse(ctx context.Context, input *model.NewCourse) (*model.User, error) {
	panic(fmt.Errorf("not implemented: CreateCourse - createCourse"))
}

// UpdateCourse is the resolver for the updateCourse field.
func (r *mutationResolver) UpdateCourse(ctx context.Context, data interface{}) (*model.User, error) {
	panic(fmt.Errorf("not implemented: UpdateCourse - updateCourse"))
}

// DeleteCourse is the resolver for the deleteCourse field.
func (r *mutationResolver) DeleteCourse(ctx context.Context, id *string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: DeleteCourse - deleteCourse"))
}

// DeleteManyCourse is the resolver for the deleteManyCourse field.
func (r *mutationResolver) DeleteManyCourse(ctx context.Context, id []string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: DeleteManyCourse - deleteManyCourse"))
}

// Courses is the resolver for the Courses field.
func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	panic(fmt.Errorf("not implemented: Courses - Courses"))
}

// Course is the resolver for the course field.
func (r *queryResolver) Course(ctx context.Context) (*model.Course, error) {
	panic(fmt.Errorf("not implemented: Course - course"))
}
