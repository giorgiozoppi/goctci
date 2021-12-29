package threestacks

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShouldBeAbleToInsertValues(t *testing.T) {
	stack := NewStack()
	stack.Push(5, 0)
	stack.Push(5, 1)
	stack.Push(5, 2)
	item, err := stack.Pop(0)
	require.Nil(t, err)
	require.Equal(t, 5, item)
	item, err = stack.Pop(1)
	require.Nil(t, err)
	require.Equal(t, 5, item)
	item, err = stack.Pop(2)
	require.Nil(t, err)
	require.Equal(t, 5, item)
}
func TestShouldRespectBoundaries(t *testing.T) {
	stack := NewStack()
	stack.Push(5, 1)
	item, err := stack.Pop(1)
	require.Nil(t, err)
	require.Equal(t, 5, item)
	item, err = stack.Pop(1)
	require.Equal(t, "stack undeflow", err.Error())
	require.Equal(t, item, 0)
}
