package combinationSum2

import (
	"sort"
)

var (
	sum  int
	res  [][]int
	path []int
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	sum = 0
	res = make([][]int, 0)
	path = make([]int, 0)
	backtrace(candidates, target, 0)
	return res
}

func backtrace(candidates []int, target, k int) {
	if sum == target {
		res = append(res, append([]int{}, path...))
		return
	}
	for i := k; i < len(candidates) && candidates[i]+sum <= target; i++ {
		if i > k && candidates[i] == candidates[i-1] {
			continue
		}
		path = append(path, candidates[i])
		sum += candidates[i]
		backtrace(candidates, target, i+1)
		path = path[:len(path)-1]
		sum -= candidates[i]
	}
}
