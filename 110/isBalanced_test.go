package isBalanced

import (
	"github.com/wuxiao889/goleetcode/util"
	. "imooc.com/wuxiao/learnGo/lang/tree"
	"math"
	"testing"
)

const null = math.MaxInt

func Test_isBalanced(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{NewTree(util.Str2Arr(" [1,2,2,3,3,null,null,4,4]"))},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.args.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
