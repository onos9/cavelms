/*
 * send email
 *
 */
package mail

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/smtp"
	"strings"

	"github.com/pkg/errors"
)

type Mailer struct {
	FromAddr   string   `json:"fromAddr"`
	ToAddrs    []string `json:"toAddrs"`
	CC         []string `json:"cc"`
	BCC        []string `json:"bcc"`
	Subject    string   `json:"subject"`
	Body       string   `json:"body"`
	Attachment []byte   `json:"attachment"`
	Filename   string   `json:"filename"`

	IMAPServer string `json:"imapserver"`
	SMTPServer string `json:"smtpserver"`
	Password   string `json:"password"`
	SSL        bool   `json:"ssl"`

	host     string
	boundary string
	auth     smtp.Auth
	*smtp.Client
}

func NewMailer(config *Mailer) error {
	m := Mailer{}

	b, err := json.Marshal(config)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	return nil
}

func (m *Mailer) SendMail() error {
	m.boundary = "**=myohmy689407924327"
	m.host, _, _ = net.SplitHostPort(m.SMTPServer)

	err := m.dialSmtp(m.SSL)
	if err != nil {
		return err
	}

	m.auth = smtp.PlainAuth("", m.FromAddr, m.Password, m.host)
	if err := m.Auth(m.auth); err != nil {
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

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         m.host,
	}

	// Connect to non-secure SMTP Server
	if !ssl {
		c, err := smtp.Dial(m.SMTPServer)
		if err != nil {
			return err
		}
		c.StartTLS(tlsconfig)

		m.Client = c
		return nil
	}

	// // Connect to secure SMTP Server running on 465 with an ssl connection
	conn, err := tls.Dial("tcp", m.SMTPServer, tlsconfig)
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
	h := m.setHeaders()
	c := m.setBody()
	contents := append(h, c...)
	if m.Filename != "" {
		a := m.attachFile()
		contents = append(contents, a...)
	}
	contents = append(contents, fmt.Sprintf("\r\n--MIXED-%s--", m.boundary))
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

func (m *Mailer) attachFile() []string {
	b := make([]byte, base64.StdEncoding.EncodedLen(len(m.Attachment)))
	base64.StdEncoding.Encode(b, m.Attachment)
	attachment := []string{
		fmt.Sprintf("\r\n--MIXED-%s", m.boundary),
		fmt.Sprintf("Content-Type: %s", http.DetectContentType(m.Attachment)),
		"Content-Transfer-Encoding: base64",
		fmt.Sprintf("Content-Disposition: attachment;filename=%s", m.Filename),
		fmt.Sprintf("\r\n%s", b),
	}

	return attachment
}

func (m *Mailer) setBody() []string {
	//place HTML message
	content := []string{
		fmt.Sprintf("\r\n%s", m.Body),
		fmt.Sprintf("\r\n--ALTERNATIVE-%s--", m.boundary),
		fmt.Sprintf("\r\n--RELATED-%s--", m.boundary),
	}

	return content
}

func (m *Mailer) setHeaders() []string {
	//basic email headers
	headers := []string{
		fmt.Sprintf("From: Adullam <%s>", m.FromAddr),
		fmt.Sprintf("To: %s", strings.Join(m.ToAddrs, ";")),
		fmt.Sprintf("Subject: %s", m.Subject),
	}
	if len(m.CC) > 0 {
		headers = append(headers, fmt.Sprintf("Cc: %s", strings.Join(m.CC, ";")))
	}

	if len(m.BCC) > 0 {
		headers = append(headers, fmt.Sprintf("Bcc: %s", strings.Join(m.BCC, ";")))
	}

	headers = append(headers, "MIME-Version: 1.0")
	headers = append(headers, fmt.Sprintf("Content-Type: multipart/mixed; boundary=MIXED-%s", m.boundary))
	headers = append(headers, fmt.Sprintf("\r\n--MIXED-%s", m.boundary))
	headers = append(headers, fmt.Sprintf("Content-Type: multipart/related; boundary=RELATED-%s", m.boundary))
	headers = append(headers, fmt.Sprintf("\r\n--RELATED-%s", m.boundary))
	headers = append(headers, fmt.Sprintf("Content-Type: multipart/alternative; boundary=ALTERNATIVE-%s", m.boundary))
	headers = append(headers, fmt.Sprintf("\r\n--ALTERNATIVE-%s", m.boundary))
	headers = append(headers, "Content-Type: text/html; charset=utf-8")

	return headers
}

// func init() {
// 	m := &Mailer{
// 		FromAddr:   "admin@adullam.ng",
// 		SMTPServer: "smtppro.zoho.com:465",
// 		SSL:        true,
// 		Password:   "#1414Bruno#",
// 	}

// 	err := NewMailer(m)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	file, err := os.Open("qservers.pdf")
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	defer file.Close()

// 	fileInfo, _ := file.Stat()
// 	var size int64 = fileInfo.Size()

// 	buffer := make([]byte, size)

// 	// read file content to buffer
// 	file.Read(buffer)

// 	// m.Attachment = buffer
// 	// m.Filename = fileInfo.Name()

// 	m.ToAddrs = []string{"onosbrown.saved@gmail.com"}
// 	m.Subject = "Golang example send mail in HTML format with attachment"
// 	m.Body = "<html><body><h1>Hi There</h1><p>this is sample email (with attachment) sent via golang program</p></body></html>"
// 	err = m.SendMail()
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func init() {
// 	m := gomail.NewMessage()
// 	m.SetHeader("From", "Adullam <admin@adullam.ng>")
// 	m.SetHeader("To", "onosbrown.saved@gmail.com")
// 	m.SetAddressHeader("Cc", "onosbrown.saved@gmail.com", "Dan")
// 	m.SetHeader("Subject", "Hello!")
// 	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
// 	// m.Attach("/home/Alex/lolcat.jpg")

// 	d := gomail.NewDialer("smtppro.zoho.com", 465, "admin@adullam.ng", "#1414Bruno#")

// 	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
// 	// Send the email to Bob, Cora and Dan.
// 	if err := d.DialAndSend(m); err != nil {
// 		log.Println(err)
// 	}
// }
