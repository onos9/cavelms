package utils

import (
	"bytes"
	"html/template"
	"io/ioutil"

	"github.com/cavelms/pkg/mail"
)

func ParseHtml(f string) (string, error) {
	bs, err := ioutil.ReadFile(f)
	if err != nil {
		return "", err
	}

	return string(bs), nil
}

// func init() {
// 	data := map[string]interface{}{
// 		"code":     12345,
// 		"fullname": "fullName",
// 	}

// 	body, err := ParseTemplate("signup", data)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println(body)
// }

func ParseTemplate(tpl string, m interface{}) (string, error) {

	fs := mail.Template
	t, err := template.ParseFS(fs, "template/"+tpl+".html")
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, m); err != nil {
		return "", err
	}

	return buf.String(), nil
}
