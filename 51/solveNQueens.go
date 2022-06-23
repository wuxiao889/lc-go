package main

import (
	"fmt"
	"strings"
)

func main() {
	queens := solveNQueens(8)
	for i := range queens {
		fmt.Println("___")
		for j := range queens[i] {
			fmt.Println(queens[i][j])
		}
	}
}

var chessBoard [][]string
var res [][]string

func solveNQueens(n int) [][]string {
	res = make([][]string, 0)
	chessBoard = make([][]string, n)
	for i := range chessBoard {
		chessBoard[i] = make([]string, n)
		for j := range chessBoard {
			chessBoard[i][j] = "."
		}
	}
	backtrace(0)
	return res
}

func backtrace(row int) {
	if row == len(chessBoard) {
		temp := make([]string, len(chessBoard))
		for i := range chessBoard {
			temp[i] = strings.Join(chessBoard[i], "")
		}
		res = append(res, temp)
		return
	}
	for col := range chessBoard {
		if !isValid(row, col) {
			continue
		}
		chessBoard[row][col] = "Q"
		backtrace(row + 1)
		chessBoard[row][col] = "."
	}
}

func isValid(row int, col int) bool {
	for i := row; i >= 0; i-- {
		if chessBoard[i][col] == "Q" {
			return false
		}
	}
	for i, j := row-1, col-1; j >= 0 && i >= 0; i, j = i-1, j-1 {
		if chessBoard[i][j] == "Q" {
			return false
		}
	}
	for i, j := row-1, col+1; j < len(chessBoard) && i >= 0; i, j = i-1, j+1 {
		if chessBoard[i][j] == "Q" {
			return false
		}
	}
	return true
}
