package tree

// https://leetcode.cn/problems/dui-cheng-de-er-cha-shu-lcof/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return recur(root.Left, root.Right)
}

func recur(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil || right.Val != left.Val {
		return false
	}
	return recur(left.Left, right.Right) && recur(left.Right, right.Left)
}
