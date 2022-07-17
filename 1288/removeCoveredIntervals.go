package removeCoveredIntervals

import "sort"

func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		} else {
			return intervals[i][1] > intervals[j][1]
		}
	})
	pre := intervals[0]
	res := 1
	for _, v := range intervals {
		if v[0] == pre[0] || v[1] <= pre[1] {
			continue
		} else {
			pre = v
			res++
		}
	}
	return res
}
