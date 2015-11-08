package monitors

import "time"

type ValueWithTimestamp struct {
	Value     int
	Timestamp int64
}

type CircularBuffer struct {
	Values []ValueWithTimestamp
	next   int
}

func NewCircularBuffer(capacity int) CircularBuffer {
	sl := make([]ValueWithTimestamp, capacity)
	return CircularBuffer{sl, 0}
}

func (buf *CircularBuffer) All(value int) bool {
	ok := true
	for _, val := range buf.Values {
		ok = ok && (val.Value == value)
	}
	return ok
}

func (buf *CircularBuffer) Append(val int) {
	if buf.next == len(buf.Values) {
		buf.next = 0
	}

	buf.Values[buf.next] = ValueWithTimestamp{val, time.Now().Unix()}
	buf.next++
}
