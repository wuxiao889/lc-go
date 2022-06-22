package convertBST

import (
	. "imooc.com/wuxiao/learnGo/lang/tree"
)

var sum int

func convertBST(root *TreeNode) *TreeNode {
	sum = 0
	inorder(root)
	return root
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Right)
	sum += root.Value
	root.Value = sum
	inorder(root.Left)
}
