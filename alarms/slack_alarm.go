package alarms

import (
	"github.com/gerty-monit/core/monitors"
	"net/http"
	"strings"
)

type SlackAlarm struct {
	WebhookUrl string
}

func NewSlackAlarm(url string) *SlackAlarm {
	return &SlackAlarm{url}
}

func (alarm SlackAlarm) Name() string {
	return "Slack Channel Alarm"
}

func (alarm SlackAlarm) NotifyError(monitor monitors.Monitor) {
	body := strings.NewReader(`{ "text": "` + monitor.Name() + ` has raised an alarm" }`)
	_, err := http.Post(alarm.WebhookUrl, "application/json", body)

	if err != nil {
		logger.Printf("error while reporting to slack: %v", err)
	}
}
