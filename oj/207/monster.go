package main

import "fmt"

var (
	n    int
	a, b [100001]int
)

func main() {
	p2()
}

func p1() {
	var c [100001][2]int
	fmt.Scanf("%d\n", &n)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d%d\n", &a[i], &b[i])
	}
	l, k := 1, 0
	for i := 1; i <= n; i++ {
		for ; k >= l && b[i] >= c[k][0]; k-- {
		}
		k++
		c[k][0] = b[i]
		c[k][1] = a[i]
		fmt.Println(c[l][0])
		for ; k >= l && c[l][1] == i; l++ {
		}
	}
}

func p2() {
	c := make([][2]int, 0, 100001)
	fmt.Scanf("%d\n", &n)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d%d\n", &a[i], &b[i])
	}
	for i := 1; i <= n; i++ {
		for len(c) > 0 && b[i] > c[len(c)-1][0] {
			c = c[:len(c)-1]
		}
		c = append(c, [2]int{b[i], a[i]})
		fmt.Println(c[0][0])
		for len(c) > 0 && c[0][1] == i {
			c = c[1:]
		}
	}
}
