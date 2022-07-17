package LRUCache

import "testing"

func TestLruCache(t *testing.T) {
	lru := Constructor(2)
	lru.Get(2)
	lru.Put(2, 6)
	lru.Get(1)
	lru.Put(1, 5)
	lru.Put(1, 2)
	lru.Get(1)
	lru.Get(1)
}
