package stack

import "errors"

var ErrStackOverflow error = errors.New("stack capcity full. cannot continue operation")
var ErrStackUnderflow error = errors.New("stack is empty. cannot continue operation")

type Stack[T any] struct {
	capacity int
	stack    []T
}

func New[T any](capacity int) *Stack[T] {
	if capacity < 0 {
		panic("capacity of stack must be positive")
	}
	return &Stack[T]{
		capacity: capacity,
		stack:    make([]T, 0, capacity),
	}
}

func (s *Stack[T]) Push(p T) error {
	if len(s.stack) >= s.capacity {
		return ErrStackOverflow
	}

	s.stack = append(s.stack, p)
	return nil
}

func (s *Stack[T]) Pop() (T, error) {
	var zeroValue T
	if s.IsEmpty() {
		return zeroValue, ErrStackUnderflow
	}

	//get index of last item
	last := len(s.stack) - 1

	//retrieve last item to return to caller
	pop := s.stack[last]

	//if stack contains pointers, empty value to avoid retaining references
	s.stack[last] = zeroValue

	//shorten stack by 1 after popping
	s.stack = s.stack[:last]

	return pop, nil
}

func (s *Stack[T]) Top() (T, error) {
	if len(s.stack) == 0 {
		var zero T
		return zero, ErrStackUnderflow
	}

	//return last item in stack without removing
	return s.stack[len(s.stack)-1], nil
}

func (s *Stack[T]) Size() int {
	return len(s.stack)
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}
