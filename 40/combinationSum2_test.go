package combinationSum2

import (
	"fmt"
	"github.com/wuxiao889/goleetcode/util"
	"testing"
)

var tests = []struct {
	candidates []int
	target     int
}{
	{
		candidates: util.Str2Arr("[1,1,2,5,6,7,10]"),
		target:     8,
	},
	{
		candidates: util.Str2Arr("[3,1,3,5,1,1]"),
		target:     8,
	},
}

func TestXXX(t *testing.T) {
	for _, tt := range tests {
		fmt.Println(combinationSum2(tt.candidates, tt.target))
		fmt.Println("-----")
	}
}
