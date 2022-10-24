package controller

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/cavelms/internal/model"
)

// CreateRole is the resolver for the createRole field.
func (r *mutationResolver) CreateRole(ctx context.Context, name string) (*model.Role, error) {
	panic(fmt.Errorf("not implemented: CreateRole - createRole"))
}

// UpdateRole is the resolver for the updateRole field.
func (r *mutationResolver) UpdateRole(ctx context.Context, input *model.UpdateRole) (*model.Role, error) {
	panic(fmt.Errorf("not implemented: UpdateRole - updateRole"))
}

// DeleteRole is the resolver for the deleteRole field.
func (r *mutationResolver) DeleteRole(ctx context.Context, id string) (*model.Role, error) {
	panic(fmt.Errorf("not implemented: DeleteRole - deleteRole"))
}

// Role is the resolver for the role field.
func (r *queryResolver) Role(ctx context.Context, id string) (*model.Role, error) {
	panic(fmt.Errorf("not implemented: Role - role"))
}

// Roles is the resolver for the roles field.
func (r *queryResolver) Roles(ctx context.Context) ([]*model.Role, error) {
	panic(fmt.Errorf("not implemented: Roles - roles"))
}

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) Create(ctx context.Context, name string) (*model.Role, error) {
	panic(fmt.Errorf("not implemented: Create - create"))
}
func (r *mutationResolver) Update(ctx context.Context, input *model.UpdateRole) (*model.Role, error) {
	panic(fmt.Errorf("not implemented: Update - update"))
}
func (r *mutationResolver) Delete(ctx context.Context, id string) (*model.Role, error) {
	panic(fmt.Errorf("not implemented: Delete - delete"))
}
