package _905

import (
	"fmt"
	"testing"
)

var grid1 = [][]int{
	{1, 0, 1, 0, 1},
	{1, 1, 1, 1, 1},
	{0, 0, 0, 0, 0},
	{1, 1, 1, 1, 1},
	{1, 0, 1, 0, 1},
}

var grid2 = [][]int{
	{0, 0, 0, 0, 0},
	{1, 1, 1, 1, 1},
	{0, 1, 0, 1, 0},
	{0, 1, 0, 1, 0},
	{1, 0, 0, 0, 1},
}

func TestXXX(t *testing.T) {
	fmt.Println(countSubIslands(grid1, grid2))
}
