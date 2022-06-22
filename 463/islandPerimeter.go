package _63

var row, col int

func islandPerimeter(grid [][]int) int {
	row, col = len(grid), len(grid[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 1 {
				return dfs(i, j, grid)
			}
		}
	}
	return 0
}

func dfs(m, n int, grid [][]int) int {
	if m < 0 || m >= row || n < 0 || n >= col || grid[m][n] == 0 {
		return 1
	}
	grid[m][n] = 2
	return dfs(m-1, n, grid) + dfs(m+1, n, grid) + dfs(m, n+1, grid) + dfs(m, n-1, grid)
}
