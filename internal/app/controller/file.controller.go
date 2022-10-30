package controller

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/cavelms/internal/model"
)

// CreateFile is the resolver for the createFile field.
func (r *mutationResolver) CreateFile(ctx context.Context, input model.NewFile) (*model.File, error) {
	document, err := r.Service.CreateFile(ctx, input)
	if err != nil {
		return nil, err
	}

	return document, nil
}

// UpdateFile is the resolver for the updateFile field.
func (r *mutationResolver) UpdateFile(ctx context.Context, input interface{}) (*model.File, error) {
	document, err := r.Service.UpdateFile(ctx, input)
	if err != nil {
		return nil, err
	}

	return document, nil
}

// UploadFiles is the resolver for the uploadFiles field.
func (r *mutationResolver) UploadFiles(ctx context.Context, input []*model.UploadFile) ([]*model.File, error) {
	panic(fmt.Errorf("not implemented: UploadFiles - uploadFiles"))
}

// File is the resolver for the file field.
func (r *queryResolver) File(ctx context.Context, id string) (*model.File, error) {
	files, err := r.Service.GetFileByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return files, nil
}

// Files is the resolver for the files field.
func (r *queryResolver) Files(ctx context.Context, userID string) ([]*model.File, error) {
	files, err := r.Service.GetFiles(ctx, userID)
	if err != nil {
		return nil, err
	}

	return files, nil
}
