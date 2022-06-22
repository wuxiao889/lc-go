package main

import (
	"fmt"
	. "imooc.com/wuxiao/learnGo/lang/tree"
	"strconv"
)

var paths []string

func binaryTreePaths(root *TreeNode) []string {
	paths = make([]string, 0)
	construct(root, "")
	return paths
}

func construct(root *TreeNode, path string) {
	fmt.Println("root:", root, "path:", path)
	if root == nil {
		return
	}
	path += strconv.Itoa(root.Value)
	if root.Left == nil && root.Right == nil {
		paths = append(paths, path)
		return
	}
	if root.Left != nil {
		construct(root.Left, path+"->")
	}
	if root.Right != nil {
		construct(root.Right, path+"->")
	}
}
