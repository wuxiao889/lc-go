package main

import "fmt"

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	i, j := 0, len(matrix[0])-1
	var flag int
	for i < len(matrix) && j >= 0 {
		flag = matrix[i][j]
		if flag == target {
			return true
		} else if flag > target {
			j--
		} else {
			i++
		}
	}
	return false
}
func main() {
	matrixs := [][][]int{
		{},
		{{}},
		{{1, 4, 7, 11, 15},
			{2, 5, 8, 12, 19},
			{3, 6, 9, 16, 22}},
	}
	for _, matrix := range matrixs {
		fmt.Printf("%p,%v\n", matrix, findNumberIn2DArray(matrix, 5))
	}
}
