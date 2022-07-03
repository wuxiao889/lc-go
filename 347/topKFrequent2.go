package topKFrequent

import "sort"

func topKFrequent2(nums []int, k int) []int {
	ma := make(map[int]int, len(nums))
	for i := range nums {
		ma[nums[i]]++
	}

	keys := make([]int, 0, len(ma))
	for key := range ma {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return ma[keys[i]] > ma[keys[j]]
	})

	return keys[:k]
}
