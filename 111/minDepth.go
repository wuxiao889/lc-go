package minDepth

import (
	. "imooc.com/wuxiao/learnGo/lang/tree"
)

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	height := 1
	for len(q) > 0 {
		l := len(q)
		for i := 0; i < l; i++ {
			cur := q[0]
			q = q[1:]
			if cur.Left == nil && cur.Right == nil {
				return height
			}
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		height++
	}
	return 0
}
