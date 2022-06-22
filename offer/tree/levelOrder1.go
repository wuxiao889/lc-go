package tree

func levelOrder11(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := make([][]int, 0)
	node := make([]*TreeNode, 0)
	node = append(node, root)
	var curr *TreeNode
	var count int
	for len(node) != 0 {
		arr := make([]int, 0)
		for i, len := 0, len(node); i < len; i++ {
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
		if count%2 == 1 {
			for left, right := 0, len(arr)-1; left < right; left, right = left+1, right-1 {
				arr[left], arr[right] = arr[right], arr[left]
			}
		}
		res = append(res, arr)
		count++
	}
	return res
}
