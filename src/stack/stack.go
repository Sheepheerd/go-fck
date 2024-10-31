package stack

import (
	"container/list"
)

type Stack struct {
	list *list.List
	size int
}

func New() *Stack {
	return &Stack{
		list: list.New(),
		size: 0,
	}
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Push(x interface{}) {
	s.list.PushBack(x)
	s.size++
}

func (s *Stack) Pop() interface{} {
	if s.list.Len() == 0 {
		return nil
	}
	tail := s.list.Back()
	val := tail.Value
	s.list.Remove(tail)
	s.size--
	return val
}
