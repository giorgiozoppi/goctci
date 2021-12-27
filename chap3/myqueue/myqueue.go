package myqueue

// instead using a custom stack we'll use container/list as a stack

import (
	"container/list"
	"errors"
)

type Stack struct {
	data *list.List
	top  *list.Element
}

func NewStack() *Stack {
	return &Stack{
		data: list.New(),
		top:  nil,
	}
}
func (s *Stack) Push(value interface{}) {
	s.data.PushBack(value)
	s.top = s.data.Back()
}
func (s *Stack) Empty() bool {
	return s.data.Len() == 0
}
func (s *Stack) Peek() interface{} {
	item := s.top.Value
	return item
}
func (s *Stack) Pop() (interface{}, error) {
	if s.top == nil {
		return nil, errors.New("stack undeflow")
	}
	value := s.data.Back()
	data := value.Value
	s.data.Remove(value)
	return data, nil
}

type MyQueue struct {
	data *Stack
	tmp  *Stack
}

func NewMyQueue() *MyQueue {
	return &MyQueue{
		data: NewStack(),
		tmp:  NewStack(),
	}
}
func (m *MyQueue) Offer(item interface{}) {
	m.data.Push(item)
}
func (m *MyQueue) Take() interface{} {
	for !m.data.Empty() {
		top, _ := m.data.Pop()
		m.tmp.Push(top)
	}
	retValue := m.tmp.Peek()
	for !m.tmp.Empty() {
		top, _ := m.tmp.Pop()
		m.data.Push(top)
	}
	return retValue
}
