package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	left, right := 0, 0
	window := [256]int{}
	var res int
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for right < len(s) {
		c := s[right]
		right++
		window[c]++

		for window[c] > 1 {
			d := s[left]
			window[d]--
			left++
		}

		res = max(res, right-left)
	}

	return res
}

func main() {
	arr := []string{
		"abcab cbb",
		"bbbbb",
	}
	for _, v := range arr {
		fmt.Println(lengthOfLongestSubstring(v))
	}
	c := byte('1')
	fmt.Println('1')
	fmt.Println('a')
	fmt.Println(c - 'a')
}
