package repository

import (
	"log"

	"github.com/cavelms/internal/model"
	"github.com/cavelms/pkg/mail"
)

type Mail interface {
	Send(m model.Mail) error
	Delete(id string) error
}

type mailClient struct {
	*mail.Mailer
}

func newEmailRepository() Mail {
	m := &mail.Mailer{
		FromAddr: "admin@adullam.ng",
		Password: "#1414Bruno#",
		SMTPServer: "smtppro.zoho.com:465",
		SSL:      true,
	}
	err := mail.NewMailer(m)
	if err != nil {
		log.Fatalln(err)
	}
	return &mailClient{m}
}

func (c *mailClient) Send(m model.Mail) error {
	c.ToAddrs = m.To
	c.Body = m.Body
	c.Subject = m.Subject

	err := c.SendMail()
	if err != nil {
		return err
	}

	return nil
}

func (c *mailClient) Delete(id string) error {

	return nil
}
