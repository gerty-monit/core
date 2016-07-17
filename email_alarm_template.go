package gerty

import (
	"bytes"
	"html/template"
	"os"
)

type emailContents struct {
	From               string
	To                 string
	Site               string
	MonitorName        string
	MonitorDescription string
}

var (
	file = os.Getenv("GOPATH") + "/src/github.com/gerty-monit/core/email-templates/email_alarm.html"
	tpl  = template.Must(template.New("email_alarm.html").ParseFiles(file))
)

func EmailTemplate(contents emailContents) ([]byte, error) {
	var buffer bytes.Buffer
	err := tpl.Execute(&buffer, contents)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
