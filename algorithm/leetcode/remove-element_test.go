package leetcode

//27. 移除元素 https://leetcode-cn.com/problems/remove-element/

/*
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

思路: 和283. 移动零 这个题目一样的解题思路,就是把移动零改成移动val, 将等val的元素放到数组的最后,
使用双指针, i,j; [0,i)保证区间没等于val, 每次把

*/
func removeElement(nums []int, val int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	i := 0
	for j := 0; j < l; j++ {
		if nums[j] != val {
			if j != i && nums[j] != nums[i] { // 值相等避免替换
				nums[j], nums[i] = nums[i], nums[j]
			}
			i++
		}
	}
	return i
}
