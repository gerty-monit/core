package gerty

type Tripper interface {
	Trip()
	Untrip()
	IsTripped() bool
}

type Describer interface {
	Name() string
	Description() string
}

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

func (monitor *BaseMonitor) Name() string {
	return monitor.title
}

func (monitor *BaseMonitor) Description() string {
	return monitor.description
}
