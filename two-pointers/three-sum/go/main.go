package main

import (
	"fmt"
	"slices"
)

func main() {
	input := []int{0, 0, 0}
	input2 := []int{-1, 0, 1, 2, -1, -4}
	input3 := []int{0, 0, 0, 0}
	input4 := []int{-2, 0, 1, 1, 2}
	fmt.Println(threeSum(input))
	fmt.Println(threeSum(input2))
	fmt.Println(threeSum(input3))
	fmt.Println(threeSum(input4))
}

func threeSum(nums []int) [][]int {
	slices.Sort(nums)

	answerMap := map[[3]int][]int{}

	for i := 0; i < len(nums)-1; i++ {
		first := nums[i]
		for j := i + 1; j < len(nums); j++ {
			second := nums[j]
			neededThird := 0 - (first + second)

			remainingAvailableThirds := nums[j+1:]

			index := slices.Index(remainingAvailableThirds, neededThird)
			if index != -1 {
				candidate := [3]int{first, second, neededThird}
				_, duplicate := answerMap[candidate]
				if !duplicate {
					answerMap[candidate] = []int{first, second, neededThird}
				}
			}

		}
	}
	ans := make([][]int, 0)
	for _, v := range answerMap {
		ans = append(ans, v)
	}
	return ans
}
