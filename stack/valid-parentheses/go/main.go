package main

import (
	"fmt"
	"strings"
)

func main() {
	inputs := []string{"()", "{}", "[]", "{([])}", "]", "]]", "({}})"}

	for _, v := range inputs {
		fmt.Printf("Is %s, a valid parenthesis pair? --> %t\n", v, isValid(v))
	}
}

func isValid(s string) bool {

	parenthesisPairs := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}

	stack := Stack{}

	for v := range strings.SplitSeq(s, "") {

		_, openCharacterFound := parenthesisPairs[v]

		if openCharacterFound {
			stack.Push(v)
		} else {
			lastOpenCharacter, ok := stack.Pop()
			if !ok {
				return false
			}
			requiredMatch := parenthesisPairs[lastOpenCharacter]

			if !(requiredMatch == v) {
				return false
			}
		}
	}

	if stack.Length() > 0 {
		return false
	}
	return true
}

type Stack struct {
	Stack []string
}

func (s *Stack) Push(p string) {
	s.Stack = append(s.Stack, p)
}

func (s *Stack) Pop() (string, bool) {
	if len(s.Stack) == 0 {
		return "", false
	}
	pop := s.Stack[len(s.Stack)-1]
	s.Stack = s.Stack[:len(s.Stack)-1]
	return pop, true
}

func (s *Stack) Length() int {
	return len(s.Stack)
}
