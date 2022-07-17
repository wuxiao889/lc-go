package util

import (
	"fmt"
	"testing"
)

func TestDlist(t *testing.T) {
	dl := NewDList()
	dl.PushFront(1)
	dl.PushFront(2)
	cur := dl.root.next
	for cur != &dl.root {
		fmt.Println(cur.Value.(int))
		cur = cur.next
	}
	dl.PushBack(4)
	cur = dl.root.next
	for cur != &dl.root {
		fmt.Println(cur.Value.(int))
		cur = cur.next
	}
}
