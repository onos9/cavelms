package controller

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/cavelms/internal/model"
)

// CreateNote is the resolver for the createNote field.
func (r *mutationResolver) CreateNote(ctx context.Context, input model.NewNote) (*model.Note, error) {
	panic(fmt.Errorf("not implemented: CreateNote - createNote"))
}

// Note is the resolver for the note field.
func (r *queryResolver) Note(ctx context.Context, id string) (*model.Note, error) {
	panic(fmt.Errorf("not implemented: Note - note"))
}

// Notes is the resolver for the notes field.
func (r *queryResolver) Notes(ctx context.Context, limit *int, offset *int) ([]*model.Note, error) {
	panic(fmt.Errorf("not implemented: Notes - notes"))
}
