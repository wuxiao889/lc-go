package isBalanced

import . "imooc.com/wuxiao/learnGo/lang/tree"

func isBalanced(root *TreeNode) bool {
	return maxDepth(root) != -1
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	if left == -1 {
		return -1
	}
	right := maxDepth(root.Right)
	if right == -1 {
		return -1
	}
	if right-left < 2 && left-right < 2 {
		return 1 + max(right, left)
	} else {
		return -1
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
