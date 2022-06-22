package tree

// https://leetcode.cn/problems/er-cha-shu-de-jing-xiang-lcof/
func mirrorTree(root *TreeNode) *TreeNode {
	stack := make([]*TreeNode, 0)
	if root == nil {
		return nil
	}
	stack = append(stack, root)
	var curr *TreeNode
	for len(stack) > 0 {
		curr = stack[0]
		stack = stack[1:]
		if curr == nil {
			continue
		}
		stack = append(stack, curr.Left)
		stack = append(stack, curr.Right)
		curr.Left, curr.Right = curr.Right, curr.Left
	}
	return root
}
