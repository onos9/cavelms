package controller

// This education will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/cavelms/internal/model"
)

// CreateEducation is the resolver for the createEducation field.
func (r *mutationResolver) CreateEducation(ctx context.Context, input model.NewEducation) (*model.Education, error) {
	education, err := r.Service.CreateEducation(ctx, input)
	if err != nil {
		return nil, err
	}

	return education, nil
}

// UpdateEducation is the resolver for the updateEducation field.
func (r *mutationResolver) UpdateEducation(ctx context.Context, input interface{}) (*model.Education, error) {
	education, err := r.Service.UpdateEducation(ctx, input)
	if err != nil {
		return nil, err
	}

	return education, nil
}

// Education is the resolver for the education field.
func (r *queryResolver) Education(ctx context.Context, id string) (*model.Education, error) {
	educations, err := r.Service.GetEducationByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return educations, nil
}

// Educations is the resolver for the educations field.
func (r *queryResolver) Educations(ctx context.Context, userID *string) ([]*model.Education, error) {
	educations, err := r.Service.GetEducations(ctx, userID)
	if err != nil {
		return nil, err
	}

	return educations, nil
}
