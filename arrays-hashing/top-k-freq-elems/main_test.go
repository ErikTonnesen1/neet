package main

import (
	"fmt"
	"testing"
)

// func TestMain(t *testing.T) {
// 	nums := []int{1, 2, 2, 3, 3, 3}
// 	k := 2
//
// 	fmt.Println(topKFrequent(nums, k))
// }

func TestMain2(t *testing.T) {
	nums := []int{7, 7}
	k := 1

	fmt.Println(topKFrequent(nums, k))
}
