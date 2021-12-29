package leetcode

import "sort"

// 90. 子集 II https://leetcode-cn.com/problems/subsets-ii/
/*
给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。

输入：nums = [1,2,2]
输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]

输入：nums = [0]
输出：[[],[0]]

1 <= nums.length <= 10
-10 <= nums[i] <= 10

思路
遇到重复的问题，一般都是先排序, 思路和78题一致,加上去重逻辑
*/

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	var dfsFn func(int)
	var queue []int
	dfsFn = func(start int) {
		res = append(res, append([]int(nil), queue...)) // 每次进来添加
		for i := start; i < len(nums); i++ {
			if i != start && nums[i] == nums[i-1] { // 去掉重复
				continue
			}
			queue = append(queue, nums[i])
			dfsFn(i + 1)
			queue = queue[:len(queue)-1]
		}
	}
	dfsFn(0)
	return res
}
