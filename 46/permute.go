package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3, 4}))
}

var res [][]int
var track []int
var used []bool

func permute(nums []int) [][]int {
	res = [][]int{}
	track = []int{}
	used = make([]bool, len(nums))
	backtrack(nums)
	return res
}

func backtrack(nums []int) {
	if len(track) == len(nums) {
		res = append(res, append([]int{}, track...))
		return
	}
	for i := range nums {
		if used[i] {
			continue
		}
		track = append(track, nums[i])
		used[i] = true
		backtrack(nums)
		track = track[:len(track)-1]
		used[i] = false
	}
}
