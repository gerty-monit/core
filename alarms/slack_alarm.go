package alarms

import (
	"github.com/gerty-monit/core/monitors"
	"net/http"
	"strings"
)

type SlackAlarm struct {
	WebhookUrl string
}

// ensure we always implement Alarm
var _ Alarm = (*SlackAlarm)(nil)

func NewSlackAlarm(url string) *SlackAlarm {
	return &SlackAlarm{url}
}

func (alarm SlackAlarm) Name() string {
	return "Slack Channel Alarm"
}

func (alarm *SlackAlarm) sendToSlack(message string) error {
	body := strings.NewReader(message)
	_, err := http.Post(alarm.WebhookUrl, "application/json", body)

	if err != nil {
		logger.Printf("error while reporting to slack: %v", err)
		return err
	}

	return nil
}

func (alarm *SlackAlarm) NotifyError(monitor monitors.Monitor) error {
	return alarm.sendToSlack(`{ "text": "` + monitor.Name() + ` has raised an alarm" }`)
}

func (alarm *SlackAlarm) NotifyRestore(monitor monitors.Monitor) error {
	return alarm.sendToSlack(`{ "text": "` + monitor.Name() + ` is back to normal" }`)
}
