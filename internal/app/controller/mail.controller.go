package controller

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/cavelms/internal/model"
)

// Send is the resolver for the send field.
func (r *mutationResolver) Send(ctx context.Context, input *model.NewMail) (*model.Mail, error) {
	panic(fmt.Errorf("not implemented: Send - send"))
}

// DeleteMail is the resolver for the deleteMail field.
func (r *mutationResolver) DeleteMail(ctx context.Context, id string) (*model.Mail, error) {
	panic(fmt.Errorf("not implemented: DeleteMail - deleteMail"))
}

// Mail is the resolver for the mail field.
func (r *queryResolver) Mail(ctx context.Context, id string) (*model.Mail, error) {
	panic(fmt.Errorf("not implemented: Mail - mail"))
}

// Mails is the resolver for the mails field.
func (r *queryResolver) Mails(ctx context.Context) ([]*model.Mail, error) {
	panic(fmt.Errorf("not implemented: Mails - mails"))
}
