package minDepth

import (
	. "imooc.com/wuxiao/learnGo/lang/tree"
	"testing"
)

func TestXXX(t *testing.T) {
	var tests = []struct {
		t     *TreeNode
		depth int
	}{{
		NewTree([]int{1, 2, 3, 4, 5}),
		2,
	},
	}
	for _, tt := range tests {
		if res := minDepth(tt.t); res != tt.depth {
			t.Errorf("Expected %d , Got %d\n", tt.depth, res)
		}
	}
}
