package leetcode

//26. 删除有序数组中的重复项 https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/

/*
给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

输入：nums = [1,1,2]
输出：2, nums = [1,2]
解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。

输入：nums = [0,0,1,1,1,2,2,3,3,4]
输出：5, nums = [0,1,2,3,4]
解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素。

0 <= nums.length <= 3 * 10^4
-10^4 <= nums[i] <= 10^4
nums 已按升序排列

思路:
利用分区的思想 [0,k) 区间是没有重复元素的区域,每次填充无重复元素的区间

*/

func removeDuplicates(nums []int) int {
	l := len(nums)
	if l < 2 {
		return l
	}

	k := 1          // [0,k)是非重复元素
	curR := nums[0] // 当前重复的元素的值
	for i := 1; i < l; i++ {
		if nums[i] != curR { // 非重复元素
			curR = nums[i]
			nums[k] = curR
			k++
		}
	}
	//nums = append(nums[:k])
	return k
}
