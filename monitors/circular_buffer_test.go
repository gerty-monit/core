package monitors

import (
	"testing"
	"testing/quick"
)

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

func TestCircularBufferLength(t *testing.T) {
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
				t.Errorf("element %d was retrieved but was not present in original array (%v)", r.Value, c.arr)
			}
		}
	}
}

func containsAll(container, arr []int) bool {
	for i := range arr {
		if !contains(container, arr[i]) {
			return false
		}
	}
	return true
}

func contains(arr []int, i int) bool {
	for _, j := range arr {
		if i == j {
			return true
		}
	}
	return false
}

func TestShouldAlwaysReturnLastAppendedValues(t *testing.T) {

	f := func(gencap SmallInt, elementsToAdd []int) bool {
		capacity := gencap.value
		start := 0
		if len(elementsToAdd)-capacity > 0 {
			start = len(elementsToAdd) - capacity
		}
		expected := elementsToAdd[start:]
		buffer := NewCircularBuffer(capacity)
		for _, el := range elementsToAdd {
			buffer.Append(el)
		}

		retrieved := []int{}
		for _, v := range buffer.GetValues() {
			retrieved = append(retrieved, v.Value)
		}
		return containsAll(retrieved, expected)
	}

	if err := quick.Check(f, nil); err != nil {
		t.Errorf("error %v", err)
	}
}
