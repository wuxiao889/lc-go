package pathSum

import (
	. "imooc.com/wuxiao/learnGo/lang/tree"
)

var res [][]int
var path []int

func pathSum(root *TreeNode, target int) [][]int {
	res = make([][]int, 0)
	path = make([]int, 0)
	recur(root, target)
	return res
}

func recur(root *TreeNode, n int) {
	if root == nil {
		return
	}
	path = append(path, root.Value)
	n -= root.Value
	if n == 0 && root.Left == nil && root.Right == nil {
		p := make([]int, len(path))
		copy(p, path)
		res = append(res, p)
	}
	recur(root.Left, n)
	recur(root.Right, n)
	path = path[:len(path)-1]
}
