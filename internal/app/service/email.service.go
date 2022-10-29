package service

import (
	"github.com/cavelms/internal/app/repository"
	"github.com/cavelms/internal/model"
	"github.com/cavelms/pkg/mail"
	"github.com/cavelms/pkg/utils"
)

type mailService interface {
	SendMail(tpl string, data *model.NewMail) (*model.Mail, error)
	DeleteMail(*model.Mail) error
}

type mailer struct {
	*repository.Repository
}

func newMailService(repo *repository.Repository) mailService {
	return &mailer{
		Repository: repo,
	}
}

func (m *mailer) SendMail(tpl string, data *model.NewMail) (*model.Mail, error) {
	body, err := utils.ParseTemplate(tpl, data.Body)
	if err != nil {
		return nil, err
	}

	msg := mail.Mailer{
		ToAddrs: data.To,
		Subject: data.Subject,
		Body:    body,
	}

	if data.Attach {
		fs := mail.Template
		file, err := fs.Open("template/reference_form.doc")
		if err != nil {
			return nil, err
		}
		defer file.Close()
		fileInfo, _ := file.Stat()
		size := fileInfo.Size()

		buffer := make([]byte, size)
		file.Read(buffer)
		msg.Attachment = buffer
		msg.Filename = fileInfo.Name()
	}

	err = m.Mail.Send(msg)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
func (m *mailer) DeleteMail(*model.Mail) error { return nil }
