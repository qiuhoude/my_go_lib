package leetcode

import (
	"sort"
)

// 16. 最接近的三数之和 https://leetcode-cn.com/problems/3sum-closest/

/*
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。
返回这三个数的和。假定每组输入只存在唯一答案。

例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).

思路:
排序 + 3 指针
*/

func threeSumClosest(nums []int, target int) int {
	// 思路: 1. 排序
	// 2. 也是使用三指针
	sort.Ints(nums)
	length := len(nums)
	ret := nums[0] + nums[1] + nums[2]

	for i := 0; i < length; i++ {
		l := i + 1
		r := length - 1
		if i > 0 && nums[i] == nums[i-1] { // 重复的去掉
			continue
		}
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if abs(target-sum) < abs(target-ret) {
				ret = sum
			}
			if target == sum { // 找到目的直接是最接近目标的
				ret = sum
				return ret
			} else if sum < target {
				l++
			} else if sum > target {
				r--
			}
		}
	}
	return ret
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
