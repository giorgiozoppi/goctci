package sortedstack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type IntComparable struct {
	Value int
}

func (in IntComparable) LessEqual(i interface{}) bool {
	a, b := in.Value, i.(IntComparable).Value
	return a <= b
}
func TestShouldPushAndPopCorrectly(t *testing.T) {
	stack := NewSortedStack()
	stack.Push(IntComparable{
		Value: 7})
	stack.Push(IntComparable{
		Value: 4})

	stack.Push(IntComparable{
		Value: 2})
	stack.Push(IntComparable{
		Value: 10})
	item, _ := stack.Pop()
	value := item.(IntComparable).Value
	require.Equal(t, 2, value)
}
