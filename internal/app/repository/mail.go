package repository

import (
	"github.com/cavelms/internal/model"
	"github.com/cavelms/pkg/mail"
)

type Mail interface {
	Send(tpl string, m model.Mail) error
	Delete(id string) error
}

type mailClient struct {
	*mail.Mailer
}

func newEmailRepository() Mail {
	m := &mail.Mailer{
		FromAddr: "admin@adullam.ng",
		Password: "#1414Bruno#",
		SSL:      true,
	}
	return &mailClient{m}
}

func (c *mailClient) Send(tpl string, m model.Mail) error {
	c.ToAddrs = m.ToAddrs
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
