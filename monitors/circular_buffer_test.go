package monitors

import "testing"

func TestShouldBeFineIfAllOk(t *testing.T) {
	buffer := NewCircularBuffer(5)
	values := []int{0, 0, 0, 0, 0}

	for _, val := range values {
		buffer.Append(val)
	}

	if buffer.All(0) == false {
		t.Fatalf("it should be all zeroes")
	}
}

func TestShouldRemoveOldValues(t *testing.T) {
	buffer := NewCircularBuffer(5)
	values := []int{0, 0, 0, 0, 0, 0, 1, 0, 1, 0}

	for _, val := range values {
		buffer.Append(val)
	}

	if buffer.All(0) {
		t.Fatalf("it shouldnt be all zeroes")
	}
}
