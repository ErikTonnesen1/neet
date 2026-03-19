/*
Problem:
	Given an integer array nums and an integer k, return the k most frequent elements within the array.

	The test cases are generated such that the answer is always unique.

	You may return the output in any order.
*/

/*
	Example 1:

	Input: nums = [1,2,2,3,3,3], k = 2

	Output: [2,3]
	Example 2:

	Input: nums = [7,7], k = 1

	Output: [7]
*/

package main

import (
	"sort"
)

/*
	Immediate thoughts
	- create hashmap k: num, v: # of times in nums
	- after map is created, can create a slice of the v's and sort
	- return [0] and [1] of the sorted slice
*/

func topKFrequent(nums []int, k int) []int {
	frequencyMap := make(map[int]int)
	for _, v := range nums {
		if _, found := frequencyMap[v]; found == true {
			frequencyMap[v]++
		} else {
			frequencyMap[v] = 1
		}
	}

	var freqSlice [][]int

	for k, v := range frequencyMap {
		freqSlice = append(freqSlice, []int{k, v})
	}

	sort.Slice(freqSlice, func(i, j int) bool {
		return freqSlice[i][1] > freqSlice[j][1] // > for descending
	})

	var result []int
	for i := range k {
		result = append(result, freqSlice[i][0])
	}
	return result
}
