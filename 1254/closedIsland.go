package _254

var row, col int

func closedIsland(grid [][]int) int {
	row, col = len(grid), len(grid[0])
	var res int
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 0 && dfs(i, j, grid) {
				res++
			}
		}
	}
	return res
}

func dfs(i, j int, grid [][]int) bool {
	if i < 0 || j < 0 || i == row || j == col || grid[i][j] == 1 {
		return true
	}
	if i == 0 || i == row-1 || j == 0 || j == col-1 {
		return false
	}
	grid[i][j] = 1
	return dfs(i-1, j, grid) && dfs(i, j+1, grid) && dfs(i+1, j, grid) && dfs(i, j-1, grid)
}
