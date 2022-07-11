package main

import "fmt"

var (
	n, m    int
	a, v, w [1001]int
)

func main() {
	p2()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func p1() {
	var f [1001][1001]int
	c := make([][]int, 1001)
	fmt.Scanf("%d%d\n", &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d%d%d\n", &a[i], &v[i], &w[i])
		c[a[i]] = append(c[a[i]], i)
	}
	for i := 1; i <= 1000; i++ {
		for j := 0; j <= m; j++ {
			f[i][j] = f[i-1][j]
			for _, k := range c[i] {
				if v[k] <= j {
					f[i][j] = max(f[i][j], f[i-1][j-v[k]]+w[k])
				}
			}
		}
	}
	fmt.Println(f[1000][m])
}

func p2() {
	var f [1001]int
	c := make([][]int, 1001)
	fmt.Scanf("%d%d\n", &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d%d%d\n", &a[i], &v[i], &w[i])
		c[a[i]] = append(c[a[i]], i)
	}
	for i := 1; i <= 1000; i++ {
		for j := m; j > 0; j-- {
			for _, k := range c[i] {
				if v[k] <= j {
					f[j] = max(f[j], f[j-v[k]]+w[k])
				}
			}
		}
	}
	fmt.Println(f[m])
}
