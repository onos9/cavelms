package controller

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/cavelms/internal/model"
)

// CreateReferee is the resolver for the createReferee field.
func (r *mutationResolver) CreateReferee(ctx context.Context, input model.NewReferee) (*model.Referee, error) {
	document, err := r.Service.CreateReferee(ctx, input)
	if err != nil {
		return nil, err
	}

	return document, nil
}

// UpdateReferee is the resolver for the updateReferee field.
func (r *mutationResolver) UpdateReferee(ctx context.Context, input interface{}) (*model.Referee, error) {
	document, err := r.Service.UpdateReferee(ctx, input)
	if err != nil {
		return nil, err
	}

	return document, nil
}

// Referee is the resolver for the referee field.
func (r *queryResolver) Referee(ctx context.Context, id string) (*model.Referee, error) {
	files, err := r.Service.GetRefereeByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return files, nil
}

// Referees is the resolver for the referees field.
func (r *queryResolver) Referees(ctx context.Context, userID *string) ([]*model.Referee, error) {
	files, err := r.Service.GetReferees(ctx, userID)
	if err != nil {
		return nil, err
	}

	return files, nil
}
