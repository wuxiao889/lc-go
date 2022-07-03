package mergeKLists

type ListNode struct {
	Val  int
	Next *ListNode
}

type Lq []*ListNode

func (q *Lq) push(node *ListNode) {
	*q = append(*q, node)
	q.up(len(*q) - 1)
}

func (q *Lq) pop() *ListNode {
	old := *q
	n := len(*q) - 1
	old[0], old[n] = old[n], old[0]
	old.down(0, n)

	node := (*q)[n]
	(*q)[n] = nil
	*q = old[:n]
	return node
}

func (q Lq) up(j int) {
	for {
		i := (j - 1) / 2
		if i == j {
			break
		}
		if q[i].Val < q[j].Val {
			break
		}
		q[i], q[j] = q[j], q[i]
		j = i
	}
}

func (q Lq) down(i, n int) {
	for {
		j1 := 2*i + 1
		if j1 >= n {
			break
		}
		j := j1
		if j2 := j1 + 1; j2 < n && q[j2].Val < q[j1].Val {
			j = j2
		}
		if q[i].Val < q[j].Val {
			break
		}
		q[i], q[j] = q[j], q[i]
		i = j
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
	preh := &ListNode{}
	lq := Lq{}
	for i := range lists {
		lq.push(lists[i])
	}
	cur := preh
	for len(lq) > 0 {
		node := lq.pop()
		if node.Next != nil {
			lq.push(node.Next)
		}
		cur.Next = node
		cur = cur.Next
		//cur.next = nil
		//并行赋值不能这样写
		//cur.Next, cur = node, cur.Next
	}
	return preh.Next
}
