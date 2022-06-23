package combinationSum

var res [][]int
var path []int
var sum int

func combinationSum(candidates []int, target int) [][]int {
	res = make([][]int, 0)
	path = make([]int, 0)
	backtrace(candidates, target, 0)
	return res
}

func backtrace(candidates []int, target int, start int) {
	if sum == target {
		res = append(res, append([]int{}, path...))
		return
	}
	if sum > target {
		return
	}
	for i := start; i < len(candidates); i++ {
		path = append(path, candidates[i])
		sum += candidates[i]
		backtrace(candidates, target, i)
		path = path[:len(path)-1]
		sum -= candidates[i]
	}
	return
}
