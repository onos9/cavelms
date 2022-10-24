package service

import (
	"context"
	"encoding/json"

	"github.com/cavelms/internal/app/repository"
	"github.com/cavelms/internal/model"
)

type educationService interface {
	CreateEducation(ctx context.Context, input model.NewEducation) (*model.Education, error)
	UpdateEducation(ctx context.Context, data interface{}) (*model.Education, error)
	GetEducationByID(ctx context.Context, id string) (*model.Education, error)
	GetEducations(ctx context.Context, userId *string) ([]*model.Education, error)
}

type education struct {
	*repository.Repository
}

func newEducationService(repo *repository.Repository) educationService {
	return &education{
		Repository: repo,
	}
}

func (r *education) CreateEducation(ctx context.Context, input model.NewEducation) (*model.Education, error) {
	education := model.Education{
		UserID:   input.UserID,
		Institution: input.Institution,
		Degree:    input.Degree,
		GraduationYear:    input.GraduationYear,
	}

	err := r.DB.Create(&education)
	if err != nil {
		return nil, err
	}
	return &education, nil
}
func (r *education) UpdateEducation(ctx context.Context, data interface{}) (*model.Education, error) {
	education := model.Education{}

	ub, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(ub, &education)
	if err != nil {
		return nil, err
	}

	err = r.DB.UpdateOne(&education)
	if err != nil {
		return nil, err
	}

	err = r.DB.FetchByID(&education)
	if err != nil {
		return nil, err
	}

	return &education, nil
}
func (r *education) GetEducations(ctx context.Context, userId *string) ([]*model.Education, error) {
	education := new(model.Education)
	educations := []model.Education{}

	if userId != nil {
		education.UserID = *userId
		err := r.DB.FetchByUserID(&educations, education)
		if err != nil {
			return nil, err
		}
	} else {
		err := r.DB.FetchAll(&educations, education)
		if err != nil {
			return nil, err
		}
	}

	fileList := []*model.Education{}
	for i := 0; i < len(educations); i++ {
		fileList = append(fileList, &educations[i])
	}

	return fileList, nil
}

func (r *education) GetEducationByID(ctx context.Context, id string) (*model.Education, error) {
	education := new(model.Education)
	education.ID = id
	err := r.DB.FetchByID(education)
	if err != nil {
		return nil, err
	}

	return education, nil
}
