package countPrimes

import (
	"fmt"
	"testing"
)

func TestXXX(t *testing.T) {
	fmt.Println(countPrimes(10))
	fmt.Println(countPrimes(1))
	fmt.Println(countPrimes(5))
	//fmt.Println(countPrimes(100))
}

func Test_countPrimes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPrimes(tt.args.n); got != tt.want {
				t.Errorf("countPrimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
