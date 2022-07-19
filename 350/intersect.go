package intersect

import (
	"sort"
)

//https://leetcode.cn/problems/intersection-of-two-arrays-ii/solution/jin-jie-san-wen-by-user5707f/

//哈希表，为了降低空间复杂度，选择遍历较短的数组
//时间复杂度o(m+n)
//空间复杂度o(min(m,n))
func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}
	var res []int
	mp := make(map[int]int, 10000)
	for i := range nums1 {
		mp[nums1[i]]++
	}
	for i := range nums2 {
		if _, ok := mp[nums2[i]]; ok && mp[nums2[i]] > 0 {
			res = append(res, nums2[i])
			mp[nums2[i]]--
		}
	}
	return res
}

//排序+双指针
//时间复杂度(mlogm + nlogn)
//
func intersect_1(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	var res []int
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			res = append(res, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return res
}

//func sort(nums []int) {
//	quicksort(nums, 0, len(nums)-1)
//}
//
//func quicksort(nums []int, l, r int) {
//	if l >= r {
//		return
//	}
//	rand.Seed(time.Now().UnixNano())
//	n := l + rand.Int()%(r-l+1)
//	nums[l], nums[n] = nums[n], nums[l]
//
//	x := nums[l]
//	i, j := l, r
//	for i < j {
//		for i < j && nums[j] > x {
//			j--
//		}
//		if i < j {
//			nums[i] = nums[j]
//			i++
//		}
//		for i < j && nums[i] < x {
//			i++
//		}
//		if i < j {
//			nums[j] = nums[i]
//			j--
//		}
//	}
//	nums[i] = x
//	quicksort(nums, l, i-1)
//	quicksort(nums, i+1, r)
//}
