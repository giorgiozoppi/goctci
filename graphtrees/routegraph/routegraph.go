package routegraph

type Node struct {
	Name     string
	Children []*Node
}
type Graph struct {
	Nodes []*Node
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make([]*Node, 0),
	}
}
func (g *Graph) AddNode(n *Node) {
	g.Nodes = append(g.Nodes, n)
}
func NewNode(name string) *Node {
	return &Node{
		Name:     name,
		Children: make([]*Node, 0),
	}
}
func HasPath(source *Node, dest *Node) bool {
	visited := make([]*Node, 0, 10)
	pathFound := false
	visited = append(visited, source)
	for !pathFound && len(visited) > 0 {
		currentNode := visited[len(visited)-1]
		if currentNode.Name == dest.Name {
			pathFound = true
			break
		} else {
			visited = visited[:len(visited)-1]
			for _, child := range currentNode.Children {
				visited = append(visited, child)
			}
		}

	}
	return pathFound
}
