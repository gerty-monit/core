package monitors

import "testing"

func TestShouldBeFineIfAllOk(t *testing.T) {
	buffer := NewCircularBuffer(5)
	values := []bool{true, true, true, true, true}

	for _, val := range values {
		buffer.Append(val)
	}

	if buffer.AllOk() == false {
		t.Fatalf("it should be all ok")
	}
}

func TestShouldRemoveOldValues(t *testing.T) {
	buffer := NewCircularBuffer(5)
	values := []bool{true, true, true, true, true, true, false, true, false, true}

	for _, val := range values {
		buffer.Append(val)
	}

	if buffer.AllOk() == true {
		t.Fatalf("it shouldnt be all ok")
	}
}
