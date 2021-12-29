package routegraph

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShouldHasRoute(t *testing.T) {
	g := NewGraph()
	c := NewNode("C")
	k := NewNode("K")
	k.Children = append(k.Children, NewNode("E"))
	c.Children = append(c.Children, k)
	c.Children = append(c.Children, NewNode("F"))
	l := NewNode("L")
	n := NewNode("S")
	n.Children = append(n.Children, NewNode("A"))
	n.Children = append(n.Children, l)
	n.Children = append(n.Children, c)
	g.AddNode(n)
	require.Equal(t, true, HasPath(c, NewNode("E")))
}
func TestShouldHasNotRoute(t *testing.T) {
	g := NewGraph()
	c := NewNode("C")
	k := NewNode("K")
	k.Children = append(k.Children, NewNode("E"))
	c.Children = append(c.Children, k)
	c.Children = append(c.Children, NewNode("F"))
	l := NewNode("L")
	n := NewNode("S")
	n.Children = append(n.Children, NewNode("A"))
	n.Children = append(n.Children, l)
	n.Children = append(n.Children, c)
	g.AddNode(n)
	require.Equal(t, false, HasPath(c, NewNode("Z")))
}
