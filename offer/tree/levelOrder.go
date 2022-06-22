package tree

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	arr := []int{}
	l := list.New()
	l.PushBack(root)
	var curr *TreeNode
	for l.Front() != nil {
		curr = l.Front().Value.(*TreeNode)
		l.Remove(l.Front())
		if curr.Left != nil {
			l.PushBack(curr.Left)
		}
		if curr.Right != nil {
			l.PushBack(curr.Right)
		}
		arr = append(arr, curr.Val)
	}
	return arr
}

func levelOrder1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	arr := make([]int, 0)
	node := make([]*TreeNode, 0)
	node = append(node, root)
	var curr *TreeNode
	for len(node) != 0 {
		curr = node[0]
		node = node[1:]
		if curr.Left != nil {
			node = append(node, curr.Left)
		}
		if curr.Right != nil {
			node = append(node, curr.Right)
		}
		arr = append(arr, curr.Val)
	}
	return arr
}
