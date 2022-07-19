package matrixReshape

func matrixReshape(mat [][]int, r int, c int) [][]int {
	n, m := len(mat), len(mat[0])
	//检查参数，重塑前后矩阵元素个数要相同
	if r*c != n*m {
		return mat
	}
	//行计数器，列计数器
	rc, cc := 0, 0
	res := make([][]int, r)
	for i := range res {
		res[i] = make([]int, c)
		for j := range res[i] {
			//列计数器满了，行计数器指向下一行，列计数器归零
			if cc == m {
				rc++
				cc = 0
			}
			res[i][j] = mat[rc][cc]
			cc++
		}
	}
	return res
}
