package monitors

type CircularBuffer struct {
	Values []Status
	next   int
}

func NewCircularBuffer(capacity int) CircularBuffer {
	sl := make([]Status, capacity)
	return CircularBuffer{sl, 0}
}

func (buf CircularBuffer) AllOk() bool {
	ok := true
	for _, val := range buf.Values {
		ok = ok && (val == OK || val == UN)
	}
	return ok
}

func (buf CircularBuffer) AllNok() bool {
	ok := true
	for _, val := range buf.Values {
		ok = ok && (val == NOK || val == UN)
	}
	return ok
}

func (buf *CircularBuffer) Append(val bool) {
	if buf.next == len(buf.Values) {
		buf.next = 0
	}
	if val {
		buf.Values[buf.next] = OK
	} else {
		buf.Values[buf.next] = NOK
	}
	buf.next++
}
