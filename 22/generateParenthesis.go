package main

import "fmt"

var res []string
var path []byte

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	res = make([]string, 0)
	path = make([]byte, 0)
	backtrace(0, 0, n)
	return res
}

func backtrace(left, right, n int) {
	if left > n || right > n {
		return
	}
	if left < right {
		return
	}
	if left == right && left == n {
		s := string(path)
		res = append(res, s)
		return
	}
	path = append(path, '(')
	backtrace(left+1, right, n)
	path = path[:len(path)-1]

	path = append(path, ')')
	backtrace(left, right+1, n)
	path = path[:len(path)-1]
}
