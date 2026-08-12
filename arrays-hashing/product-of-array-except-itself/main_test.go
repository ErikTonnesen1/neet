package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	strs := []int{1, 2, 3, 4} //24
	fmt.Println(productExceptSelf(strs))
}
