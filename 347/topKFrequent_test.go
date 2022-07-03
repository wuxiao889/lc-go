package topKFrequent

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestTopKFrequent(t *testing.T) {
	arrs := make([][]int, 1)
	rand.Seed(time.Now().Unix())
	for i := range arrs {
		lens := rand.Intn(20) + 2
		arrs[i] = make([]int, lens)
		for j := range arrs[i] {
			arrs[i][j] = rand.Intn(5)
		}
	}
	fmt.Println(arrs)
	for i := range arrs {
		fmt.Println(topKFrequent3(arrs[i], 2))
	}
}
