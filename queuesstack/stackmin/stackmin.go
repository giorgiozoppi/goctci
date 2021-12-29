package stackmin

type StackMin struct {
	values  *Stack
	minimus *Stack
}

func NewStackMin() *StackMin {
	return &StackMin{
		values:  NewStack(),
		minimus: NewStack(),
	}
}
func (s *StackMin) Top() interface{} {
	return s.values.Peek()
}
func (s *StackMin) Min() interface{} {
	return s.minimus.Peek()
}
func (s *StackMin) Pop() interface{} {
	value := s.values.Peek()
	if !s.minimus.IsEmpty() {
		min := s.minimus.Peek()
		if value == min {
			s.minimus.Pop()
		}
	}
	s.values.Pop()
	return value
}
func (s *StackMin) Push(item interface{}) {
	value := item.(int)
	min := value
	if s.minimus.IsEmpty() {
		s.minimus.Push(item)
		min = value
	} else {
		min = s.minimus.Peek().(int)
	}
	if value < min {
		s.minimus.Push(value)
	}
	s.values.Push(value)
}
