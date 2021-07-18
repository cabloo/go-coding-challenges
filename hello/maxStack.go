package hello

import "fmt"

type Stack struct {
	values []int
	maxes  []int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(values ...int) {
	s.values = append(s.values, values...)
	for _, value := range values {
		if max, _ := s.Max(); max <= value {
			s.maxes = append(s.maxes, value)
		}
	}
}

func (s *Stack) Pop() (value int, ok bool) {
	if len(s.values) == 0 {
		return
	}

	last := len(s.values) - 1
	value = s.values[last]
	if max, _ := s.Max(); max == value {
		s.maxes = s.maxes[:len(s.maxes)-1]
	}
	s.values = s.values[:last]
	ok = true
	return
}

func (s *Stack) Max() (value int, ok bool) {
	if len(s.maxes) == 0 {
		return
	}

	value = s.maxes[len(s.maxes)-1]
	ok = true
	return
}

func TestStack() {
	s := NewStack()
	s.Push(1, 2, 4, 4, 2, 5)
	for range s.values {
		vals := s.values
		max, maxOk := s.Max()
		val, valOk := s.Pop()
		fmt.Println(vals, max, val, maxOk, valOk)
	}
	max, maxOk := s.Max()
	val, valOk := s.Pop()
	fmt.Println(s.values, max, val, maxOk, valOk)
}
