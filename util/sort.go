package util

import (
	"fmt"
	"math/rand"
)

//时间复杂度o(nlogn)
func quickSort(nums []int, l, r int) {
	fmt.Println(nums)
	if l >= r {
		return
	}
	n := l + rand.Int()%(r-l+1)
	//随机一个数把他放在left作为x
	//否则时间复杂度最差会到o(n^2)
	//比如1，2，3，4，5
	//每次都要遍历到n个元素
	//o = n * (n-1) * ... * 1
	nums[l], nums[n] = nums[n], nums[l]
	x := nums[l]
	i, j := l, r
	//关键是要记住每次都要判断i<j
	for i < j {
		//j往左找第一个小于等于x的元素
		for i < j && nums[j] > x {
			j--
		}
		//i<j才要交换，才要i++
		if i < j {
			nums[i] = nums[j]
			//i要向右移动一位
			i++
		}
		for i < j && nums[i] < x {
			i++
		}
		if i < j {
			nums[j] = nums[i]
			j--
		}
	}
	nums[i] = x
	quickSort(nums, l, i-1)
	quickSort(nums, i+1, r)
}

//多关键字排序，先比较值，再比较下标
func quickSortStable(nums []int, i, j int) {

}
