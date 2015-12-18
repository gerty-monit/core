package monitors

const (
	UN = iota
	OK
	NOK
)

type Monitor interface {
	Stater
	Describer
}

type Stater interface {
	Check() int
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

func AllFailed(m Monitor) bool {
	for i := range m.Values() {
		if m.Values()[i].Value != NOK {
			return false
		}
	}
	return true
}
