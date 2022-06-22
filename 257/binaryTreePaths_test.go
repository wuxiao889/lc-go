package main

import (
	"fmt"
	. "imooc.com/wuxiao/learnGo/lang/tree"
	"testing"
)

var roots = []*TreeNode{
	tree1(),
}

func tree1() *TreeNode {
	root := NewTree([]int{1, 2, 3, 4})
	return root
}

func TestXXX(t *testing.T) {
	for _, root := range roots {
		fmt.Println(binaryTreePaths(root))
	}
}
