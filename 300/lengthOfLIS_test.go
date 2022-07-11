package lengthOfLIS

import (
	"github.com/wuxiao889/goleetcode/util"
	"testing"
)

func Test_lengthOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{util.Str2Arr("[10,9,2,5,3,7,101,18]")},
			want: 4,
		},
		{
			name: "2",
			args: args{util.Str2Arr("[0,1,0,3,2,3]")},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
