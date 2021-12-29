package minimaltree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TreeShouldCreatedSameWeight(t *testing.T) {
	var rootTree *TreeNode = nil
	started := []int{1, 3, 4, 7, 10, 12, 13, 15, 20, 23}
	visited := make([]int, 0)
	rootTree = BuildTree(&started, rootTree, 0, len(started)-1)
	InOrderVisit(rootTree, &visited)
	require.Equal(t, started, visited)
}
