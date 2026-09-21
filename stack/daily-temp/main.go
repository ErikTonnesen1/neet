package main

import (
	"fmt"
	"slices"
)

func main() {
	input := []int{30, 38, 30, 36, 35, 40, 28}
	input2 := []int{89, 62, 70, 58, 47, 47, 46, 76, 100, 70}
	ans := dailyTemperatures(input)
	fmt.Println(ans)

	expected := []int{8, 1, 5, 4, 3, 2, 1, 1, 0, 0}
	ans2 := dailyTemperatures(input2)
	fmt.Println(ans2)
	fmt.Printf("Expected result: %t\n", slices.Equal(expected, ans2))
}

type stack []pair

func (s *stack) peek() pair {
	if len(*s) == 0 {
		return pair{}
	}
	return (*s)[len((*s))-1]
}

func (s *stack) push(p pair) {
	*s = append(*s, p)
}

func (s *stack) pop() pair {
	last := len(*s) - 1
	res := (*s)[last]
	*s = (*s)[:last]
	return res
}

type pair struct {
	value int
	index int
}

func peek[T any](arr []T) (top T, index int) {
	index = len(arr) - 1
	return arr[index], index
}

func dailyTemperatures(temperatures []int) []int {
	s := stack{}
	ans := make([]int, len(temperatures))
	for i, v := range temperatures {
		peek := s.peek()
		for len(s) != 0 && peek.value < v {
			pop := s.pop()
			ans[pop.index] = i - pop.index
			peek = s.peek()
		}
		if len(s) == 0 || peek.value >= v {
			s.push(pair{value: v, index: i})
		}
	}
	return ans
}
