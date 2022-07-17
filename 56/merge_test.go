package merge

import (
	"github.com/wuxiao889/goleetcode/util"
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{util.Str2A2rr("[[1,3],[2,6],[8,10],[15,18]]")},
			want: util.Str2A2rr("[[1,6],[8,10],[15,18]]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
