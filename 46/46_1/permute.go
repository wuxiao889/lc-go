package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3, 4}))
}

var res [][]int

func permute(nums []int) [][]int {
	res = make([][]int, 0)
	track := make([]int, 0)
	used := make([]bool, len(nums))
	backtrack(nums, track, used)
	return res
}

func backtrack(nums, track []int, used []bool) {
	if len(track) == len(nums) {
		//fmt.Printf("before append:track:%v ,adress of 0th element:%p\n", track, track)
		res = append(res, append([]int{}, track...))
		return
	}
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		track = append(track, nums[i])
		used[i] = true
		backtrack(nums, track, used)
		track = track[:len(track)-1]
		used[i] = false
	}
}
