package leetcode

import (
	"sort"
)

// 349. 两个数组的交集 https://leetcode-cn.com/problems/intersection-of-two-arrays/

/*

给定两个数组，编写一个函数来计算它们的交集。

输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2]
输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[9,4]

思路:
1. 直接用hash map进行去重, 会增加空间复杂度, o(m)或o(n)
2. 先进行两个数组排序,通过每个数组一个指针,进行向后移动 . 如果已经排序好的情况下时间复杂度 O(m+n), 空间复杂度O(1)
prometheus 中series 进行求交集的算法
进一步优化思路, 拿范围小的数组nums[0],nums[len-1],去范围大的数组二分查找(变体)找重叠的区间, 再进行2的操作
*/
func intersection2(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	var res []int
	idx1, idx2 := 0, 0
	for idx1 < len(nums1) && idx2 < len(nums2) {
		if nums1[idx1] > nums2[idx2] {
			idx2++
		} else if nums1[idx1] < nums2[idx2] {
			idx1++
		} else {
			tmp := nums1[idx1]
			res = append(res, tmp)
			for idx1 < len(nums1) && nums1[idx1] == tmp {
				idx1++
			}
			for idx2 < len(nums2) && nums2[idx2] == tmp {
				idx2++
			}
		}
	}
	return res
}

func intersection1(nums1 []int, nums2 []int) []int {
	set1 := make(map[int]bool, len(nums1))
	for _, c := range nums1 {
		set1[c] = true
	}
	var res []int
	for _, c := range nums2 {
		if isIn, ok := set1[c]; ok {
			if isIn {
				res = append(res, c)
				set1[c] = false
			}
		}
	}
	return res
}
