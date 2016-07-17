package gerty

type Result int

const (
	UN Result = iota
	OK
	NOK
)

type BaseMonitor struct {
	title       string
	description string
	tripped     bool
}

func NewBaseMonitor(title, description string) *BaseMonitor {
	return &BaseMonitor{title, description, false}
}

func (m *BaseMonitor) Trip() {
	m.tripped = true
}

func (m *BaseMonitor) Untrip() {
	m.tripped = false
}

func (m *BaseMonitor) IsTripped() bool {
	return m.tripped
}

type Monitor interface {
	Stater
	Describer
	Tripper
}

type Tripper interface {
	Trip()
	Untrip()
	IsTripped() bool
}

type Stater interface {
	Check() Result
	Values() []ValueWithTimestamp
}

type Describer interface {
	Name() string
	Description() string
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
