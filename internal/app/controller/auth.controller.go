package controller

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/cavelms/graph/generated"
	"github.com/cavelms/internal/model"
)

// SignIn is the resolver for the signIn field.
func (r *mutationResolver) SignIn(ctx context.Context, email string, password string) (*model.User, error) {
	auth, err := r.Service.SignIn(ctx, email, password)
	if err != nil {
		return nil, err
	}

	return auth, nil
}

// SignUp is the resolver for the signUp field.
func (r *mutationResolver) SignUp(ctx context.Context, fullName string, email string, password string, role string) (*model.User, error) {
	user, err := r.Service.SignUp(ctx, fullName, email, password, role)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// LogOut is the resolver for the logOut field.
func (r *mutationResolver) LogOut(ctx context.Context, email string, password string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: LogOut - logOut"))
}

// ForgetPassword is the resolver for the forgetPassword field.
func (r *mutationResolver) ForgetPassword(ctx context.Context, email string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: ForgetPassword - forgetPassword"))
}

// ResetPassword is the resolver for the resetPassword field.
func (r *mutationResolver) ResetPassword(ctx context.Context, email string, password string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: ResetPassword - resetPassword"))
}

// ChangePassword is the resolver for the changePassword field.
func (r *mutationResolver) ChangePassword(ctx context.Context, email string, token string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: ChangePassword - changePassword"))
}

// VerifyEmail is the resolver for the verifyEmail field.
func (r *mutationResolver) VerifyEmail(ctx context.Context, id string, code string) (*model.User, error) {
	user, err := r.Service.VerifyEmail(ctx, id, code)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Refresh is the resolver for the refresh field.
func (r *queryResolver) Refresh(ctx context.Context) (*model.User, error) {
	auth, err := r.Service.RefreshToken(ctx)
	if err != nil {
		return nil, err
	}

	return auth, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
