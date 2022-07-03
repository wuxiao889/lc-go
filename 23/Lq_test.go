package mergeKLists

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestLq(t *testing.T) {
	lq := Lq{}
	for i := 0; i < 10; i++ {
		lq.push(&ListNode{Val: rand.Intn(5)})
	}
	for len(lq) > 0 {
		fmt.Printf("%v ", lq.pop().Val)
	}
}
