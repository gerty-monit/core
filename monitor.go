package gerty

type Result int

const (
	UN Result = iota
	OK
	NOK
)

type Monitor interface {
	Stater
	Describer
	Tripper
}

type Stater interface {
	Check() Result
	Values() []ValueWithTimestamp
}

type Group struct {
	Name     string
	Monitors []Monitor
}

func all(m Monitor, status Result) bool {
	for i := range m.Values() {
		if m.Values()[i].Value != status {
			return false
		}
	}
	return true
}

func AllFailed(m Monitor) bool {
	return all(m, NOK)
}

func AllOk(m Monitor) bool {
	return all(m, OK)
}
