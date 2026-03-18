package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	target := 7

	fmt.Println("Two sum indices: ", twoSum(slice, target))
}

func Test2(t *testing.T) {
	nums := []int{4, 5, 6}
	target := 10

	fmt.Println("Two sum indices: ", twoSum(nums, target))
}
