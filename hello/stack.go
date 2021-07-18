package hello

import "fmt"

type StringStack []string

const END int = 1

func (s StringStack) Pop() (value string, newS StringStack, error int) {
	newS = s
	last := len(newS) - 1
	if last < 0 {
		error = END
		return
	}
	value = newS[last]
	newS[last] = ""
	newS = newS[:last]
	return
}

func (s StringStack) Append(value string) StringStack {
	return append(s, value)
}

func (s StringStack) PrintTop() {
	if len(s) < 1 {
		fmt.Println(nil)
		return
	}
	fmt.Println(s[len(s)-1])
}

func NewStringStack() StringStack {
	return StringStack{}
}

func TestStringStack() {
	s := NewStringStack()
	s = s.Append("test")
	s = s.Append("test2")

	s.PrintTop()
	v, s, err := s.Pop()
	fmt.Println(v, s, err)
	s.PrintTop()
	v, s, err = s.Pop()
	fmt.Println(v, s, err)
	s.PrintTop()
	v, s, err = s.Pop()
	fmt.Println(v, s, err)
	if err == END {
		fmt.Println("Reached end")
	}
	s.PrintTop()
}
