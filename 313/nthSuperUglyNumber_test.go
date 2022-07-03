package nthSuperUglyNumber

import (
	"github.com/wuxiao889/goleetcode/util"
	"testing"
)

func Test_nthSuperUglyNumber(t *testing.T) {
	type args struct {
		n      int
		primes []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				n:      12,
				primes: util.StrToArr("[2,7,13,19]"),
			},
			want: 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nthSuperUglyNumber(tt.args.n, tt.args.primes); got != tt.want {
				t.Errorf("nthSuperUglyNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
