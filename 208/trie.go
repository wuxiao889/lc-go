package trie

type Trie struct {
	isEnd bool
	next  map[byte]*Trie
}

func Constructor() Trie {
	return Trie{
		isEnd: false,
		next:  map[byte]*Trie{},
	}
}

func (this *Trie) Insert(word string) {
	cur := this
	for i := range word {
		c := word[i]
		if cur.next[c-'a'] == nil {
			cur.next[c-'a'] = &Trie{
				isEnd: false,
				next:  map[byte]*Trie{},
			}
		}
		cur = cur.next[c-'a']
	}
	cur.isEnd = true
}

func (this *Trie) Search(word string) bool {
	cur := this
	for i := range word {
		c := word[i]
		if cur.next[c-'a'] == nil {
			return false
		}
		cur = cur.next[c-'a']
	}
	return cur.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for i := range prefix {
		c := prefix[i]
		if cur.next[c-'a'] == nil {
			return false
		}
		cur = cur.next[c-'a']
	}
	return true
}

/**
// * Your Trie object will be instantiated and called as such:
// * obj := Constructor();
// * obj.Insert(word);
// * param_2 := obj.Search(word);
// * param_3 := obj.StartsWith(prefix);
// */
