package exist

var row, col int
var Board [][]byte
var Word string
var used [][]bool

func exist(board [][]byte, word string) bool {
	Board, Word = board, word
	row, col = len(board), len(board[0])
	used = make([][]bool, len(board))
	for i := range used {
		used[i] = make([]bool, len(board[0]))
	}
	for i := range board {
		for j := range board[i] {
			if board[i][j] == word[0] {
				if backtrace(i, j, 0, used) {
					return true
				}
			}
		}
	}
	return false
}

var dirs = [][]int{{0, -1}, {-1, 0}, {0, +1}, {+1, 0}}

func backtrace(i, j, k int, used [][]bool) bool {
	if k == len(Word) {
		return true
	}
	if i < 0 || j < 0 || i == row || j == col || used[i][j] {
		return false
	}
	if Board[i][j] != Word[k] {
		return false
	}
	used[i][j] = true
	for _, dir := range dirs {
		if backtrace(i+dir[0], j+dir[1], k+1, used) {
			return true
		}
	}
	used[i][j] = false
	return false
}
