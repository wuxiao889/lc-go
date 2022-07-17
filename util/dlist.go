package util

type DListNode struct {
	pre, next *DListNode
	Value     any
}

type DList struct {
	len  int
	root DListNode
}

func NewDList() *DList {
	dl := new(DList)
	dl.root.next = &dl.root
	dl.root.pre = &dl.root
	return dl
}

func (l *DList) insert(e, at *DListNode) *DListNode {
	e.pre = at
	e.next = at.next
	at.next = e
	e.next.pre = e
	l.len++
	return e
}

func (l *DList) insertValue(value any, at *DListNode) *DListNode {
	return l.insert(&DListNode{Value: value}, at)
}

func (l *DList) PushFront(value any) *DListNode {
	return l.insertValue(value, &l.root)
}

func (l *DList) PushBack(value any) *DListNode {
	return l.insertValue(value, l.root.pre)
}

func (l *DList) move(e, at *DListNode) {
	if e == at {
		return
	}
	e.pre.next = e.next
	e.next.pre = e.pre

	e.pre = at
	e.next = at.next
	at.next = e
	e.next.pre = e
}

func (l *DList) Remove(e *DListNode) {
	e.pre.next = e.next
	e.next.pre = e.pre
	e.pre = nil
	e.next = nil
	l.len--
}

func (l *DList) MoveToFront(e *DListNode) {
	if l.root.next == e {
		return
	}
	l.move(e, &l.root)
}

func (l *DList) MoveToBack(e *DListNode) {
	if l.root.pre == e {
		return
	}
	l.move(e, l.root.pre)
}

func (l *DList) Len() int {
	return l.len
}

func (l *DList) Front() *DListNode {
	if l.Len() == 0 {
		return nil
	}
	return l.root.next
}

func (l *DList) Back() *DListNode {
	if l.Len() == 0 {
		return nil
	}
	return l.root.pre
}
