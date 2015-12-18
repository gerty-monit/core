package alarms

import (
	"bytes"
	"github.com/gerty-monit/core/monitors"
	"html/template"
	"log"
	"net/smtp"
	"os"
)

var logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

var templateFile = os.Getenv("GOPATH") + "/src/github.com/gerty-monit/core/alarms/email_alarm.html"
var tpl = template.Must(template.New("email_alarm.html").ParseFiles(templateFile))

type emailData struct {
	From    string
	To      string
	Site    string
	Monitor monitors.Monitor
}

type EmailAlarm struct {
	Host string
	Port string
	User string
	Pass string
	From string
	To   string
	Home string
}

func NewEmailAlarm(host, port, user, pass, from, to, home string) *EmailAlarm {
	return &EmailAlarm{host, port, user, pass, from, to, home}
}

func (alarm EmailAlarm) Name() string {
	return "Email Alarm"
}

func (alarm EmailAlarm) NotifyError(monitor monitors.Monitor) {
	auth := smtp.PlainAuth("", alarm.User, alarm.Pass, alarm.Host)
	to := alarm.To
	var buffer bytes.Buffer
	data := emailData{alarm.From, to, "http://status.galeno.omniasalud.com", monitor}

	err := tpl.Execute(&buffer, data)
	if err != nil {
		logger.Printf("error while creating the email template %v", err)
	}

	address := alarm.Host + ":" + alarm.Port
	logger.Printf("template generated, sending email to `%s`", address)
	err = smtp.SendMail(address, auth, alarm.From, []string{to}, buffer.Bytes())
	if err != nil {
		logger.Printf("error while seding the alarm email %v", err)
	}
}
