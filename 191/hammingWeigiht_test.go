package hammingWeight

import (
	"fmt"
	"testing"
)

func TestXX(t *testing.T) {
	var nums [10]int
	for i := range nums {
		fmt.Printf("%d %b\n", i, i)
		fmt.Println(hammingWeight(uint32(i)))
	}
}
