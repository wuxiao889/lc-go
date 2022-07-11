package trie

import "fmt"

func ExampleTrie() {
	obj := Constructor()
	obj.Insert("hello")
	obj.Insert("h")
	fmt.Println(obj.Search("hello"))
	fmt.Println(obj.StartsWith("he"))
	//output:
	//true
	//true
}
