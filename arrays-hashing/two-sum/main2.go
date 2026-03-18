package main

func twoSum2(nums []int, target int) []int {
	prevMap := make(map[int]int)
	for i, v := range nums {
		diff := target - v
		if _, found := prevMap[diff]; found != false {
			return []int{prevMap[diff], i}
		}
		prevMap[diff] = i
	}
	return []int{}
}
