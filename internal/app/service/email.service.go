package service

import (
	"github.com/cavelms/internal/app/repository"
	"github.com/cavelms/internal/model"
	"github.com/cavelms/pkg/utils"
)

type mailService interface {
	SendMail(tpl string, data *model.NewMail) (*model.Mail, error)
	DeleteMail(*model.Mail) error
}

type mail struct {
	*repository.Repository
}

func newMailService(repo *repository.Repository) mailService {
	return &mail{
		Repository: repo,
	}
}

func (m *mail) SendMail(tpl string, data *model.NewMail) (*model.Mail, error) {
	body, err := utils.ParseTemplate(tpl, data.Body)
	if err != nil {
		return nil, err
	}

	mail := model.Mail{
		To:      data.To,
		Subject: data.Subject,
		Body:    body,
	}

	err = m.Mail.Send(mail)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
func (m *mail) DeleteMail(*model.Mail) error { return nil }
