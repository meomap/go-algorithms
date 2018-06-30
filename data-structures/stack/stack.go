// Package stack provides Stack implementation
package stack

import "fmt"

type Stack struct {
	Top   uint
	Items []interface{}
}

func New() *Stack {
	return &Stack{}
}

func (s Stack) String() string {
	return fmt.Sprintf("top=%d last_item=%#v", s.Top, s.Items[s.Top-1])
}

func (s *Stack) Push(el interface{}) {
	s.Top++
	s.Items = append(s.Items, el)
}

func (s *Stack) Pop() interface{} {
	s.Top--
	return s.Items[s.Top]
}
