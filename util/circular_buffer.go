package util

type CircularBuffer struct {
	Values []int
	next   int
}

func NewCircularBuffer(capacity int) CircularBuffer {
	sl := make([]int, capacity)
	return CircularBuffer{sl, 0}
}

func (buf *CircularBuffer) All(value int) bool {
	ok := true
	for _, val := range buf.Values {
		ok = ok && (val == value)
	}
	return ok
}

func (buf *CircularBuffer) Append(val int) {
	if buf.next == len(buf.Values) {
		buf.next = 0
	}
	
	buf.Values[buf.next] = val
	buf.next++
}
