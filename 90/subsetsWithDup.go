package subsetsWithDup

import (
	"fmt"
	"sort"
)

var res [][]int
var path []int

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res = make([][]int, 0)
	path = make([]int, 0)
	backtrace(nums, 0)
	return res
}

func backtrace(nums []int, index int) {
	fmt.Println(path)
	res = append(res, append([]int{}, path...))
	for i := index; i < len(nums); i++ {
		if i > index && nums[i] == nums[i-1] {
			continue
		}
		path = append(path, nums[i])
		backtrace(nums, i+1)
		path = path[:len(path)-1]
	}
}
