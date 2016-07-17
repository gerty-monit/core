package gerty

type Alarm interface {
	Name() string
	NotifyError(Monitor) error
	NotifyRestore(Monitor) error
}
