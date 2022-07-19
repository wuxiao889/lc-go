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

//log层，每一层复杂度，每一层每个元素只会被操作1次，总共n次。
//o(n) = nlogn
//时间复杂度总是nlogn，因为总会被分成log层，然后比较n次，与原数组有没有排序，是什么数组无关
func mergeSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	mid := l + (r-l)>>1
	//递归排序
	mergeSort(nums, l, mid)
	mergeSort(nums, mid+1, r)
	//合并
	tmp := make([]int, r-l+1)
	i, j, top := l, mid+1, 0
	for i <= mid && j <= r {
		if nums[i] < nums[j] {
			tmp[top] = nums[i]
			i++
		} else {
			tmp[top] = nums[j]
			j++
		}
		top++
	}
	for i <= mid {
		tmp[top] = nums[i]
		top, i = top+1, i+1
	}
	for j <= r {
		tmp[top] = nums[j]
		top, j = top+1, j+1
	}
	copy(nums[l:r+1], tmp)
}

//0 n 2*n+1 2*n+2 (n-1)/2
//从小到大排序，这是一个大根堆
func heapSort(nums []int) {
	n := len(nums)
	//堆化，从最后一个父节点开始down
	for i := n/2 - 1; i >= 0; i-- {
		down(nums, i, n)
	}
	//堆排序，把堆顶的和末尾交换，交换后堆顶开始down,并把以排序的排除在外
	for i := n - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		down(nums, 0, i)
	}
}

//i要down的元素，n时下界
func down(nums []int, i, n int) {
	for {
		j1 := 2*i + 1
		//左儿子出界
		if j1 >= n {
			return
		}
		j := j1
		//有右儿子，选出最大的儿子
		if j2 := j1 + 1; j2 < n && nums[j2] > nums[j1] {
			j = j2
		}
		//如果满足父大于儿子，返回
		if nums[i] > nums[j] {
			return
		}
		nums[i], nums[j] = nums[j], nums[i]
		//继续向下down
		i = j
	}
}
