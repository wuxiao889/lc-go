package kSmallestPairs

import (
	"github.com/wuxiao889/goleetcode/util"
	"reflect"
	"testing"
)

func Test_kSmallestPairs(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test1",
			args: args{
				nums1: util.Str2Arr("[1,1,2]"),
				nums2: util.Str2Arr("[1,2,3]"),
				k:     2,
			},
			want: util.Str2A2rr("[1,1],[1,1]"),
		}, {
			name: "test2",
			args: args{
				nums1: util.Str2Arr("[1,7,11]"),
				nums2: util.Str2Arr("[2,4,6]"),
				k:     3,
			},
			want: util.Str2A2rr("[1,2],[1,4],[1,6]"),
		},
		{
			name: "test3",
			args: args{
				nums1: util.Str2Arr("[1,2]"),
				nums2: util.Str2Arr("[3]"),
				k:     3,
			},
			want: util.Str2A2rr("[[1 3] [2 3]]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kSmallestPairs(tt.args.nums1, tt.args.nums2, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kSmallestPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
