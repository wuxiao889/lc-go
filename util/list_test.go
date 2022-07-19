package util

import (
	"reflect"
	"testing"
)

func TestNewListNode(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "empty slice",
			args: args{[]int{}},
			want: []int{},
		},
		{
			name: "1",
			args: args{[]int{1, 2, 3}},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if nums := ListToSlice(Arr2List(tt.args.nums)); !reflect.DeepEqual(nums, tt.want) {
				t.Errorf("got %v, want %v", nums, tt.want)
			}
		})
	}
}
