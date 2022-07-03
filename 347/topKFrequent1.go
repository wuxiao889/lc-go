package topKFrequent

//
//type priQueue []struct {
//	val int
//	fre int
//}
//
//func down(q priQueue, i, n int) {
//	for {
//		j1 := 2*i + 1
//		if j1 >= n {
//			return
//		}
//		j := j1
//		if j2 := 2*i + 2; j2 < n && q[j2].fre < q[j1].fre {
//			j = j2
//		}
//		if q[j].fre > q[i].fre {
//			return
//		}
//
//		q[i], q[j] = q[j], q[i]
//		i = j
//	}
//}
//
//func heapsort(q priQueue) {
//	n := len(q)
//	for i := n/2 - 1; i >= 0; i-- {
//		down(q, i, n)
//	}
//	for i := n - 1; i >= 0; i-- {
//		q[0], q[i] = q[i], q[0]
//		down(q, 0, i)
//	}
//}
//
//func topKFrequent1(nums []int, k int) []int {
//	pq := priQueue{}
//
//	ma := make(map[int]int, len(nums))
//	for i := range nums {
//		ma[nums[i]]++
//	}
//
//	for k, v := range ma {
//		pq = append(pq, struct {
//			val int
//			fre int
//		}{val: k, fre: v})
//	}
//	heapsort(pq)
//
//	var res []int
//	for i := 0; i < k; i++ {
//		res = append(res, pq[i].val)
//	}
//	return res
//}
