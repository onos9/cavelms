package controller

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/cavelms/internal/model"
)

// CreateQuiz is the resolver for the createQuiz field.
func (r *mutationResolver) CreateQuiz(ctx context.Context, input model.NewQuiz) (*model.Quiz, error) {
	panic(fmt.Errorf("not implemented: CreateQuiz - createQuiz"))
}

// Quiz is the resolver for the quiz field.
func (r *queryResolver) Quiz(ctx context.Context, id string) (*model.Quiz, error) {
	panic(fmt.Errorf("not implemented: Quiz - quiz"))
}

// Quizzes is the resolver for the quizzes field.
func (r *queryResolver) Quizzes(ctx context.Context, limit *int, offset *int) ([]*model.Quiz, error) {
	panic(fmt.Errorf("not implemented: Quizzes - quizzes"))
}
