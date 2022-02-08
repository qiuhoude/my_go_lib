package leetcode

import "sort"

// 169. 多数元素 https://leetcode-cn.com/problems/majority-element/

/*
给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。

输入：[3,2,3]
输出：3

输入：[2,2,1,1,1,2,2]
输出：2

进阶：
尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。

思路:
1. hash表计数,最基础计数

2. 先排序, 将数组一份为2,多数元素一定在中间  时间复杂度 O(nlogn)

3. 摩尔投票法
假设第一个元素为多数元素,最开始count=1 count==0就换元素,遇到相同就count++,不同就count--,,
因为多数元素是超过n/2的元素,所以count一定是个正数
*/

func majorityElement(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	return nums[n/2]
}

func majorityElement3(nums []int) int {
	res := nums[0]
	count := 0
	for _, v := range nums {
		if count == 0 {
			res = v
		}
		if res == v {
			count++
		} else {
			count--
		}
	}
	return res
}
