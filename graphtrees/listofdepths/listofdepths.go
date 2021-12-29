package main

import (
	"container/list"
	"errors"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func removeIndex(pos int, s []*TreeNode) []*TreeNode {
	copy(s[pos:], s[pos+1:])
	s[len(s)-1] = nil
	s = s[:len(s)-1]
	return s
}
func ListODepths(root *TreeNode) (*list.List, error) {
	if root == nil {
		return nil, errors.New("empty tree")
	}
	items := list.New()
	toVisit := make([]*TreeNode, 0)
	toVisit = append(toVisit, root)
	toVisit = append(toVisit, nil)
	partialList := make([]*TreeNode, 0)
	for len(toVisit) > 0 {
		currentNode := toVisit[0]
		toVisit = removeIndex(0, toVisit)
		if currentNode != nil {
			partialList = append(partialList, currentNode)
			if currentNode.Left != nil {
				toVisit = append(toVisit, currentNode.Left)
			}
			if currentNode.Right != nil {
				toVisit = append(toVisit, currentNode.Right)
			}
		} else {
			items.PushBack(partialList)
			partialList = make([]*TreeNode, 0)
			if len(toVisit) > 0 {
				toVisit = append(toVisit, nil)
			}
		}
	}
	return items, nil
}
