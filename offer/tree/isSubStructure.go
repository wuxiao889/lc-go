package tree

// https://leetcode.cn/problems/shu-de-zi-jie-gou-lcof/
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	return isContain(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func isContain(A, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil || A.Val != B.Val {
		return false
	}
	return isContain(A.Left, B.Left) && isContain(A.Right, B.Right)
}
