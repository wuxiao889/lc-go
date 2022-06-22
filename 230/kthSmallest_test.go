package kthSmallest

import (
	. "imooc.com/wuxiao/learnGo/lang/tree"
	"testing"
)

var tests = []struct {
	root *TreeNode
	k    int
	val  int
}{
	{
		NewBST([]int{0, 10, 3, 1, 2, 5, 9}),
		1,
		0,
	},
	{
		NewBST([]int{0, 10, 3, 1, 2, 5, 9}),
		4,
		3,
	},
}

func TestKthSmallest(t *testing.T) {
	for _, tt := range tests {
		ans := kthSmallest(tt.root, tt.k)
		if ans != tt.val {
			t.Errorf("expected %d got %d\n", ans, tt.val)
		}
	}
}
