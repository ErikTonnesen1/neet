package main

import (
	"fmt"
	"slices"
)

func main() {
	//nums := []int{1, 2, 3, -3, 0}
	nums2 := []int{-1, 0, 1, 2, -1, -4}
	// fmt.Println(threeSum(nums))

	fmt.Println(threeSum(nums2))

}

// [1, 2, 3, 4, 5]
// [1, 2] = 3
// nums[1, len(nums)].indexOf(-3) ? return triplet : move on

func threeSum(nums []int) [][]int {
	fmt.Println(nums)
	threeSums := make([][]int, 0)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			neededValue := 0 - sum

			index := slices.Index(nums, neededValue)
			fmt.Printf("Needed num: %d, Index: %d\n", neededValue, index)
			if index != -1 {
				threeSums = append(threeSums, []int{nums[i], nums[j], nums[index]})
			}
		}
	}
	return threeSums
}
