/*
 * send email
 *
 */
package mail

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/smtp"
	"strings"

	"github.com/pkg/errors"
)

type Mailer struct {
	FromAddr string   `json:"fromAddr"`
	ToAddrs  []string `json:"toAddrs"`
	Subject  string   `json:"subject"`
	Body     string   `json:"body"`

	IMAPServer string `json:"imapserver"`
	SMTPServer string `json:"smtpserver"`
	Password   string `json:"password"`
	SSL        bool   `json:"ssl"`

	host string
	*smtp.Client
}

func init() {}

func (m *Mailer) SendMail() error {
	err := m.dialSmtp(m.SSL)
	if err != nil {
		return err
	}

	auth := smtp.PlainAuth("", m.FromAddr, m.Password, m.host)
	if err := m.Auth(auth); err != nil {
		return err
	}

	if err := m.Mail(m.FromAddr); err != nil {
		return err
	}

	recipients, err := m.getRecipients()
	if err != nil {
		return err
	}

	msg := m.getMessage(recipients)
	err = m.send(msg)
	if err != nil {
		return err
	}

	return m.Quit()
}

func (m *Mailer) dialSmtp(ssl bool) error {
	m.host, _, _ = net.SplitHostPort("smtppro.zoho.com:465")

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         m.host,
	}

	// Connect to non-secure SMTP Server
	if !ssl {
		c, err := smtp.Dial("smtppro.zoho.com:465")
		if err != nil {
			return err
		}
		c.StartTLS(tlsconfig)

		m.Client = c
		return nil
	}

	// // Connect to secure SMTP Server running on 465 with an ssl connection
	conn, err := tls.Dial("tcp", "smtppro.zoho.com:465", tlsconfig)
	if err != nil {
		return err
	}

	m.Client, err = smtp.NewClient(conn, m.host)
	if err != nil {
		return err
	}

	return nil
}

func (m *Mailer) send(msg string) error {
	wc, err := m.Data()
	if err != nil {
		return err
	}
	defer wc.Close()

	// write the body
	_, err = wc.Write([]byte(msg))
	if err != nil {
		return err
	}

	return nil
}

func (m *Mailer) getMessage(rcpt string) string {
	contents := []string{
		fmt.Sprintf("From: Adullam <%s>", m.FromAddr),
		fmt.Sprintf("To: %s", rcpt),
		fmt.Sprintf("Subject: %s", m.Subject),
		"MIME-Version: 1.0",
		"Content-Type: text/html; charset=\"utf-8\"",
		fmt.Sprintf("Body: %s", m.Body),
	}

	return strings.Join(contents, "\r\n")
}

func (m *Mailer) getRecipients() (string, error) {
	var errs []string
	var recipients []string
	for _, addr := range m.ToAddrs {
		if err := m.Rcpt(addr); err != nil {
			errs = append(errs, errors.Wrapf(err, "establishing mail recipient '%s'", addr).Error())
			continue
		}
		recipients = append(recipients, addr)
	}
	if len(errs) > 0 {
		return "", errors.New(strings.Join(errs, "; "))
	}
	return strings.Join(recipients, ", "), nil
}
