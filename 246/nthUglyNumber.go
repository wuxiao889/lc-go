package nthUglyNumber

import (
	"container/heap"
	"sort"
)

type Heap struct{ sort.IntSlice }

func (h *Heap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *Heap) Pop() interface{} {
	oldSlice := h.IntSlice
	item := oldSlice[len(oldSlice)-1]
	h.IntSlice = oldSlice[:len(oldSlice)-1]
	return item
}

func nthUglyNumber(n int) int {
	h := &Heap{sort.IntSlice{}}
	used := map[int]struct{}{1: {}}
	nums := []int{2, 3, 5}
	heap.Push(h, 1)

	for i := 1; ; i++ {
		res := heap.Pop(h).(int)
		if i == n {
			return res
		}
		for i := range nums {
			num := res * nums[i]
			if _, has := used[num]; !has {
				heap.Push(h, num)
				used[num] = struct{}{}
			}
		}
	}
}
