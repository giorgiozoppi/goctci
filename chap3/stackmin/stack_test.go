package stackmin

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShouldPushAndPopCorrectly(t *testing.T) {
	stack := NewStack()
	stack.Push(5)
	stack.Push(3)
	stack.Push(1)
	item := stack.Pop()
	value := item.(int)
	require.Equal(t, 1, value)
}
func TestShouldCheckIfEmptyIsValid(t *testing.T) {
	stack := NewStack()
	stack.Push(4)
	stack.Push(8)
	for !stack.IsEmpty() {
		stack.Pop()
	}
	require.Equal(t, true, stack.IsEmpty())
}
