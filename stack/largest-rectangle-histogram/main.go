package main

import "fmt"

/*
Input: heights = [7,1,7,2,2,4]
Output: 8

Constraints:
- len(input) >= 1
- height >= 0


Flow

1. Iterate through heights, build pair {height, start}

	If height >= previous height => push onto stack
		by pushing onto stack, the previous items have their right bound increased by 1
	If height < previous height =>
		while height < previous height
			pop off top of stack
			find area of popped -> pop.height * (RB - LB)
			if area > largestArea => largestArea = are

			current.start - 1 (since it could fit in )

*/

type pair struct {
	height int
	start  int
}

type stack []*pair

func (s *stack) peek() *pair {
	if s.isEmpty() {
		return nil
	}
	return (*s)[len(*s)-1]
}

func (s *stack) pop() *pair {
	if s.isEmpty() {
		return nil
	}
	p := s.peek()
	*s = (*s)[:len(*s)-1]
	return p
}
func (s *stack) push(p *pair) {
	*s = append(*s, p)
}

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func largestRectangleArea(heights []int) int {
	var largestArea int
	s := stack{}
	for i, v := range heights {
		p := pair{
			height: v,
			start:  i,
		}

		if s.isEmpty() {
			s.push(&p)
			continue
		}

		peek := s.peek()
		if p.height > peek.height {
			s.push(&p)
			continue
		}
		if p.height == peek.height {
			continue
		}

		leftBound := 0
		for s.peek() != nil && p.height < s.peek().height {
			pop := s.pop()
			area := pop.height * (p.start - pop.start)
			largestArea = max(largestArea, area)
			leftBound = pop.start
		}
		p.start = leftBound
		s.push(&p)
	}

	for !s.isEmpty() {
		p := s.pop()
		area := p.height * (len(heights) - p.start)
		largestArea = max(largestArea, area)
	}

	return largestArea
}

func main() {
	input := []int{7, 1, 7, 2, 2, 4}
	area := largestRectangleArea(input)
	fmt.Printf("Area: %d\n", area)

	heights := []int{2, 1, 5, 6, 2, 3}
	area = largestRectangleArea(heights)
	fmt.Printf("Area: %d\n", area)
}
