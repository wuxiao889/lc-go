package grayCode

func grayCode(n int) []int {
	res := []int{0, 1}
	add := 2
	for i := 1; i < n; i++ {
		n := len(res)
		for j := n - 1; j >= 0; j-- {
			res = append(res, res[j]+add)
		}
		add *= 2
	}
	return res
}
