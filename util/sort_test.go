package util

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func Test_mergeSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		n := rand.Intn(20)
		nums := make([]int, n)
		for i := range nums {
			nums[i] = rand.Intn(20)
		}
		nums2 := make([]int, n)
		copy(nums2, nums)
		sort.Ints(nums2)
		heapSort(nums)
		//mergeSort(nums, 0, len(nums)-1)
		if !reflect.DeepEqual(nums, nums2) {
			t.Logf("failed to sort:\ngot:%v\nexpected:%v", nums, nums2)
		}
		//quickSort(nums, 0, len(nums)-1)
	}
}
