package topKFrequent

import "fmt"

func heapify(arr []int) {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		fmt.Println(arr)
		down(arr, i, n)
	}
}

func down(arr []int, i int, n int) {
	for {
		j1 := 2*i + 1
		if j1 >= n {
			return
		}
		j := j1
		if j2 := 2*i + 2; j2 < n && ma[arr[j2]] > ma[arr[j1]] {
			j = j2
		}
		if ma[arr[j]] <= ma[arr[i]] {
			return
		}
		arr[i], arr[j] = arr[j], arr[i]
		i = j
	}
}

var ma map[int]int

func topKFrequent3(nums []int, k int) []int {
	ma = make(map[int]int, len(nums))
	for i := range nums {
		ma[nums[i]]++
	}
	fmt.Println(ma)
	keys := make([]int, 0, len(ma))

	for key := range ma {
		keys = append(keys, key)
	}
	fmt.Println(keys)
	heapify(keys)
	fmt.Println(keys)

	var res []int
	n := len(keys)
	for i := n - 1; i >= n-k; i-- {
		res = append(res, keys[0])
		keys[0], keys[i] = keys[i], keys[0]
		down(keys, 0, i)
	}
	return res
}
