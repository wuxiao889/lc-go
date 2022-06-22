package main

func reverseStr(s string, k int) string {
	str := []byte(s)
	for i := 0; i < len(s); i += 2 * k {
		sub := str[i:min(i+k, len(s))]
		start, end := 0, len(sub)-1
		for start < end {
			sub[start], sub[end] = sub[end], sub[start]
			end--
			start++
		}
	}
	return string(str)
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
