package sortedArrayToBST

import . "imooc.com/wuxiao/learnGo/lang/tree"

func sortedArrayToBST(nums []int) *TreeNode {
	return binary(nums, 0, len(nums)-1)
}

func binary(nums []int, lo, hi int) *TreeNode {
	if lo > hi {
		return nil
	}
	mid := lo + (hi-lo)>>1
	root := &TreeNode{Value: nums[mid]}
	root.Left = binary(nums, lo, mid-1)
	root.Right = binary(nums, mid+1, hi)
	return root
}
