package main

import (
	"fmt"
)

func BuildTree(data *[]int, root *TreeNode, lower int, upper int) *TreeNode {
	if upper < lower {
		return nil
	}
	if root == nil {
		pivot := (lower + upper) / 2
		root = new(TreeNode)
		root.Val = (*data)[pivot]
		root.Left = BuildTree(data, root.Left, lower, pivot-1)
		root.Right = BuildTree(data, root.Right, pivot+1, upper)
	}
	return root
}
func InOrderVisit(t *TreeNode) {
	if t != nil {
		InOrderVisit(t.Left)
		fmt.Printf("%v", t.Val)
		InOrderVisit(t.Right)
	}
}
func main() {
	var rootTree *TreeNode = nil
	started := []int{1, 2, 3, 4, 5, 6, 7, 8}
	rootTree = BuildTree(&started, rootTree, 0, len(started)-1)
	fmt.Printf("Visit the tree :")
	InOrderVisit(rootTree)
	fmt.Printf("\n")

	lists, err := ListODepths(rootTree)
	if err != nil {
		panic(err)
	}
	for e := lists.Front(); e != nil; e = e.Next() {
		item := e.Value.([]*TreeNode)
		for _, value := range item {
			fmt.Printf("Value :%d ", value.Val)
		}
		fmt.Println("-")
	}
}
