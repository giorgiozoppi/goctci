package minimaltree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(data *[]int, root *TreeNode, lower int, upper int) *TreeNode {
	offset := upper - lower
	if offset <= 0 {
		return root
	}
	if root == nil {
		pivot := lower + (offset / 2)
		root = new(TreeNode)
		root.Val = (*data)[pivot]
		root.Left = BuildTree(data, root.Left, lower, pivot)
		root.Right = BuildTree(data, root.Right, pivot+1, upper)
	}
	return root
}
func InOrderVisit(t *TreeNode, visited *[]int) {
	if t != nil {
		InOrderVisit(t.Left, visited)
		*visited = append(*visited, t.Val)
		InOrderVisit(t.Right, visited)
	}
}
