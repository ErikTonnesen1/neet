package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	nums2 := []int{0, 0, 0, 0}
	nums3 := []int{-2, 0, 1, 1, 2}

	fmt.Println(threeSum(nums))
	fmt.Println(threeSum(nums2))
	fmt.Println(threeSum(nums3))

}

// [1, 2, 3, 4, 5]
// [1, 2] = 3
// nums[1, len(nums)].indexOf(-3) ? return triplet : move on

//sort input arr -> nlogn
// select first elem = a
// iterate thru array with two pointers,
// n + 1 = b, and len(arr)-1 = c
// 0 - n = arr[b] + arr[c]
// if arr[b] + arr[c] > 0-n => c--
// if arr[b] + arr[c] < 0-n => b++
// after all that, a++

/*
time-complexity
	sort = nlogn
	two for loops = n^2
	overal: n^2

space complexity
O(n) - map
O(n) - array

total O(n)
*/

type triplet struct {
	a, b, c int
}

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	threeSums := make([][]int, 0)

	sumMap := make(map[triplet]bool)

	for a := 0; a < len(nums)-1; a++ {
		b := a + 1
		c := len(nums) - 1

		target := 0 - nums[a]

		for b < c {
			sum := nums[b] + nums[c]
			if sum == target {
				_, ok := sumMap[triplet{nums[a], nums[b], nums[c]}]
				if !ok {
					sumMap[triplet{nums[a], nums[b], nums[c]}] = true
					threeSums = append(threeSums, []int{nums[a], nums[b], nums[c]})
				}
				b++
				c--
			} else if sum < target {
				b++
			} else if sum > target {
				c--
			}
		}
	}
	return threeSums
}
