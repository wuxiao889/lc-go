package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var res int
	for x != 0 {
		res = res*10 + x%10
		x = x / 10
		if res > math.MaxInt32 || res < -math.MaxInt32-1 {
			return 0
		}
	}
	return res
}

func main() {
	arr := []int{
		1999,
		9991,
		-1,
		-99885,
	}
	for _, v := range arr {
		fmt.Println(v, reverse(v))
	}
}
