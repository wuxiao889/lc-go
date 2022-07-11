package kSmallestPairs

import "container/heap"

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	n, m, res := len(nums1), len(nums2), make([][]int, 0, k)
	flag := n > m
	if flag {
		n, m, nums1, nums2 = m, n, nums2, nums1
	}
	if n > k {
		n = k
	}
	h := make(hp, 0, n)
	for i := 0; i < n; i++ {
		h = append(h, []int{nums1[i] + nums2[0], i, 0})
	}
	heap.Init(&h)

	for h.Len() > 0 && len(res) < k {
		i := heap.Pop(&h).([]int)
		a := i[1]
		b := i[2]
		if flag {
			res = append(res, []int{nums2[b], nums1[a]})
		} else {
			res = append(res, []int{nums1[a], nums2[b]})
		}
		if b+1 < m {
			heap.Push(&h, []int{nums1[a] + nums2[b+1], a, b + 1})
		}
	}
	return res
}

type hp [][]int

func (h hp) Less(i, j int) bool  { return h[i][0] < h[j][0] }
func (h hp) Len() int            { return len(h) }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Pop() interface{}   { o := *h; item := o[len(o)-1]; *h = o[:len(o)-1]; return item }
func (h *hp) Push(a interface{}) { *h = append(*h, a.([]int)) }
