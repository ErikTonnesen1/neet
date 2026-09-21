package main

import (
	"fmt"
	"slices"
	"strconv"
)

/*
  Given a list of strings, and that they represent a reverse polish notation equation

  Return the evaluation of the expression
*/

func main() {
	input := []string{"1", "2", "+", "3", "*", "4", "-"}
	input2 := []string{"4", "13", "5", "/", "+"}
	ans := evalRPN(input)
	ans2 := evalRPN(input2)
	fmt.Printf("Answer is: %d\n", ans)
	fmt.Printf("Answer is: %d\n", ans2)
}

type stack []string

var operators = map[string]func(x, y int) int{
	"+": func(x, y int) int { return x + y },
	"-": func(x, y int) int { return x - y },
	"/": func(x, y int) int { return x / y },
	"*": func(x, y int) int { return x * y },
}

// My solution :(
func evalRPN(tokens []string) int {
	s := stack{}
	buf := make([]string, 0)

	for _, token := range slices.Backward(tokens) {
		s = append(s, token)
	}

	size := (len(s) - 1)
	for size != -1 {
		//pop item off stack
		v := s[size]
		//resize stack
		s = s[:size]
		size--

		if operation, found := operators[v]; found {
			if len(buf) < 2 {
				panic("operator requires at least 2 operands")
			}
			x, xerr := strconv.Atoi(buf[len(buf)-2])
			y, yerr := strconv.Atoi(buf[len(buf)-1])
			buf = buf[:len(buf)-2]
			if xerr != nil || yerr != nil {
				panic("operands must be of type int")
			}

			eval := operation(x, y)
			s = append(s, strconv.Itoa(eval))
			size++
		} else {
			buf = append(buf, v)
		}
	}

	ans, err := strconv.Atoi(buf[0])
	if err != nil {
		panic("final product is not an int")
	}

	return ans
}

// ChatGPT solution:
func evalRPNChatGPT(tokens []string) int {
	stack := []int{}

	for _, token := range tokens {
		if operation, found := operators[token]; found {
			y := stack[len(stack)-1]
			x := stack[len(stack)-2]

			stack = stack[:len(stack)-2]
			stack = append(stack, operation(x, y))
		} else {
			n, _ := strconv.Atoi(token)
			stack = append(stack, n)
		}
	}

	return stack[0]
}
