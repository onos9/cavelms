package service

import (
	"context"
	"encoding/json"

	"github.com/cavelms/internal/app/repository"
	"github.com/cavelms/internal/model"
	"github.com/cavelms/pkg/mail"
	"github.com/cavelms/pkg/utils"
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

	user := model.User{ID: input.UserID}
	err = r.DB.FetchByID(&user)
	if err != nil {
		return nil, err
	}

	data := map[string]interface{}{
		"fullName":    user.FullName,
		"refereeName": input.FullName,
		"upload_link": "https://dev.beznet.org/signup/" + input.UserID,
	}

	body, err := utils.ParseTemplate("reference", data)
	if err != nil {
		return nil, err
	}

	msg := mail.Mailer{
		ToAddrs: []string{input.Email},
		Subject: "Adullam Reference",
		Body:    body,
	}

	// fs := mail.Template
	// file, err := fs.Open("template/reference_form.doc")
	// if err != nil {
	// 	return nil, err
	// }
	// defer file.Close()
	// fileInfo, _ := file.Stat()
	// size := fileInfo.Size()

	// buffer := make([]byte, size)
	// file.Read(buffer)
	// msg.Attachment = buffer
	// msg.Filename = fileInfo.Name()

	err = r.Mail.Send(msg)
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
