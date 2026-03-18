package dynamicarray

import (
	"fmt"
	"testing"
)

func TestSlices(t *testing.T) {
	slice := make([]int, 1)

	fmt.Printf("Slice length: %d, slice capacity: %d\n", len(slice), cap(slice))

	//
	// fmt.Printf("Slice length: %d, slice capacity: %d\n", len(slice), cap(slice))
}
