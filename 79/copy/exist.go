package copy

func exist(board [][]byte, word string) bool {
	cache := make([][]bool, len(board))
	for i := range board {
		cache[i] = make([]bool, len(board[0]))
	}

	res := false

	var dfs func(i, j, n int)
	dfs = func(i, j, n int) {
		if n == len(word) {
			res = true
			return
		}
		cache[i][j] = true
		if i-1 >= 0 && !cache[i-1][j] && board[i-1][j] == word[n] {
			dfs(i-1, j, n+1)
		}
		if i+1 < len(board) && !cache[i+1][j] && board[i+1][j] == word[n] {
			dfs(i+1, j, n+1)
		}
		if j-1 >= 0 && !cache[i][j-1] && board[i][j-1] == word[n] {
			dfs(i, j-1, n+1)
		}
		if j+1 < len(board[0]) && !cache[i][j+1] && board[i][j+1] == word[n] {
			dfs(i, j+1, n+1)
		}
		cache[i][j] = false
		return
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				dfs(i, j, 1)
			}
		}
	}

	return res
}
