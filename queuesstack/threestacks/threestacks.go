package threestacks

import "errors"

const (
	FirstStack = iota
	SecondStack
	ThirdStack
	StackSize = 2048
	MaxSize   = StackSize * 3
)

type Stack struct {
	arena     []int
	firstTop  int
	secondTop int
	thirdTop  int
}

func NewStack() *Stack {
	return &Stack{
		arena:     make([]int, MaxSize, MaxSize),
		firstTop:  -1,
		secondTop: StackSize - 1,
		thirdTop:  (StackSize * 2) - 1,
	}
}
func (s *Stack) Pop(index int) (int, error) {
	value, err := 0, errors.New("invalid input")
	switch index {
	case FirstStack:
		{
			value, err = extractItem(s, &s.firstTop, 0)
		}
	case SecondStack:
		{
			value, err = extractItem(s, &s.secondTop, StackSize)
		}
	case ThirdStack:
		{
			value, err = extractItem(s, &s.thirdTop, StackSize*2)
		}
	default:
		{
			value, err = 0, errors.New("invalid index")
		}
	}
	return value, err
}
func (s *Stack) Push(item int, index int) error {
	err := errors.New("invalid input")
	switch index {
	case FirstStack:
		{
			err = pushItem(s, &s.firstTop, StackSize, item)
		}
	case SecondStack:
		{
			err = pushItem(s, &s.secondTop, StackSize*2, item)
		}
	case ThirdStack:
		{
			err = pushItem(s, &s.thirdTop, StackSize*3, item)
		}
	default:
		{
			err = errors.New("invalid index")
		}
	}
	return err
}
func extractItem(s *Stack, top *int, startIndex int) (int, error) {
	if *top < startIndex {
		return 0, errors.New("stack undeflow")
	}
	data := s.arena[*top]
	*top--
	return data, nil
}
func pushItem(s *Stack, top *int, endIndex int, item int) error {
	if *top >= endIndex-1 {
		return errors.New("stack overflow")
	}
	*top++
	s.arena[*top] = item
	return nil
}
