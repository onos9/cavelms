package controller

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/cavelms/internal/model"
)

// CreateVideo is the resolver for the createVideo field.
func (r *mutationResolver) CreateVideo(ctx context.Context, input model.NewVideo) (*model.Video, error) {
	panic(fmt.Errorf("not implemented: CreateVideo - createVideo"))
}

// Video is the resolver for the video field.
func (r *queryResolver) Video(ctx context.Context, id string) (*model.Video, error) {
	panic(fmt.Errorf("not implemented: Video - video"))
}

// Videos is the resolver for the videos field.
func (r *queryResolver) Videos(ctx context.Context, limit *int, offset *int) ([]*model.Video, error) {
	panic(fmt.Errorf("not implemented: Videos - videos"))
}
