package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cavelms/internal/app/repository"
	"github.com/cavelms/internal/model"
)

type userService interface {
	CreateUser(ctx context.Context, input model.NewUser) (*model.User, error)
	UpdateUser(ctx context.Context, data interface{}) (*model.User, error)
	DeleteUser(ctx context.Context, id string) (*model.User, error)
	DeleteManyUsers(ctx context.Context, id []string) (*model.User, error)
	GetUsers(ctx context.Context) ([]*model.User, error)
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
}

type user struct {
	*repository.Repository
}

func newUserService(repo *repository.Repository) userService {
	return &user{repo}
}

// CreateUser is the resolver for the createUser field.
func (u *user) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

// UpdateUser is the resolver for the updateUser field.
func (u *user) UpdateUser(ctx context.Context, data interface{}) (*model.User, error) {
	user := model.User{}

	ub, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(ub, &user)
	if err != nil {
		return nil, err
	}

	err = u.DB.UpdateOne(user)
	if err != nil {
		return nil, err
	}

	err = u.DB.FetchByID(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// DeleteUser is the resolver for the deleteUser field.
func (u *user) DeleteUser(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: DeleteUser - deleteUser"))
}

// DeleteManyUsers is the resolver for the deleteManyUsers field.
func (u *user) DeleteManyUsers(ctx context.Context, id []string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: DeleteManyUsers - deleteManyUsers"))
}

// GetUsers is the resolver for the getUsers field.
func (u *user) GetUsers(ctx context.Context) ([]*model.User, error) {
	user := new(model.User)
	users := []model.User{}
	err := u.DB.FetchAll(&users, user)
	if err != nil {
		return nil, err
	}

	userList := []*model.User{}
	for i := 0; i < len(users); i++{
		userList = append(userList, &users[i])
	}

	return userList, nil
}

// GetUserByID is the resolver for the getUserById field.
func (u *user) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	user := new(model.User)
	user.ID = id
	err := u.DB.FetchByID(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByEmail is the resolver for the getUserByEmail field.
func (u *user) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: GetUserByEmail - getUserByEmail"))
}
