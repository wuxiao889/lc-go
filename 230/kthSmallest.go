package kthSmallest

import (
	"fmt"
	. "imooc.com/wuxiao/learnGo/lang/tree"
)

func kthSmallest(root *TreeNode, k int) int {
	var c, res int
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		fmt.Println(root.Value)
		if c++; c == k {
			res = root.Value
			return
		}
		inorder(root.Right)
	}
	inorder(root)
	return res
}
