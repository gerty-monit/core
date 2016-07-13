package monitors

import (
	"math/rand"
	"reflect"
)

type SmallInt struct {
	value int
}

func (SmallInt) Generate(r *rand.Rand, size int) reflect.Value {
	smallInt := SmallInt{r.Intn(50) + 1}
	return reflect.ValueOf(smallInt)
}
