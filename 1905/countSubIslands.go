package _905

import "fmt"

var row, col int

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	row, col = len(grid1), len(grid1[0])
	var res int
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid2[i][j] == 1 {
				if dfs(i, j, grid1, grid2) {
					res++
				}
			}
		}
	}
	return res
}

func dfs(i, j int, grid1, grid2 [][]int) bool {
	if i < 0 || j < 0 || i >= row || j >= col || grid2[i][j] == 0 {
		return true
	}
	if grid1[i][j] == 0 {
		return false
	}
	grid2[i][j] = 0
	fmt.Println("---------")
	for _, v := range grid2 {
		fmt.Println(v)
	}
	return dfs(i, j-1, grid1, grid2) && dfs(i-1, j, grid1, grid2) && dfs(i, j+1, grid1, grid2) && dfs(i+1, j, grid1, grid2)
}
