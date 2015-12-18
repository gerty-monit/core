package alarms

import (
	"github.com/gerty-monit/core/monitors"
)

type Alarm interface {
	Name() string
	NotifyError(monitors.Monitor)
}
