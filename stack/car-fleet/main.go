package main

import (
	"fmt"
	"slices"
)

func main() {
	var (
		target   = 0
		expected = 0
		fleets   = 0
	)

	target = 10
	var (
		position = []int{1, 4}
		speed    = []int{3, 2}
	)
	expected = 1

	fleets = carFleet(target, position, speed)
	fmt.Printf("Scenario 1, had %d, fleets.\n", fleets)
	fmt.Printf("Correct answer: %t\n", expected == fleets)

	target = 10
	var (
		position2 = []int{4, 1, 0, 7}
		speed2    = []int{2, 2, 1, 1}
	)

	expected = 3

	fleets = carFleet(target, position2, speed2)
	fmt.Printf("Scenario 1, had %d, fleets.\n", fleets)
	fmt.Printf("Correct answer: %t\n", expected == fleets)

	target = 12
	var (
		position3 = []int{10, 8, 0, 5, 3}
		speed3    = []int{2, 4, 1, 1, 3}
	)

	expected = 3

	fleets = carFleet(target, position3, speed3)
	fmt.Printf("Scenario 1, had %d, fleets.\n", fleets)
	fmt.Printf("Correct answer: %t\n", expected == fleets)

	target = 10
	expected = 2
	var (
		position4 = []int{6, 8}
		speed4    = []int{3, 2}
	)

	fleets = carFleet(target, position4, speed4)
	fmt.Printf("Scenario 1, had %d, fleets.\n", fleets)
	fmt.Printf("Correct answer: %t\n", expected == fleets)

}

//need to know for each car, how far until it gets to next car && how long until get to destination
// if time to next car <= time to destination = +1 car fleet

//for each car, need to know
// - time to next car
// - time to final destination

// time to destination (tdd) = target - position / speed
// time to next car (ttn) = (nextCarPos - currPos) / (nextCarSpeed - currSpeed)
// 	- if currPost != nextCarPos && nextCarSpeed >= currSpeed = skip
// 	- if currPos == nextCarPos && nextCarSpeed == currSpeed ==> +1 car fleet

//monotonic stack -- sort by carPosition initially, calculate ttdn and place on stack
// [{carPosition, ttd, ttn}, {carPosition, ttd, ttn}]

//Once all on stack, pop iteratiely and find:
// -- ttn if applicable
//	|-----> will catch up before destination? (fleet?)

//Edge cases:
// 1. Only 1 car

type car struct {
	position int
	speed    int
	ttd      float64
}

type stack []car

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) peek() *car {
	if len(*s) == 0 {
		return nil
	}
	return &(*s)[len(*s)-1]
}

func (s *stack) push(c car) {
	*s = append((*s), c)
}

func (s *stack) pop() car {
	l := len(*s) - 1
	car := (*s)[l]
	*s = (*s)[:l]

	return car
}

func carFleet(target int, position, speed []int) int {
	cars := make([]car, len(position))
	for i, p := range position {
		car := car{
			position: p,
			speed:    speed[i],
			ttd:      (float64)(target-p) / (float64)(speed[i]),
		}
		cars[i] = car
	}

	slices.SortFunc(cars, func(a, b car) int {
		return b.position - a.position
	})

	s := make(stack, 0)
	for _, v := range cars {
		if s.isEmpty() {
			s.push(v)
			continue
		}
		p := s.peek()
		if v.ttd > p.ttd {
			s.push(v)
		}
	}
	return len(s)
}
