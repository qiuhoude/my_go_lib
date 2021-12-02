package leetcode

import "sort"

// 217. 存在重复元素 https://leetcode-cn.com/problems/contains-duplicate/

/*
给定一个整数数组，判断是否存在重复元素。
如果存在一值在数组中出现至少两次，函数返回 true 。如果数组中每个元素都不相同，则返回 false 。

输入: [1,2,3,1]
输出: true

输入: [1,2,3,4]
输出: false

输入: [1,1,1,3,3,4,3,2,4,2]
输出: true

思路:
1.丢到hash表记录次数
2.排序后比较前后值,记录相同值的次数
*/

func containsDuplicate(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	tab := make(map[int]bool, len(nums))
	for _, v := range nums {
		if _, ok := tab[v]; ok {
			return true
		}
		tab[v] = true
	}
	return false
}

func containsDuplicate1(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	sort.Ints(nums)
	sameCnt := 0 // 记录相等次数
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			sameCnt++
		} else {
			sameCnt = 0 // 相同值
		}
		if sameCnt >= 1 {
			return true
		}
	}
	return false
}
