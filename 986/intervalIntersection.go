package intervalIntersection

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	var res [][]int
	for i, j := 0, 0; i < len(firstList) && j < len(secondList); {
		f, s := firstList[i], secondList[j]
		left, right := max(f[0], s[0]), min(f[1], s[1])
		if left <= right {
			res = append(res, []int{left, right})
		}
		if f[1] > s[1] {
			j++
		} else {
			i++
		}
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
