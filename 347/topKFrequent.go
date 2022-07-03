package topKFrequent

import "container/heap"

type A struct {
	val int
	fre int
}

type As []A

func (a As) Len() int {
	return len(a)
}

func (a As) Less(i, j int) bool {
	return a[i].fre > a[j].fre
}

func (a As) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a *As) Push(x any) {
	*a = append(*a, x.(A))
}

func (a *As) Pop() any {
	old := *a
	n := a.Len()
	item := old[n-1]
	*a = old[:n-1]
	return item.val
}

var kk int

func topKFrequent(nums []int, k int) []int {
	pq := &As{}
	kk = k
	ma := make(map[int]int, len(nums))
	for i := range nums {
		ma[nums[i]]++
	}
	for k, v := range ma {
		heap.Push(pq, A{
			val: k,
			fre: v,
		})
	}
	var res []int
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(pq).(int))
	}
	return res
}
