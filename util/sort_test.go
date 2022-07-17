package util

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Test_quickSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		n := rand.Intn(20)
		nums := make([]int, n)
		for i := range nums {
			nums[i] = rand.Intn(20)
		}
		fmt.Println(nums)
		quickSort(nums, 0, len(nums)-1)
		fmt.Println(nums)
	}
}
