package LRUCache

import (
	"container/list"
)

type LRUCache struct {
	capacity int
	//为什么用双向链表，因为删除操作需要前驱节点，单向链表无法保证链表的结构
	//双向链表能够保证增改移动的o(1)
	list *list.List
	//map保存对应key值的node节点，实现o(1)的查找
	mp map[int]*list.Element
}

//元素要有key，value。因为只有value删除时就不知道删除map中的哪个
type entry struct {
	key int
	val int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		list:     list.New(),
		mp:       make(map[int]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.mp[key]; ok {
		kv := v.Value.(*entry)
		this.list.MoveToFront(v)
		return kv.val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	//如果存在，更新值，并把节点移到最前
	if v, ok := this.mp[key]; ok {
		this.list.MoveToFront(v)
		ele := v.Value.(*entry)
		//ele是指向结构体的指针类型
		//go语言中，ele.val = (*ele).val
		ele.val = value
		return
	} else {
		//为什么要元素要是指针类型，因为要更新结构体的值
		element := this.list.PushFront(&entry{key, value})
		this.mp[key] = element
	}
	if this.list.Len() > this.capacity {
		back := this.list.Back()
		this.list.Remove(back)
		kv := back.Value.(*entry)
		delete(this.mp, kv.key)
	}
}
