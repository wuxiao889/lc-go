package combine

var res [][]int
var path []int
var nn, kk int

func combine(n int, k int) [][]int {
	nn, kk = n, k
	res = make([][]int, 0)
	path = make([]int, 0)
	backtrace(1)
	return res
}

func backtrace(index int) {
	if len(path) == kk {
		res = append(res, append([]int{}, path...))
		return
	}
	for i := index; i <= nn; i++ {
		path = append(path, i)
		backtrace(i + 1)
		path = path[:len(path)-1]
	}
}
