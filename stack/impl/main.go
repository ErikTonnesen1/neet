package main

import (
	"errors"
	"fmt"
	"stack-impl/stack"
	"strings"
)

func main() {
	s := stack.New[string](10)

	s.Push("Henry!")
	s.Push("On")
	s.Push("Going")
	s.Push("What's'")
	s.Push("Hello,")

	var message strings.Builder
	for !s.IsEmpty() {
		m, err := s.Pop()

		if err != nil {
			panic(err)
		}

		if strings.HasPrefix(m, "H") {
			message.WriteString(m)
			message.WriteString(" ")
		}
	}

	_, err := s.Top()
	if err != nil {
		if errors.Is(err, stack.ErrStackUnderflow) {
			fmt.Println("Stack fully processed")
		}
	}

	fmt.Printf("Printing message: %s", message.String())

}
