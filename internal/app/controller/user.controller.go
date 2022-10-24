package controller

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/cavelms/internal/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, data interface{}) (*model.User, error) {
	users, err := r.Service.UpdateUser(ctx, data)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id *string) (*model.User, error) {
	return &model.User{}, nil
}

// DeleteManyUsers is the resolver for the deleteManyUsers field.
func (r *mutationResolver) DeleteManyUsers(ctx context.Context, id []string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: DeleteManyUsers - deleteManyUsers"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	users, err := r.Service.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	user, err := r.Service.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
