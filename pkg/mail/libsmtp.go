package mail

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"net/smtp"
	"path/filepath"
	"strings"
)

var (
	address  = "smtppro.zoho.com:465"
	username = "admin@adullam.ng"
	password = "#1414Bruno#"
)

type Sender struct {
	auth smtp.Auth
}

type Message struct {
	To          []string
	CC          []string
	BCC         []string
	Subject     string
	Body        string
	Attachments map[string][]byte
}

func New() *Sender {
	host, _, _ := net.SplitHostPort(address)
	auth := smtp.PlainAuth("", username, password, host)
	return &Sender{auth}
}

func (s *Sender) Send(m *Message) error {
	host, _, _ := net.SplitHostPort(address)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	// // Connect to secure SMTP Server running on 465 with an ssl connection
	conn, err := tls.Dial("tcp", address, tlsconfig)
	if err != nil {
		return err
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		return err
	}

	if err := c.Auth(s.auth); err != nil {
		return err
	}

	if err := c.Mail(username); err != nil {
		return err
	}

	if err := c.Rcpt(strings.Join(m.To, ", ")); err != nil {
		return err
	}

	wc, err := c.Data()
	if err != nil {
		return err
	}
	defer wc.Close()

	// write the body
	_, err = wc.Write(m.ToBytes())
	if err != nil {
		return err
	}

	return nil
}

func NewMessage(s, b string) *Message {
	return &Message{Subject: s, Body: b, Attachments: make(map[string][]byte)}
}

func (m *Message) AttachFile(src string) error {
	b, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	_, fileName := filepath.Split(src)
	m.Attachments[fileName] = b
	return nil
}

func (m *Message) ToBytes() []byte {
	buf := bytes.NewBuffer(nil)
	withAttachments := len(m.Attachments) > 0
	buf.WriteString(fmt.Sprintf("Subject: %s\n", m.Subject))
	buf.WriteString(fmt.Sprintf("To: %s\n", strings.Join(m.To, ",")))
	if len(m.CC) > 0 {
		buf.WriteString(fmt.Sprintf("Cc: %s\n", strings.Join(m.CC, ",")))
	}

	if len(m.BCC) > 0 {
		buf.WriteString(fmt.Sprintf("Bcc: %s\n", strings.Join(m.BCC, ",")))
	}

	buf.WriteString("MIME-Version: 1.0\n")
	writer := multipart.NewWriter(buf)
	boundary := writer.Boundary()
	if withAttachments {
		buf.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\n", boundary))
		buf.WriteString(fmt.Sprintf("--%s\n", boundary))
	} else {
		buf.WriteString("Content-Type: text/html; charset=utf-8\n")
	}

	buf.WriteString(m.Body)
	if withAttachments {
		for k, v := range m.Attachments {
			buf.WriteString(fmt.Sprintf("\n\n--%s\n", boundary))
			buf.WriteString(fmt.Sprintf("Content-Type: %s\n", http.DetectContentType(v)))
			buf.WriteString("Content-Transfer-Encoding: base64\n")
			buf.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=%s\n", k))

			b := make([]byte, base64.StdEncoding.EncodedLen(len(v)))
			base64.StdEncoding.Encode(b, v)
			buf.Write(b)
			buf.WriteString(fmt.Sprintf("\n--%s", boundary))
		}

		buf.WriteString("--")
	}

	fmt.Println(buf.String())
	return buf.Bytes()
}

// func init() {
// 	sender := New()
// 	m := NewMessage("Golang example send mail in HTML format with attachment", "<html><body><h1>Hi There</h1><p>this is sample email (with attachment) sent via golang program</p></body></html>")
// 	m.To = []string{"Adullam <onosbrown.saved@gmail.com>"}
// 	m.CC = []string{"onos9@gmail.com", "suport@adullam.ng"}
// 	m.BCC = []string{"admin@dullam.ng"}
// 	// m.AttachFile("/path/to/file")
// 	fmt.Println(sender.Send(m))
// }
