package stackmin

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShouldPushAndPopSorted(t *testing.T) {
	stack := NewStackMin()
	stack.Push(1)
	stack.Push(8)
	stack.Push(10)
	item := stack.Min()
	value := item.(int)
	require.Equal(t, 1, value)
}
