package main

import "fmt"

//Use a map to store sums:
// - k: sum , v: indices
// return the smallest indices first
// - O(n^2)

type indicePair struct {
	indices []int
}

func twoSum(nums []int, target int) []int {

	sumMap := make(map[int][]indicePair)

	for i, v := range nums {
		for j := i; j < len(nums); j++ {
			if i != j {
				sum := v + nums[j]
				sumMap[sum] = append(sumMap[sum], indicePair{[]int{i, j}})
			}
		}
	}

	fmt.Printf("SumMap: %+v\n", sumMap)

	smallestIndicePair := sumMap[target][0]
	return smallestIndicePair.indices
}
