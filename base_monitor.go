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
	name        string
	description string
	tripped     bool
}

func NewBaseMonitor(name, description string) *BaseMonitor {
	return &BaseMonitor{name, description, false}
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
	return monitor.name
}

func (monitor *BaseMonitor) Description() string {
	return monitor.description
}
