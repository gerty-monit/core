package monitors

type Status int

const (
	UN Status = iota
	OK
	NOK
)

type Monitor interface {
	Stater
	Describer
}

type Stater interface {
	Check() bool
	Values() []Status
}

type Describer interface {
	Name() string
	Description() string
}
