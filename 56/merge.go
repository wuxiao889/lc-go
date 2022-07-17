package merge

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Sort(A(intervals))
	var res [][]int
	res = append(res, intervals[0])
	for _, v := range intervals {
		pre := res[len(res)-1]
		if v[0] == pre[0] || v[1] < pre[1] {
			continue
		}
		if v[0] > pre[1] {
			res = append(res, v)
		} else {
			res[len(res)-1][1] = v[1]
		}
	}
	return res
}

type A [][]int

func (a A) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a A) Len() int      { return len(a) }
func (a A) Less(i, j int) bool {
	if a[i][0] != a[j][0] {
		return a[i][0] < a[j][0]
	} else {
		return a[i][1] > a[j][1]
	}
}
