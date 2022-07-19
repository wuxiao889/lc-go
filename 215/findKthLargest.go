package findKthLargest

import (
	"math/rand"
	"time"
)

//o(n)的复杂度，利用快排排序后选定的元素位置固定
//第k大，那么再排序后的数组中的第k-1个索引位置
//当快排i==j==k时，返回nums[i]
//利用快排的分治，减少不必要的排序
func findKthLargest(nums []int, k int) int {
	return quickSort(nums, 0, len(nums)-1, k-1)
}

func quickSort(arr []int, l, r, k int) int {
	//加入随机性，避免极端案例
	rand.Seed(time.Now().UnixNano())
	n := l + rand.Int()%(r-l+1)
	arr[l], arr[n] = arr[n], arr[l]

	x := arr[l]
	i, j := l, r
	for i < j {
		for i < j && arr[j] < x {
			j--
		}
		if i < j {
			arr[i] = arr[j]
			i++
		}
		for i < j && arr[i] > x {
			i++
		}
		if i < j {
			arr[j] = arr[i]
			j--
		}
	}
	arr[i] = x
	if i == k {
		return arr[i]
	}
	if i > k {
		return quickSort(arr, l, i-1, k)
	} else {
		return quickSort(arr, i+1, r, k)
	}
}

//堆排序方法
func heapSort(nums []int, k int) int {
	n := len(nums)
	for i := n/2 - 1; i >= 0; i-- {
		down(nums, i, n)
	}
	for i := n - 1; i > n-k; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		down(nums, 0, i)
	}
	return nums[0]
}

func down(nums []int, i, n int) {
	for {
		j1 := 2*i + 1
		if j1 >= n {
			return
		}
		j := j1
		if j2 := j1 + 1; j2 < n && nums[j2] > nums[j1] {
			j = j2
		}
		if nums[i] > nums[j] {
			return
		}
		nums[i], nums[j] = nums[j], nums[i]
		i = j
	}
}
