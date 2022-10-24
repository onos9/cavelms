package controller

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/cavelms/internal/model"
)

// CreateSetting is the resolver for the createSetting field.
func (r *mutationResolver) CreateSetting(ctx context.Context, input model.NewSetting) (*model.Video, error) {
	panic(fmt.Errorf("not implemented: CreateSetting - createSetting"))
}

// Setting is the resolver for the setting field.
func (r *queryResolver) Setting(ctx context.Context, id string) (*model.Role, error) {
	panic(fmt.Errorf("not implemented: Setting - setting"))
}

// Settings is the resolver for the settings field.
func (r *queryResolver) Settings(ctx context.Context, limit *int, offset *int) ([]*model.Setting, error) {
	panic(fmt.Errorf("not implemented: Settings - settings"))
}
