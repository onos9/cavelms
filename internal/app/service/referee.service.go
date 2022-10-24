package service

import (
	"context"
	"encoding/json"

	"github.com/cavelms/internal/app/repository"
	"github.com/cavelms/internal/model"
)

type refereeService interface {
	CreateReferee(ctx context.Context, input model.NewReferee) (*model.Referee, error)
	UpdateReferee(ctx context.Context, data interface{}) (*model.Referee, error)
	GetRefereeByID(ctx context.Context, id string) (*model.Referee, error)
	GetReferees(ctx context.Context, userId *string) ([]*model.Referee, error)
}

type referee struct {
	*repository.Repository
}

func newRefereeService(repo *repository.Repository) refereeService {
	return &referee{
		Repository: repo,
	}
}

func (r *referee) CreateReferee(ctx context.Context, input model.NewReferee) (*model.Referee, error) {
	referee := model.Referee{
		UserID:   input.UserID,
		FullName: input.FullName,
		Email:    input.Email,
		Phone:    input.Phone,
	}

	err := r.DB.Create(&referee)
	if err != nil {
		return nil, err
	}
	return &referee, nil
}
func (r *referee) UpdateReferee(ctx context.Context, data interface{}) (*model.Referee, error) {
	referee := model.Referee{}

	ub, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(ub, &referee)
	if err != nil {
		return nil, err
	}

	err = r.DB.UpdateOne(&referee)
	if err != nil {
		return nil, err
	}

	err = r.DB.FetchByID(&referee)
	if err != nil {
		return nil, err
	}

	return &referee, nil
}
func (r *referee) GetReferees(ctx context.Context, userId *string) ([]*model.Referee, error) {
	referee := new(model.Referee)
	referees := []model.Referee{}

	if userId != nil {
		referee.UserID = *userId
		err := r.DB.FetchByUserID(&referees, referee)
		if err != nil {
			return nil, err
		}
	} else {
		err := r.DB.FetchAll(&referees, referee)
		if err != nil {
			return nil, err
		}
	}

	fileList := []*model.Referee{}
	for i := 0; i < len(referees); i++ {
		fileList = append(fileList, &referees[i])
	}

	return fileList, nil
}

func (r *referee) GetRefereeByID(ctx context.Context, id string) (*model.Referee, error) {
	referee := new(model.Referee)
	referee.ID = id
	err := r.DB.FetchByID(referee)
	if err != nil {
		return nil, err
	}

	return referee, nil
}
