package pilestack

import "errors"

const (
	MaxThreaSold = 1024
)

type SetOfStacks struct {
	buckets []*Stack
	top     int
}

func NewSetOfStacks() *SetOfStacks {
	return &SetOfStacks{
		buckets: make([]*Stack, 1, 10),
		top:     -1,
	}
}
func (s *SetOfStacks) Push(item interface{}) {
	if s.top < 0 {
		current := new(Stack)
		current.Push(item)
		s.buckets = append(s.buckets, current)
	} else {
		index := s.top / MaxThreaSold
		current := s.buckets[index]
		if current != nil {
			if current.Len() < MaxThreaSold {
				s.buckets[index].Push(item)
			} else {
				newStack := NewStack()
				newStack.Push(item)
				s.buckets = append(s.buckets, newStack)
			}
		} else {
			s.buckets[index] = new(Stack)
			s.buckets[index].Push(item)
		}
	}
	s.top++
}
func (s *SetOfStacks) Pop() (interface{}, error) {
	if s.top < 0 {
		return nil, errors.New("stack underflow")
	}
	index := s.top / MaxThreaSold
	current := s.buckets[index]
	retValue := current.Pop()
	return retValue, nil
}
