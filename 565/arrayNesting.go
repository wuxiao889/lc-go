package arrayNesting

func arrayNesting(nums []int) (ans int) {
	n := len(nums)
	for i := range nums {
		cnt := 0
		for nums[i] < n {
			i, nums[i] = nums[i], n
			cnt++
		}
		if cnt > ans {
			ans = cnt
		}
	}
	return
}

//var maxl int
//var curl int
//var used []bool
//var mem []int
//var path []int
//
//func arrayNesting(nums []int) int {
//	maxl = 0
//	curl = 0
//	used = make([]bool, len(nums))
//	mem = make([]int, len(nums))
//	for i := range mem {
//		mem[i] = -1
//	}
//	path = make([]int, 0)
//	for i := range nums {
//		if mem[i] == -1 {
//			backtrace(nums, i, i)
//		}
//	}
//	return maxl
//}
//
//func backtrace(nums []int, k, j int) {
//	if used[j] {
//		maxl = max(maxl, curl)
//		for i := range path {
//			mem[path[i]] = len(path) - i
//		}
//		return
//	}
//	if nums[j] < len(nums) {
//		if mem[nums[j]] != -1 {
//			mem[j] = mem[nums[j]] + 1
//			return
//		}
//		path = append(path, j)
//		curl++
//		used[j] = true
//		backtrace(nums, k, nums[j])
//		used[j] = false
//		curl--
//		path = path[:len(path)-1]
//	}
//}
//
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
