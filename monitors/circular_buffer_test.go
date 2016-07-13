package monitors

import "testing"

var cases = []struct {
	name           string
	arr            []int
	capacity       int
	expectedLength int
}{
	{"appending no elements should have length zero", []int{}, 5, 0},
	{"appending few elements should have pararmeters length", []int{1, 2, 3, 4}, 5, 4},
	{"appending more elements than capacity should have cap length", []int{1, 2, 3, 4}, 2, 2},
}

func contains(arr []int, i int) bool {
	for _, j := range arr {
		if i == j {
			return true
		}
	}
	return false
}

func TestCircularBufferAppending(t *testing.T) {
	for _, c := range cases {
		buffer := NewCircularBuffer(c.capacity)
		for _, v := range c.arr {
			buffer.Append(v)
		}

		retrieved := buffer.GetValues()

		// test length.
		if len(retrieved) != c.expectedLength {
			t.Errorf(c.name)
		}

		// sanity check: test all returned elements are present.

		for _, r := range retrieved {
			if !contains(c.arr, r.Value) {
				t.Errorf("element %d was retrieved but was not present in original array (%v)", c.arr)
			}
		}
	}
}

func TestShouldBeFineIfAllOk(t *testing.T) {
	buffer := NewCircularBuffer(5)
	values := []int{0, 0, 0, 0, 0}

	for _, val := range values {
		buffer.Append(val)
	}

	if buffer.All(0) == false {
		t.Errorf("it should be all zeroes")
	}
}

func TestShouldRemoveOldValues(t *testing.T) {
	buffer := NewCircularBuffer(5)
	values := []int{0, 0, 0, 0, 0, 0, 1, 0, 1, 0}

	for _, val := range values {
		buffer.Append(val)
	}

	if buffer.All(0) {
		t.Errorf("it shouldnt be all zeroes")
	}
}
