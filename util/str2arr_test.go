package util

import (
	"reflect"
	"testing"
)

func TestStrToArr(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{"[]"},
			want: []int{},
		},
		{
			name: "2",
			args: args{"[2]"},
			want: []int{2},
		},
		{
			name: "3",
			args: args{"[1,2,3]"},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Str2Arr(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Str2Arr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrToA2rr(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "0",
			args: args{"[]"},
			want: [][]int{},
		},

		{
			name: "",
			args: args{"[[1]]"},
			want: [][]int{{1}},
		},
		{
			name: "2",
			args: args{"[[1,1,1,0,0],[0,1,1,1,1],[0,0,0,0,0],[1,0,0,0,0],[1,1,0,1,1]]"},
			want: [][]int{{1, 1, 1, 0, 0}, {0, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 1, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Str2A2rr(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Str2A2rr() = %v, want %v", got, tt.want)
			}
		})
	}
}
