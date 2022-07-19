package main

import (
	"fmt"
)

var (
	m, n int
	a    [1001]int
	b    [1001]int
	f    [1001][1001]int
)

func main() {
	fmt.Scanf("%d%d\n", &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	fmt.Scanf("\n")
	for i := 1; i <= m; i++ {
		fmt.Scanf("%d", &b[i])
	}
	fmt.Scanf("\n")
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			f[i][j] = max(f[i-1][j], f[i][j-1])
			if a[i] == b[j] {
				f[i][j] = max(f[i][j], f[i-1][j-1]+1)
			}
		}
	}
	fmt.Println(f[n][m])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
