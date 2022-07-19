package main

import "fmt"

//http://oj.daimayuan.top/course/5/problem/161

func main() {
	p3()
}

// v[i] 物品i的重量
// w[i] 物品i的价值
// n 物品数
// m 总重量
var (
	v    [1001]int
	w    [1001]int
	n, m int
)

func p1() {
	// f[i][j]定义为前i个物品重量为j时的最大收益
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
				f[i][j] = max(f[i-1][j], f[i-1][j-v[i]]+w[i])
			}
		}
	}
	fmt.Println(f[n][m])
}

func p2() {
	var (
		f = make([]int, 1001)
		g = make([]int, 1001)
	)
	fmt.Scanf("%d%d\n", &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d%d\n", &v[i], &w[i])
	}
	for i := 1; i <= n; i++ {
		for j := 0; j <= m; j++ {
			if j < v[i] {
				g[j] = f[j]
			} else {
				g[j] = max(f[j], f[j-v[i]]+w[i])
			}
		}
		copy(f, g)
	}
	fmt.Println(f[m])
}

func p3() {
	var (
		f = make([]int, 1001)
	)
	fmt.Scanf("%d%d\n", &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d%d\n", &v[i], &w[i])
	}
	for i := 1; i <= n; i++ {
		for j := m; j >= v[i]; j-- {
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
