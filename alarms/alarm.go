package alarms

import (
	"github.com/gerty-monit/core/monitors"
	"log"
	"os"
)

var logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

type Alarm interface {
	Name() string
	NotifyError(monitors.Monitor) error
	NotifyRestore(monitors.Monitor) error
}
