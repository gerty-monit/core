package gerty

import (
	"net/smtp"
)

// ensure we always implement Alarm
var _ Alarm = (*EmailAlarm)(nil)

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

func (alarm EmailAlarm) NotifyError(monitor Monitor) error {
	auth := smtp.PlainAuth("", alarm.User, alarm.Pass, alarm.Host)
	data := emailContents{alarm.From, alarm.To, "http://status.galeno.omniasalud.com",
		monitor.Name(), monitor.Description()}

	bytes, err := EmailTemplate(data)
	if err != nil {
		logger.Printf("error while creating the email template %v", err)
		return err
	}

	address := alarm.Host + ":" + alarm.Port
	logger.Printf("template generated, sending email to `%s`", address)
	err = smtp.SendMail(address, auth, alarm.From, []string{alarm.To}, bytes)
	if err != nil {
		logger.Printf("error while sending the alarm email %v", err)
		return err
	}

	return nil
}

// TODO: create email template with restore message.
func (alarm EmailAlarm) NotifyRestore(monitor Monitor) error {
	return nil
}
