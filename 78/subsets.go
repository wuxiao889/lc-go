package subsets

var res [][]int
var path []int

func subsets(nums []int) [][]int {
	res = make([][]int, 0)
	path = make([]int, 0)
	backtrace(0, nums)
	return res
}

func backtrace(pos int, nums []int) {
	res = append(res, append([]int{}, path...))
	for i := pos; i < len(nums); i++ {
		path = append(path, nums[i])
		backtrace(i+1, nums)
		path = path[:len(path)-1]
	}
}
