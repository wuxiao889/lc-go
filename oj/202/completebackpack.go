package main

import "fmt"

func main() {
	p2()
}

var (
	n, m int
	v, w [1001]int
)

func p1() {
	var f [1001][1001]int
	fmt.Scanf("%d%d\n", &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d%d\n", &v[i], &w[i])
	}
	for i := 1; i <= n; i++ {
		for j := 0; j <= m; j++ {
			if j < v[i] {
				f[i][j] = f[i-1][j]
			} else {
				f[i][j] = max(f[i-1][j], f[i][j-v[i]]+w[i])
			}
		}
	}
	fmt.Println(f[n][m])
}

func p2() {
	var f [1001]int
	fmt.Scanf("%d%d\n", &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d%d\n", &v[i], &w[i])
	}
	for i := 1; i <= n; i++ {
		for j := v[i]; j <= m; j++ {
			f[j] = max(f[j], f[j-v[i]]+w[i])
		}
	}
	fmt.Println(f[m])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
