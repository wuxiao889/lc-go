package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3, 4}))
}

var res [][]int
var track []int

func permute(nums []int) [][]int {
	res = make([][]int, 0)
	track = make([]int, 0)
	used := make([]bool, len(nums))
	backtrack(nums, used)
	return res
}

func backtrack(nums []int, used []bool) {
	if len(track) == len(nums) {
		//fmt.Printf("track:%v ,adress of 0th element:%p \n", track, track)
		//res = append(res, track)
		res = append(res, append([]int{}, track...))
		return
	}
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		track = append(track, nums[i])
		used[i] = true
		backtrack(nums, used)
		track = track[:len(track)-1]
		used[i] = false
	}
}
