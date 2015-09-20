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
	Check() bool
	Values() []int
}

type Describer interface {
	Name() string
	Description() string
}
