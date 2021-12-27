package sortedstack

// instead using a custom stack we'll use container/list as a stack

import (
	"container/list"
	"errors"
)

type Comparable interface {
	LessEqual(i interface{}) bool
}

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
func (s *Stack) Top() interface{} {
	if s.top != nil {
		item := s.top.Value
		return item
	}
	return nil
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

type SortedStack struct {
	sorted *Stack
	tmp    *Stack
}

func NewSortedStack() *SortedStack {
	return &SortedStack{
		sorted: NewStack(),
		tmp:    NewStack(),
	}
}
func (m *SortedStack) IsEmpty() bool {
	return m.sorted.Empty()
}

func (m *SortedStack) Peek() Comparable {
	return m.sorted.Top().(Comparable)
}
func (m *SortedStack) Pop() (Comparable, error) {
	value, err := m.sorted.Pop()
	return value.(Comparable), err
}
func (m *SortedStack) Push(item Comparable) {
	top := m.sorted.Top()
	if top != nil {
		if item.LessEqual(top) {
			m.sorted.Push(item)
		} else {
			for !m.sorted.Empty() {
				tmp, _ := m.sorted.Pop()
				m.tmp.Push(tmp)
			}
		}
	}
	m.tmp.Push(item)
	for !m.tmp.Empty() {
		tmp, _ := m.tmp.Pop()
		m.sorted.Push(tmp)
	}
}
