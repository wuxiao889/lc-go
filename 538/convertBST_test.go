package convertBST

import (
	. "imooc.com/wuxiao/learnGo/lang/tree"
	"testing"
)

var tests = []struct {
	r1, r2 *TreeNode
}{
	{
		r1: NewTree([]int{4, 1, 6, 0, 2, 5, 7, Null, Null, Null, 3, Null, Null, Null, 8}),
		r2: NewTree([]int{30, 36, 21, 36, 35, 26, 15, Null, Null, Null, 33, Null, Null, Null, 8}),
	},
}

func TestConvertBst(t *testing.T) {
	for _, tt := range tests {
		if ans := convertBST(tt.r1); !IsSameTree(ans, tt.r2) {
			t.Errorf("error")
		}
	}
}
