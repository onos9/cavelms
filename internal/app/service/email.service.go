package service

import (
	"github.com/cavelms/internal/app/repository"
	"github.com/cavelms/internal/model"
)

type mailService interface {
	SendMail(*model.Mail) (*model.Mail, error)
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

func (m *mail) SendMail(*model.Mail) (*model.Mail, error) {return nil, nil}
func (m *mail) DeleteMail(*model.Mail) error              {return nil}
