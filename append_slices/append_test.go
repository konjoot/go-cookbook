package append_slices

import (
	"testing"
)

func TestSliceAppend(t *testing.T) {
	slice := make([]int, 4, 8)

	slice1 := slice[:]

	for i := 0; i < 10; i++ {
		slice1[i%4] = i
		slice = append(slice, 99)
		t.Logf("slice1 = %v", slice1)
		t.Logf("slice  = %v", slice)
	}
}
